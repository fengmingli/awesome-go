package main

import (
	"encoding/base64"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// dtle浮点数序列化算法
func DecimalValueFromStringMysql(value string) string {
	//初始化一个大整数
	sum := new(big.Int)

	//如果为负数，将-号去掉，并标记为负数。
	isNeg := false
	if value[0] == '-' {
		value = value[1:]
		isNeg = true
	}

	// 将字符串转变为整数，例如：96646.13 -> 9654513
	for i := range value {
		if '0' <= value[i] && value[i] <= '9' {
			sum.Mul(sum, big.NewInt(10))
			offset := int(value[i] - '0')
			sum.Add(sum, big.NewInt(int64(offset)))
		}
	}

	bs := sum.Bytes()
	if len(bs) == 0 {
		bs = []byte{0}
	}

	//如果时负数，取反，再加1
	if isNeg {
		for i := len(bs) - 1; i >= 0; i-- {
			bs[i] = ^bs[i]
		}
		for i := len(bs) - 1; i >= 0; i-- {
			bs[i] += 1
			if bs[i] != 0x00 {
				break
			}
		}
	} else if bs[0] > 0x7f { //如果是正数，首字节大于0x7f（127），添加一个字节0x00
		bs2 := make([]byte, len(bs)+1)
		bs2[0] = 0x00
		copy(bs2[1:], bs)
		bs = bs2
	}
	//fmt.Printf("原字符串:%s ，2 进制补码为：%%0%db\n", value, bs)
	//添加日志
	//fmt.Printf(fmt.Sprintf("原始数据:%s ，补码表示为: %08b\n", value, bs))
	return base64.StdEncoding.EncodeToString(bs)
}

func deserialize(bigDecimal string, scale, precision int) string {
	decodedBytes, err := base64.StdEncoding.DecodeString(bigDecimal)
	if err != nil {
		return ""
	}
	//isNeg := false
	//if decodedBytes[0]&0x00 == 0 {
	//	isNeg = true
	//	decodedBytes = decodedBytes[1:]
	//}

	//将每个字节转换为2进制字符串
	var binaryStr strings.Builder
	for _, b := range decodedBytes {
		binaryStr.WriteString(fmt.Sprintf("%08b", b))
	}

	//将2进制字符串解析为大整数
	intValue := new(big.Int)
	intValue.SetString(binaryStr.String(), 2)

	//处理原始数据的符号，确保正确处理原始数据的符号。
	//如果数据的最高 4 位中的任何一位不为零，它将将整数表示为负数，然后进行调整。
	//即：最高4位和0x70进行&运行。如果不等于0成立，表示为负数，需要将负数的补码转为原码
	if decodedBytes[0]&0x70 != 0 || decodedBytes[0]&0x80 != 0 {

		//创建一个大整数 negInt 并将其设置为值为 1。
		negInt := new(big.Int).SetUint64(1)
		//左移len(decodedBytes)*8位，以创建一个负数的表示，位数与原始数据相同。
		negInt.Lsh(negInt, uint(len(decodedBytes)*8))
		//将 intValue（原始整数）减去 negInt（负数表示），从而得到正确的整数
		intValue.Sub(intValue, negInt)
	}

	//缩放因子：计算10^scale
	scaleFactor := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(scale)), nil)
	//设置有理数的分子和分母
	result := new(big.Rat).SetFrac(intValue, scaleFactor)
	//精度限制。确保小数点后不超过指定的位数
	floatString := result.FloatString(precision)

	//日志打印
	//fmt.Printf(fmt.Sprintf("反序列化过程->需要反序列化字符串：%s，正、负数判断：（首位 %08b & 0x70 == %v）true 为负，反之正数. 反序列结果：%s\n", bigDecimal, decodedBytes[0], decodedBytes[0]&0x70 != 0, floatString))

	return floatString
}

func main() {
	bech()
	//test2()
}

func bech() {
	bytesNum := int64(3) //测试 1-3个字节的数据
	for i := int64(1); i <= bytesNum; i++ {
		count := int64(0)
		exceptionData := int64(0)
		byteSum := 8 * i
		for j := int64(0); j <= 2<<(byteSum-1)-1; j++ {
			count++
			res := DecimalValueFromStringMysql(strconv.FormatInt(j, 10))
			//res2 := DecimalValueFromStringMysql(strconv.FormatInt(j-2<<(byteSum-1), 10))
			//if res == res2 {
			//	exceptionData++
			//}
			s := deserialize(res, 0, 0)
			if s != strconv.FormatInt(j, 10) {
				exceptionData++
			}
		}

		fmt.Printf("当正数字节数为:%d字节时,总数：%d,正负数序列化结果一致: %d ,无法序列化占比: %.4f \n", i, count, exceptionData, float64(exceptionData)/float64(count))

		count2 := int64(0)
		exceptionData2 := int64(0)
		for j := int64(-2<<(byteSum-1) + 1); j < 0; j++ {
			count2++
			res2 := DecimalValueFromStringMysql(strconv.FormatInt(j, 10))
			s2 := deserialize(res2, 0, 0)
			if s2 != strconv.FormatInt(j, 10) {
				exceptionData2++
			}

		}
		fmt.Printf("当负数字节数为:%d字节时,总数：%d,异常数: %d ,异常占比: %.4f \n", i, count2, exceptionData2, float64(exceptionData2)/float64(count2))
	}
}

func test2() {

	//case-1: 正数，首字节 > 0x7f，头部会加上0x00。（正常）
	fmt.Println("=====================case-1=====================")
	fmt.Println("case-1:正数，首字节 > 0x7f，头部会加上0x00。（正常）")
	case1 := ""
	result1 := DecimalValueFromStringMysql(case1)
	fmt.Printf("序列化过程->原始数据:%s ,序列化结果：%s \n", case1, result1)
	deserialize(result1, 2, 2)
	//case-2：正数，首字节 < 0x7f，头部不会加上0x00，会被识别为负数。(异常)
	fmt.Println("\n=====================case-2=====================")
	fmt.Println("case-2：正数，首字节 < 0x7f，头部不会加上0x00，会被识别为负数。(异常)")
	case2 := "-129"
	result2 := DecimalValueFromStringMysql(case2)
	fmt.Printf("序列化过程->原始数据:%s ,序列化结果：%s \n", case2, result2)
	deserialize(result2, 2, 2)
	//case-3: 负数，转为补码后，与正数首字节 < 0x7f 的补码存在交集。（异常）这里以（1 、-255）
	fmt.Println("\n=====================case-3=====================")
	fmt.Println("case-3: 负数，转为补码后，与正数首字节 < 0x7f 的补码存在交集。（异常）这里以（0.01 、-2.55）为案例")
	case31 := "126"
	result31 := DecimalValueFromStringMysql(case31)
	fmt.Printf("序列化过程->原始数据:%s ,序列化结果：%s \n", case31, result31)
	deserialize(result31, 2, 2)
	fmt.Println("------------------------------------------------")
	case32 := "-130"
	result32 := DecimalValueFromStringMysql(case32)
	fmt.Printf("序列化过程->原始数据:%s ,序列化后结果：%s\n", case32, result32)
	deserialize(result32, 2, 2)
	fmt.Println("================================================")

}
