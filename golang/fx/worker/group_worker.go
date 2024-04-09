package worker

import (
	"awesome-go/golang/fx"
	"fmt"
)

/**
 * @Author: LFM
 * @Date: 2023/8/19 23:44
 * @Since: 1.0.0
 * @Desc: TODO
 */

type GroupWorker struct {
}

func NewGroupWorker() fx.Worker {
	return GroupWorker{}

}

func (a GroupWorker) Split() {
	fmt.Println("A-Split")
}

func (a GroupWorker) Send() {
	fmt.Println("A-Send")
}
