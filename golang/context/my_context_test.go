package context

import (
	"fmt"
	"os"
	"testing"
)

func TestMyContext(t *testing.T) {
	getenv := os.Setenv("LFM", "123")

	fmt.Println(getenv)
	lookupEnv, ok := os.LookupEnv("HOME_AGE")
	if ok {
		fmt.Println(lookupEnv)
	}

	MyContext()
}
