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

type MetadataWorker struct {
}

func NewMetadataWorker() fx.Worker {
	return MetadataWorker{}

}

func (a MetadataWorker) Split() {
	fmt.Println("A-Split")
}

func (a MetadataWorker) Send() {
	fmt.Println("A-Send")
}
