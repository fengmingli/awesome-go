/**
 * @Author LFM
 * @Date 2021/9/17 7:31 下午
 * @Since V1
 */

package main

import "fmt"

func print[T any] (arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}

func gMap[T1 any, T2 any] (arr []T1, f func(T1) T2) []T2 {
	result := make([]T2, len(arr))
	for i, elem := range arr {
		result[i] = f(elem)
	}
	return result
}

func gReduce[T1 any, T2 any] (arr []T1, init T2, f func(T2, T1) T2) T2 {
	result := init
	for _, elem := range arr {
		result = f(result, elem)
	}
	return result
}

func gFilter[T any] (arr []T, in bool, f func(T) bool) []T {
	result := []T{}
	for _, elem := range arr {
		choose := f(elem)
		if (in && choose) || (!in && !choose) {
			result = append(result, elem)
		}
	}
	return result
}
func gFilterIn[T any] (arr []T, f func(T) bool) []T {
	return gFilter(arr, true, f)
}
func gFilterOut[T any] (arr []T, f func(T) bool) []T {
	return gFilter(arr, false, f)
}

func main() {
	//strs := []string{"Hello", "World",  "Generics"}
	//decs := []float64{3.14, 1.14, 1.618, 2.718 }
	//nums := []int{2,4,6,8}
	//
	//print(strs)
	//print(decs)
	//print(nums)

	//nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//sum := gReduce(nums, 0, func(result, elem int) int {
	//	return result + elem
	//})
	//fmt.Printf("Sum = %d \n", sum)


	nums := []int {0,1,2,3,4,5,6,7,8,9}
	odds := gFilterIn(nums, func (elem int) bool  {
		return elem % 2 == 1
	})
	print(odds)
}
