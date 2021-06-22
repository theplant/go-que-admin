package main

import (
	"context"
	"fmt"
	"time"

	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/theplant/go-que-admin/config"
	"github.com/tnclong/go-que"
)

func main() {
	q := config.TheQ

	worker, err := que.NewWorker(que.WorkerOptions{
		Queue:                     "import_pdf",
		Mutex:                     q.Mutex(),
		MaxLockPerSecond:          10,
		MaxBufferJobsCount:        0,
		MaxPerformPerSecond:       2,
		MaxConcurrentPerformCount: 1,
		Perform: func(ctx context.Context, job que.Job) error {
			plan := job.Plan()
			args := plan.Args
			fmt.Println("import pdf", string(args))
			if string(args) == "[3, 2, 1]" {
				fmt.Println("import pdf job done", string(args))
				return job.Done(ctx)
			}
			panic("wrong args")
		},
	})

	if err != nil {
		panic(err)
	}

	go func() {
		err := worker.Run()
		if err != nil {
			panic(err)
		}
	}()

	_, err = q.Enqueue(context.Background(), nil, que.Plan{
		Queue: "import_pdf",
		Args:  que.Args(1, 2, 3),
		RunAt: time.Now(),
		RetryPolicy: que.RetryPolicy{
			InitialInterval:        1 * time.Second,
			NextIntervalMultiplier: 1,
			IntervalRandomPercent:  0,
			MaxRetryCount:          30,
		},
	})
	if err != nil {
		panic(err)
	}

	select {}
}
