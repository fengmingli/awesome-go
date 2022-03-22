/**
 * https://go.dev/ref/spec#Type_assertions
 * @Author LFM
 * @Date 2021/12/2 9:47 上午
 * @Since V1
 */

package typeconversion

import (
	"fmt"
	"reflect"
	"unsafe"
)

//Assert 断言
//https://go.dev/ref/spec#Type_assertions
func Assert() {
	var x interface{} = 7 // x has dynamic type int and value 7
	if i, ok := x.(int); ok {
		fmt.Println(i)
		fmt.Println(reflect.TypeOf(i))
	}

	if i, ok := x.(float64); ok {
		fmt.Println(i)
		fmt.Println(reflect.TypeOf(i))
	} else {
		fmt.Println("not float64")
	}
}

//Engine ---------------------------------------------------
type Engine struct {
}

var IRouter = &Engine{}

//Forced 强制
//https://go.dev/ref/spec#Package_unsafe
func Forced() {
	fmt.Println(IRouter)
	var engine = (*Engine)(nil)
	fmt.Println(&engine)
	//==================
	var f float64 = 1
	fmt.Println("f 类型", reflect.TypeOf(f))
	bits := *(*uint64)(unsafe.Pointer(&f))
	fmt.Println("bits 类型", reflect.TypeOf(bits))
}

//Conversions 显示======================
//https://go.dev/ref/spec#Conversions
func Conversions() {
	fmt.Println(float32(2.718281828)) // 2.718281828 of type float32
	fmt.Println(complex128(1))        // 1.0 + 0.0i of type complex128
	fmt.Println(float32(0.49999999))  // 0.5 of type float32
	fmt.Println(float32(0.49999899))  // 0.5 of type float32
	fmt.Println(float64(-1e-1000))    // 0.0 of type float64
	fmt.Println(string('x'))          // "x" of type string
	//fmt.Println(string(0x266c))       // "♬" of type string
	fmt.Println(string([]byte{'a'})) // not a constant: []byte{'a'} is not a constant
	fmt.Println((*int)(nil))         // not a constant: nil is not a constant, *int is not a boolean, numeric, or string type
	//fmt.Println(int(1.2))             // illegal: 1.2 cannot be represented as an int
	//fmt.Println(string(65.0))         // illegal: 65.0 is not an integer constant
}

//PassByValue 值传递
//形参地址 0xc00011c1e8
//实际参数的地址 0xc00011c1e0
//改动后的值是  1
func PassByValue() {
	var args int64 = 1
	modifiedNumber(args) // args就是实际参数
	fmt.Printf("实际参数的地址 %p\n", &args)
	fmt.Printf("改动后的值是  %d\n", args)
}

func modifiedNumber(args int64) { //这里定义的args就是形式参数
	fmt.Printf("形参地址 %p \n", &args)
	args = 10
}

func PassByQuote() {
	var args int64 = 1
	addr := &args
	fmt.Printf("原始指针的内存地址是 %p\n", addr)
	fmt.Printf("指针变量addr存放的地址 %p\n", &addr)
	modifiedNumberV2(addr) // args就是实际参数
	fmt.Printf("改动后的值是  %d\n", args)
}

func modifiedNumberV2(addr *int64) { //这里定义的args就是形式参数
	fmt.Printf("形参地址 %p \n", &addr)
	fmt.Printf("形参值 %p \n", addr)
	*addr = 10
}

type Person struct {
	Name string
	Age  int64
}

//Struct -----------------------------------------------------------------------
func Struct() {
	per := Person{
		Name: "张三",
		Age:  int64(8),
	}
	fmt.Printf("原始struct地址是：%p\n", &per)
	modifiedAge(per)
	fmt.Println(per)
	modifiedAgeV2(&per)
	fmt.Println(per)
}
func modifiedAge(per Person) {
	fmt.Printf("V1函数里接收到struct的内存地址是：%p\n", &per)
	per.Age = 10
}
func modifiedAgeV2(per *Person) {
	fmt.Printf("V2函数里接收到struct的内存地址是：%p\n", &per)
	per.Age = 10
}

//原始struct地址是：0xc0000bc048
//函数里接收到struct的内存地址是：0xc0000bc060
//{张三 8}
//Struct -----------------------------------------------------------------------

//传值的意思是：函数传递的总是原来这个东西的一个副本，一副拷贝。
//go就是值传递，可以确认的是Go语言中所有的传参都是值传递（传值），都是一个副本，一个拷贝。
//因为拷贝的内容有时候是非引用类型（int、string、struct等这些），这样就在函数中就无法修改原内容数据；
//有的是引用类型（指针、map、slice、chan等这些），这样就可以修改原内容数据。

//是否可以修改原内容数据，和传值、传引用没有必然的关系。
