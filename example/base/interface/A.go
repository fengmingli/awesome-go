/**
 * @Author LFM
 * @Date 2021/12/9 10:18 上午
 * @Since V1
 */

package _interface

import "fmt"

//Animal 类型于Java里面的继承
type Animal struct {
	Name string
}

func (a Animal) Call() string {
	return "动物的叫声..."
}

func (a Animal) FavorFood() string {
	return "爱吃的食物..."
}

func (a Animal) GetName() string {
	return a.Name
}

type Dog struct {
	Animal
}

//FavorFood 多态
func (d Dog) FavorFood() string {
	fmt.Println(d.Animal.FavorFood())
	return "骨头"
}
func (d Dog) Call() string {
	fmt.Println(d.Animal.Call())
	return "汪汪汪"
}
