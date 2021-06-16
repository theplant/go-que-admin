package main

import (
	"context"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/theplant/go-que-admin/config"
	"github.com/tnclong/go-que"
)

func main() {
	q := config.TheQ

	worker, err := que.NewWorker(que.WorkerOptions{
		Queue: "import_pdf",
		Mutex: q.Mutex(),
		MaxLockPerSecond: 10,
		MaxBufferJobsCount: 0,
		MaxPerformPerSecond: 2,
		MaxConcurrentPerformCount: 1,
		Perform: func(ctx context.Context, job que.Job) error {
			plan := job.Plan()
			args := plan.Args
			fmt.Println("import pdf", string(args))
			return job.Done(ctx)
		},

	})

	if err != nil {
		panic(err)
	}

	err = worker.Run()
	if err != nil {
		panic(err)
	}

}
