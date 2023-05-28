package timing

import (
	"fmt"
	"testing"
)

func TestMyTiming(t *testing.T) {
	MyTiming()
}
func TestMyTiming2(t *testing.T) {
	MyTiming2()
}

func TestMyTick(t *testing.T) {
	fmt.Println("=start=")
	MyTick()
	fmt.Println("=end=")
}
