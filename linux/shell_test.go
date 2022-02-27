package linux

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	args := []string{"/Users/lifengming/Work/tmp/test.sh"}
	executor := CommandExecutor{}
	output, err := executor.ExecuteCommandWithOutput("/bin/sh", args...)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(output)
}
