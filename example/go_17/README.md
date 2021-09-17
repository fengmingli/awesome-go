#泛型
Go语言的1.17版本发布了，其中开始正式支持泛型了。虽然还有一些限制，但是可以体验了：
```go
func print[T any] (arr []T) {
	for _, v := range arr {
		fmt.Print(v)
		fmt.Print(" ")
	}
	fmt.Println("")
}
func main() {
	strs := []string{"Hello", "World",  "Generics"}
	decs := []float64{3.14, 1.14, 1.618, 2.718 }
	nums := []int{2,4,6,8}

	print(strs)
	print(decs)
	print(nums)
}
```
要让这段程序跑起来需要在编译行上加上 -gcflags=-G=3编译参数（这个编译参数会在1.18版上成为默认参数）

```shell
go run --gcflags=-G=3 generic.go 
```