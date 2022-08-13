// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type workerTask func(id string) error

type Worker struct {
	id       int
	task     workerTask
	stopChan chan struct{}
	dataChan chan string
}

type WorkerGroup struct {
	workers  []*Worker
	stopChan chan struct{}
	dataChan chan string
}

func (w *Worker) Run() error {
	defer wg.Done()
	for {
		select {
		case <-w.stopChan:
			fmt.Println("stopping this worker: ", w.id)
			return nil
		case inData := <-w.dataChan:
			fmt.Println("simulating some job: ", w.id, inData)
			w.task(inData)
			fmt.Println("completed: ", w.id)
		default:
			//fmt.Println("no job for now:", w.id)
		}
	}
}

func InitWorker(task workerTask) *WorkerGroup {
	workers := make([]*Worker, 5)
	stop := make(chan struct{})
	data := make(chan string, 5)
	for i := 0; i < 5; i++ {
		workers[i] = &Worker{
			id:       i,
			task:     task,
			stopChan: stop,
			dataChan: data,
		}
	}

	return &WorkerGroup{
		workers:  workers,
		dataChan: data,
		stopChan: stop,
	}
}

func (group *WorkerGroup) StartAll() {
	for idx, worker := range group.workers {
		wg.Add(1)
		fmt.Println("starting worker: ", idx)
		go worker.Run()
	}
}

func (group *WorkerGroup) StopAll() {
	for range group.workers {
		group.stopChan <- struct{}{}
	}
}

func main() {
	var job workerTask
	job = func(inData string) error {
		time.Sleep(2 * time.Second)
		fmt.Printf("finished some stuff: %v\n", inData)
		return nil
	}
	workerGroup := InitWorker(job)
	workerGroup.StartAll()

	inputString := "abcdefghijklmnopqrstuvwxyz"

	for i := 0; i < 20; i++ {
		fmt.Printf("sending %v to data channel\n", string(inputString[i]))
		workerGroup.dataChan <- string(inputString[i])
	}

	fmt.Println(">>>>>>>>>>>>>>>>>>>>>>>>>>>closing channel<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<<")
	workerGroup.StopAll()

	wg.Wait()
	fmt.Println("closed all workers")
}
