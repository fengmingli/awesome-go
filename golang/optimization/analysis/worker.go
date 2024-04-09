package analysis

/**
 * @Author: LFM
 * @Date: 2023/8/19 23:11
 * @Since: 1.0.0
 * @Desc: TODO
 */

type Worker interface {
	Split()
	Send()
}

type WorkerManager struct {
	Worker Worker
}

var cache = make(map[string]Worker, 0)

func (dsm *WorkerManager) DynamicDataSourceFactory(worker string) Worker {
	return cache[worker]
}

func NewWorkerManager() *WorkerManager {
	return &WorkerManager{}
}

func init() {
	cache["A"] = NewA()
	cache["B"] = NewB()
}
