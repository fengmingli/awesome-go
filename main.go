package main

import (
	"fmt"
	"sync"
	"time"
)

// JobStatus 定义工作状态类型
type JobStatus string

const (
	Pending   JobStatus = "Pending"
	Running   JobStatus = "Running"
	Completed JobStatus = "Completed"
	Failed    JobStatus = "Failed"
)

// Job 定义工作接口
type Job interface {
	ID() int
	Status() JobStatus
	Do() error
}

// CheckJob 实现 Job 接口
type CheckJob struct {
	id     int
	status JobStatus
}

func (cj *CheckJob) ID() int {
	return cj.id
}

func (cj *CheckJob) Status() JobStatus {
	return cj.status
}

func (cj *CheckJob) Do() error {
	fmt.Printf("Doing CheckJob %d\n", cj.id)
	time.Sleep(2 * time.Second) // Simulating work
	cj.status = Completed
	return nil
}

// Worker 定义工作者接口
type Worker interface {
	Start()
	Stop()
	IsIdle() bool
	CompareID(id int) bool
	AwaitCompletion()
}

// TaskWorker 实现 Worker 接口
type TaskWorker struct {
	id          int
	jobChannel  <-chan Job
	stopChan    chan struct{}
	currentJob  Job
	currentDone chan struct{}
	update      chan func()
}

// NewTaskWorker 创建一个新的 TaskWorker 实例
func NewTaskWorker(id int, jobChannel <-chan Job) *TaskWorker {
	return &TaskWorker{
		id:          id,
		jobChannel:  jobChannel,
		stopChan:    make(chan struct{}),
		currentDone: make(chan struct{}),
		update:      make(chan func()),
	}
}

// Start 启动 TaskWorker
func (w *TaskWorker) Start() {
	fmt.Printf("Worker %d starts\n", w.id)
	go func() {
		for {
			select {
			case job := <-w.jobChannel:
				fmt.Printf("Worker %d starts processing Job %d\n", w.id, job.ID())
				w.update <- func() { w.currentJob = job }
				err := job.Do()
				if err != nil {
					fmt.Printf("Worker %d encountered an error while processing Job %d: %v\n", w.id, job.ID(), err)
				}
				fmt.Printf("Worker %d finished processing Job %d\n", w.id, job.ID())
				w.update <- func() { w.currentJob = nil }
				w.currentDone <- struct{}{}
			case <-w.stopChan:
				fmt.Printf("Worker %d stops\n", w.id)
				return
			}
		}
	}()
}

// Stop 停止 TaskWorker
func (w *TaskWorker) Stop() {
	w.stopChan <- struct{}{}
}

// IsIdle 检查工作者是否空闲
func (w *TaskWorker) IsIdle() bool {
	return w.currentJob == nil
}

// CompareID 比较工作者当前处理的任务 ID 与给定 ID 是否匹配
func (w *TaskWorker) CompareID(id int) bool {
	if w.currentJob != nil {
		return w.currentJob.ID() == id
	}
	return false
}

// AwaitCompletion 等待当前任务完成
func (w *TaskWorker) AwaitCompletion() {
	<-w.currentDone
}

// Update 更新工作者状态
func (w *TaskWorker) Update(fn func()) {
	w.update <- fn
}

// Scheduler 定义调度器
type Scheduler struct {
	workers    []*TaskWorker
	jobChannel chan Job
	wg         sync.WaitGroup
}

func NewScheduler(numWorkers int, jobChannel chan Job) *Scheduler {
	scheduler := &Scheduler{
		jobChannel: jobChannel,
	}
	for i := 0; i < numWorkers; i++ {
		worker := NewTaskWorker(i+1, jobChannel)
		scheduler.workers = append(scheduler.workers, worker)
	}
	return scheduler
}

func (s *Scheduler) Start() {
	for _, worker := range s.workers {
		worker.Start()
	}
}

func (s *Scheduler) Stop() {
	for _, worker := range s.workers {
		worker.Stop()
	}
}

func (s *Scheduler) AddJob(job Job) {
	s.wg.Add(1)
	go func() {
		defer s.wg.Done()
		s.jobChannel <- job
	}()
}

func (s *Scheduler) Wait() {
	s.wg.Wait()
}

func main() {
	jobChannel := make(chan Job)
	scheduler := NewScheduler(5, jobChannel)
	scheduler.Start()

	// Simulating job generation
	for i := 1; i <= 10; i++ {
		job := &CheckJob{id: i, status: Pending}
		scheduler.AddJob(job)
	}

	scheduler.Wait()
	scheduler.Stop()
	close(jobChannel)
}
