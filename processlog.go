package batcher

import (
	"fmt"
	"time"

	"git.kanosolution.net/kano/dbflex"
	"git.kanosolution.net/kano/dbflex/orm"
	"github.com/eaciit/toolkit"
	"github.com/google/uuid"
)

type ProcessLog struct {
	orm.DataModelBase `json:"-" bson:"-"`
	ID                string    `json:"_id" bson:"_id"`
	PID               string    `json:"pid"`
	TS                time.Time `json:"ts"`
	LogType           string    `json:"logtype"`
	Message           string    `json:"message"`
	Data              toolkit.M `json:"data"`
}

func (o *ProcessLog) TableName() string {
	return "proclogs"
}

func (o *ProcessLog) GetID(_ dbflex.IConnection) ([]string, []interface{}) {
	return []string{"_id"}, []interface{}{o.ID}
}

func (o *ProcessLog) SetID(keys ...interface{}) {
	o.ID = keys[0].(string)
}

func (o *ProcessLog) PreSave(c dbflex.IConnection) error {
	if o.PID == "" {
		return fmt.Errorf("PID is mandatory")
	}

	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}
