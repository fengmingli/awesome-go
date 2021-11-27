/**
 * @Author LFM
 * @Date 2021/10/26 2:55 下午
 * @Since V1
 */

package copy

import (
	"fmt"
	"strconv"
)

func Copy(a []int, i int) {
	fmt.Println(a)
	// 删除a[i]，可以用 copy 将i+1到末尾的值覆盖到i,然后末尾-1
	fmt.Println(a[i:])
	fmt.Println(a[i+2:])
	copy(a[i:], a[i+2:])
	fmt.Println(a[i:])
	fmt.Println(a[i+1:])
	a = a[:len(a)-1]
	fmt.Println(a)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{5, 4, 3}
	fmt.Println(slice1)
	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Println(slice2)
	copy(slice1, slice2) // 只会复制slice2的3个元素到slice1的前3个位置
	fmt.Println(slice1)
}

func Common() {
	// byte转数字
	s := "12345" // s[0] 类型是byte
	num := int(s[0] - '0')
	str := string(s[0])
	b := byte(num + '0')                // '1'
	fmt.Printf("%d%s%c\n", num, str, b) // 111

	// 字符串转数字
	num1, _ := strconv.Atoi("19")
	str1 := strconv.Itoa(20)
	fmt.Printf("%d%s\n", num1, str1) // 111
}

func subsets(nums []int) [][]int {
	// 保存最终结果
	result := make([][]int, 0)
	// 保存中间结果
	list := make([]int, 0)
	backtrack(nums, 0, list, &result)
	return result
}

// nums 给定的集合
// pos 下次添加到集合中的元素位置索引
// list 临时结果集合(每次需要复制保存)
// result 最终结果
func backtrack(nums []int, pos int, list []int, result *[][]int) {
	// 把临时结果复制出来保存到最终结果
	ans := make([]int, len(list))
	copy(ans, list)
	*result = append(*result, ans)
	// 选择、处理结果、再撤销选择
	for i := pos; i < len(nums); i++ {
		list = append(list, nums[i])
		backtrack(nums, i+1, list, result)
		list = list[0 : len(list)-1]
	}
}
