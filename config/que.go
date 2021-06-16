package config

import (
	"database/sql"
	"github.com/tnclong/go-que"
	"github.com/tnclong/go-que/pg"
	"os"
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