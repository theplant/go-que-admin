package models

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	queque "github.com/tnclong/go-que"
)

type GoqueJob struct {
	ID              uint
	Queue           string
	Args            Args
	RetryPolicy     RetryPolicy
	RunAt           time.Time
	DoneAt          *time.Time
	RetryCount      int
	ExpiredAt       *time.Time
	LastErrMsg      string
	LastErrStack    string
	UniqueID        string
	UniqueLifeCycle int
	UpdatedAt       time.Time
	CreatedAt       time.Time
}

type Args []interface{}

func (p Args) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *Args) Scan(data interface{}) error {
	var byteData []byte
	switch values := data.(type) {
	case []byte:
		byteData = values
	case string:
		byteData = []byte(values)
	default:
		return errors.New("scan DayParts unsupported type of data")
	}
	return json.Unmarshal(byteData, p)
}

type RetryPolicy struct {
	queque.RetryPolicy
}

func (p RetryPolicy) Value() (driver.Value, error) {
	return json.Marshal(p)
}

func (p *RetryPolicy) Scan(data interface{}) error {
	var byteData []byte
	switch values := data.(type) {
	case []byte:
		byteData = values
	case string:
		byteData = []byte(values)
	default:
		return errors.New("scan DayParts unsupported type of data")
	}
	return json.Unmarshal(byteData, p)
}
