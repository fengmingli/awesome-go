/**
 * @Author LFM
 * @Date 2021/12/3 11:22 上午
 * @Since V1
 */

package typeconversion

import (
	"fmt"
	"strings"
)

type People interface {
	Do()
}

type Man struct {
	Name string
}

func (man Man) Do() {
	fmt.Println(man.Name)
}

type Woman struct {
	Name string
}

func (woman *Woman) Do() {
	fmt.Println(woman.Name)
}

func SetMain() {
	//Man 值
	var man People = Man{Name: "男人"}
	man.Do()
	//
	var man2 People = &Man{Name: "男人2"}
	man2.Do()

	//Woman 引用
	var woman People = &Woman{}
	woman.Do()

	//var Woman2 People = Woman{}
	//Woman2.Do()
}

func Sum(target int, sourceArr []int) [][]int {
	var arr [][]int
	for j := 0; j < len(sourceArr); j++ {
		var dd = target
		var targetArr []int
		for i := j; i < len(sourceArr); i++ {
			dd = dd - sourceArr[i]
			if dd >= 0 {
				targetArr = append(targetArr, i)
				continue
			}
			arr = append(arr, targetArr)
		}
	}
	return arr
}

func SumTarget(target int, sourceArr []int) [][]int {
	var arr [][]int
	for i := 0; i < len(sourceArr); i++ {
		fmt.Println("->", i)
		var targetArr []int
		tempTarget := target
		for j := i; j < len(sourceArr); j++ {
			fmt.Println("<-->", j)
			tempTarget = tempTarget - sourceArr[j]
			fmt.Println("==", tempTarget)
			if tempTarget >= 0 {
				targetArr = append(targetArr, j)
				continue
			}
			tempTarget = tempTarget + sourceArr[j]
		}
		arr = append(arr, targetArr)
	}
	return arr
}

func CompareChar(str string) bool {
	if len(str)%2 == 1 {
		return false
	}
	var leftArr []string
	var rightArr []string
	strArr := []byte(str)
	for i := 0; i < len(strArr); i++ {
		if string(strArr[i]) == "(" || string(strArr[i]) == "{" || string(strArr[i]) == "[" {
			leftArr = append(leftArr, string(strArr[i]))
		}

		if string(strArr[i]) == ")" || string(strArr[i]) == "}" || string(strArr[i]) == "]" {
			rightArr = append(rightArr, string(strArr[i]))
		}
	}
	if len(leftArr) != len(rightArr) {
		return false
	}
	for i := 0; i < len(leftArr); i++ {
		if strings.Compare(leftArr[i], "(") == 0 {
			if rightArr[len(rightArr)-i] != ")" {
				return false
			}
		}
		if strings.Compare(leftArr[i], "{") == 0 {
			if rightArr[len(rightArr)-i] != "}" {
				return false
			}
		}
		if strings.Compare(leftArr[i], "[") == 0 {
			if rightArr[len(rightArr)-i] != "]" {
				return false
			}
		}

	}
	return true
}
