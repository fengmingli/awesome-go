/**
 * @Author LFM
 * @Date 2021/12/2 11:16 上午
 * @Since V1
 */

package typeconversion

import (
	"fmt"
	"testing"
)

func TestForced(t *testing.T) {
	Forced()
}
func TestConversions(t *testing.T) {
	Conversions()
}

func TestPassByValue(t *testing.T) {
	PassByValue()
}

func TestPassByQuote(t *testing.T) {
	PassByQuote()
}

func TestStruct(t *testing.T) {
	Struct()
}
func TestAssert(t *testing.T) {
	Assert()
}

func TestSetMain(t *testing.T) {
	SetMain()
}

//5 10
//2,2,6,5,4
//6,3,5,4,6
func TestSum(t *testing.T) {
	weight := []int{2, 2, 6, 5, 4}
	dd := []int{6, 3, 5, 4, 6}
	sum := SumTarget(10, weight)

	if len(sum) > 0 {
		for _, value := range sum {
			var targetSum int
			for _, v := range value {
				targetSum = targetSum + dd[v]
			}
			fmt.Println(targetSum)
		}
	}
}

func TestName(t *testing.T) {
	fmt.Println(CompareChar("[{(})]"))
	fmt.Println(CompareChar("[{(})]}"))
	fmt.Println(CompareChar("[({})]"))
}
