/**
 * @Author LFM
 * @Date 2021/12/9 10:31 上午
 * @Since V1
 */

package _interface

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	animal := Animal{"中华田园犬"}
	dog := Dog{animal}
	fmt.Println(dog.GetName())
	fmt.Println(dog.Call())
	fmt.Println(dog.FavorFood())
}
