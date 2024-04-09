package fx

import (
	"awesome-go/golang/fx/worker"
	"go.uber.org/fx"
	"testing"
)

/**
 * @Author: LFM
 * @Date: 2023/8/19 23:54
 * @Since: 1.0.0
 * @Desc: TODO
 */

func TestFx(t *testing.T) {
	app := fx.New(
		fx.Provide(NewWorkerManager),           // 提供Greeter组件
		fx.Provide(worker.NewMetadataWorker()), // 提供Printer组件
		fx.Provide(worker.NewGroupWorker()),
		// 调用Printer的Print方法
	)

	app.Run()

}
