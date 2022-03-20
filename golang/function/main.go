package main

import "fmt"

type customizeInterface struct {
	MyFunc func(string) bool
}

func main() {
	c := &customizeInterface{
		MyFunc: func(s string) bool {
			if s == "s" {
				return true
			}
			return false
		},
	}
	myFunc := c.MyFunc("ss")
	fmt.Println(myFunc)
}

// A Getter loads data for a key.
type Getter interface {
	Get(key string) ([]byte, error)
}

// A GetterFunc implements Getter with a function.
type GetterFunc func(key string) ([]byte, error)

// Get implements Getter interface function
func (f GetterFunc) Get(key string) ([]byte, error) {
	return f(key)
}
