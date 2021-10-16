/**
 * @Author LFM
 * @Date 2021/10/14 8:51 下午
 * @Since V1
 */

package analysis

import "math/rand"

func generate8191() {
	nums := make([]int, 8191) // < 64KB
	for i := 0; i < 8191; i++ {
		nums[i] = rand.Int()
	}
}

func generate8192() {
	nums := make([]int, 8192) // = 64KB
	for i := 0; i < 8192; i++ {
		nums[i] = rand.Int()
	}
}

func generate(n int) {
	nums := make([]int, n) // 不确定大小
	for i := 0; i < n; i++ {
		nums[i] = rand.Int()
	}
}

//go test -gcflags=-m .
func main() {
	generate8191()
	generate8192()
	generate(1)
}
