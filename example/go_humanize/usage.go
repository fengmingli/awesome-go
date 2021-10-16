/**
 * @Author LFM
 * @Date 2021/10/11 4:22 下午
 * @Since V1
 */

package go_humanize

import (
	"fmt"
	"github.com/dustin/go-humanize"
	"time"
)

func Usage() {
	fmt.Printf("That file is %s.", humanize.Bytes(82854982000)) // That file is 83 MB.

	// This was touched 7 hours ago.
	fmt.Printf("This was touched %s.", humanize.Time(time.Now().Add(45*time.Minute)))

	fmt.Printf("You owe ￥%s.\n", humanize.Comma(6582491))

}

func Closure() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
}
