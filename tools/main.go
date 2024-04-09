package main

import (
	"encoding/base64"
	"fmt"
	"math/big"
)

func main() {

	//case1 := "-96545.13"
	//fmt.Println("====case-1 ===")
	//fmt.Printf("原始数据：%v ,dtle序列化后：%v \n", case1, DecimalValueFromStringMysql(case1))
	//fmt.Printf("原始数据：%v ,反序列化后：%v\n", DecimalValueFromStringMysql(case1), decodeBigDecimal(DecimalValueFromStringMysql(case1)))
	//
	//fmt.Println("")
	//case2 := "123.32"
	//fmt.Println("====case-2===")
	//fmt.Printf("原始数据：%v ,dtle序列化后：%v \n", case2, DecimalValueFromStringMysql(case2))
	//fmt.Printf("原始数据：%v ,反序列化后：%v\n", DecimalValueFromStringMysql(case2), decodeBigDecimal(DecimalValueFromStringMysql(case2)))
	//
	//fmt.Println("")
	//case3 := "660.23"
	//fmt.Println("====case-3 ===")
	//fmt.Printf("原始数据：%v ,dtle序列化后：%v \n", case3, DecimalValueFromStringMysql(case3))
	//fmt.Printf("原始数据：%v ,反序列化后：%v\n", DecimalValueFromStringMysql(case3), decodeBigDecimal(DecimalValueFromStringMysql(case3)))
	//
	fmt.Println("")
	case4 := "2"
	fmt.Println("====case-4 ===")
	fmt.Printf("原始数据：%v ,dtle序列化后：%v \n", case4, DecimalValueFromStringMysql(case4))
	fmt.Printf("原始数据：%v ,反序列化后：%v\n", DecimalValueFromStringMysql(case4), decodeBigDecimal(DecimalValueFromStringMysql(case4)))

	//fmt.Println("")
	case5 := "-254"
	fmt.Println("====case-5 ===")
	fmt.Printf("原始数据：%v ,dtle序列化后：%v \n", case5, DecimalValueFromStringMysql(case5))
	fmt.Printf("原始数据：%v ,反序列化后：%v\n", DecimalValueFromStringMysql(case5), decodeBigDecimal2(DecimalValueFromStringMysql(case5)))

}

func decodeBigDecimal(encoded string) *big.Float {
	decodedBytes, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Errorf("Decode fail. message:%v ", err.Error())
		return nil
	}

	//newValue := float64(0)
	//for i := 0; i < len(decodedBytes); i++ {
	// v := decodedBytes[i]
	// offset := (len(decodedBytes) - i - 1) * 8
	// newValue = newValue + float64(v << offset)
	//}
	//
	//firstByte := decodedBytes[0]
	//
	//if firstByte&0x70 != 0 {
	// newValue = newValue - math.Pow(0x100, float64(len(decodedBytes)))
	//}
	//
	//fmt.Sprintf("%f",newValue)
	fmt.Printf("%b\n", decodedBytes)
	// Create a big.Int from the byte slice
	bigInt := new(big.Int)

	if decodedBytes[0]&0x70 != 0 {
		//bigInt-=int64('1'+'00'*len(decodedBytes),16)
	}
	bigInt.SetBytes(decodedBytes)
	// Create a big.Float with the specified scale
	bigFloat := new(big.Float).SetPrec(64) // Set the desired precision
	bigFloat.SetInt(bigInt)
	scaleFactor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(2)), nil))
	bigFloat.Quo(bigFloat, scaleFactor)

	return bigFloat

	//return newValue

}

func decodeBigDecimal2(bigDecimal string) *big.Float {
	byteVals, _ := base64.StdEncoding.DecodeString(bigDecimal)
	bval := ""
	for _, c := range byteVals {
		bval += fmt.Sprintf("%08b", c)
	}
	intval := new(big.Float)
	intval, success := intval.SetString(bval)
	if !success {
		return nil
	}
	if byteVals[0]&0x70 != 0 {
		negValue := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(16), big.NewInt(int64(len(byteVals)*2)), nil))
		intval.Sub(intval, negValue)
	}
	scaleFactor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(2)), nil))
	result := new(big.Float).Quo(intval, scaleFactor)
	return result
}

//func decodeBigDecimal3(bigDecimal string, scale int, precision int) *big.Float {
//	byteVals, err := base64.StdEncoding.DecodeString(bigDecimal)
//	if err != nil {
//		return nil
//	}
//
//	bval := ""
//	for _, c := range byteVals {
//		bval += fmt.Sprintf("%08b", c)
//	}
//
//	intval := new(big.Int)
//	intval, success := intval.SetString(bval, 2)
//	if !success {
//		return nil
//	}
//
//	if byteVals[0]&0x70 != 0 {
//		negValue := new(big.Int).SetBytes([]byte("1" + strings.Repeat("00", len(byteVals))))
//		intval.Sub(intval, negValue)
//	}
//
//	scaleFactor := new(big.Float).SetInt(new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(2)), nil))
//	result := new(big.Float).Quo(intval, scaleFactor)
//	return result
//}

// value: e.g. decimal(11,5), 123.45 will be 123.45000
func DecimalValueFromStringMysql(value string) string {
	//初始化一个big的Int型
	sum := big.NewInt(0)

	//如果为负数，将-号去掉
	//绝对值
	isNeg := false
	if value[0] == '-' {
		value = value[1:]
		isNeg = true
	}

	//转变为整数 例如：96646.13 -> 9654513
	for i := range value {
		//字符'0' 48 ~ 字符'9' 57
		if '0' <= value[i] && value[i] <= '9' {
			//乘以10
			sum.Mul(sum, decimalNums[10])
			offset := value[i] - '0'
			if offset != 0 { // add 0 = do nothing
				sum.Add(sum, decimalNums[offset])
			}
		}
	}

	if isNeg {
		sum.Neg(sum) // 如果是负数，将其取反
	}

	bs := sum.Bytes()

	// 转为字节
	if len(bs) == 0 {
		bs = []byte{0}
	}

	// 如果是负数，将最高位设置为1
	if isNeg {
		bs[0] |= 0x80
	}

	return base64.StdEncoding.EncodeToString(bs)

	//bs := sum.Bytes()
	////fmt.Printf("%b\n len（%v）", bs, len(bs))
	////转为字节
	//if len(bs) == 0 {
	//	bs = []byte{0}
	//}
	//
	//if isNeg {
	//	for i := len(bs) - 1; i >= 0; i-- {
	//		bs[i] = ^bs[i] //正数取反
	//	}
	//
	//	for i := len(bs) - 1; i >= 0; i-- {
	//		bs[i] += 1 //补码
	//		if bs[i] != 0x00 {
	//			break
	//		}
	//	}
	//
	//
	//} else if bs[0] > 0x7f { //127
	//	bs2 := make([]byte, len(bs)+1)
	//	bs2[0] = 0x00
	//	copy(bs2[1:], bs)
	//	bs = bs2
	//}
	//
	//return base64.StdEncoding.EncodeToString(bs)
}

var (
	decimalNums [11]*big.Int
)

func init() {
	for i := 0; i <= 10; i++ {
		decimalNums[i] = big.NewInt(int64(i))
	}
}
