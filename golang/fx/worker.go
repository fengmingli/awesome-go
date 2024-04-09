package fx

import "awesome-go/golang/fx/worker"

/**
 * @Author: LFM
 * @Date: 2023/8/19 23:44
 * @Since: 1.0.0
 * @Desc: TODO
 */

type WorkerType string

const (
	GroupWorker    WorkerType = "GroupWorker"
	MetadataWorker WorkerType = "MetadataWorker"
)

type Worker interface {
	Split()
}

type WorkerManager struct{}

func NewWorkerManager(
	metadataWorker worker.MetadataWorker,
	groupWorker worker.GroupWorker,
) *WorkerManager {
	WorkerService[GroupWorker] = groupWorker
	WorkerService[MetadataWorker] = metadataWorker
	return &WorkerManager{}
}

var WorkerService = make(map[WorkerType]Worker, 0)

func (w *WorkerManager) DynamicDataSourceFactory(wt WorkerType) Worker {
	return WorkerService[wt]
}
