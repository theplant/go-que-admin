package config

import (
	"database/sql"
	"encoding/json"
	"os"

	"github.com/tnclong/go-que"
	"github.com/tnclong/go-que/pg"
)

var TheQ que.Queue

func init() {
	db, err := sql.Open("postgres", os.Getenv("DB_PARAMS"))
	if err != nil {
		panic(err)
	}
	TheQ, err = pg.New(db)
	if err != nil {
		panic(err)
	}
}

type Queue struct {
	Name string
	Args []*QueueArg
}

type QueueArg struct {
	Name string
	// string, int, time.Time ...
	Type string
}

var _queues []*Queue

func MustGetQueues() []*Queue {
	if _queues != nil {
		return _queues
	}

	s := os.Getenv("Queues")
	err := json.Unmarshal([]byte(s), &_queues)
	if err != nil {
		panic(err)
	}

	return _queues
}
