package worker

import (
	"fmt"
	"github.com/jasonlvhit/gocron"
)

type Worker struct {
	*gocron.Scheduler
	stopped bool
}

func New() *Worker {
	newScheduler := gocron.NewScheduler()
	return &Worker{
		Scheduler: newScheduler,
		stopped:   true,
	}
}

func (w *Worker) StartWorker() error {
	if !w.stopped {
		return fmt.Errorf("worker already started")
	}

	w.Start()
	w.stopped = false

	return nil
}
