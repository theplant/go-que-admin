package models

import "time"

type GoqueJob struct {
	ID        uint
	Queue string
	Args     string
	RetryPolicy string
	RunAt time.Time
	DoneAt time.Time
	RetryCount int
	ExpiredAt time.Time
	LastErrMsg string
	LastErrStack string
	UniqueID string
	UniqueLifeCycle int
	UpdatedAt time.Time
	CreatedAt time.Time
}
