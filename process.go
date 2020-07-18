package batcher

import (
	"fmt"
	"time"

	"git.kanosolution.net/kano/dbflex"
	"git.kanosolution.net/kano/dbflex/orm"
	"github.com/ariefdarmawan/datahub"
	"github.com/eaciit/toolkit"
	"github.com/google/uuid"
)

const (
	ProcRunning string = "Running"
	ProcHold           = "Hold"
	ProcDone           = "Done"
	ProcFail           = "Fail"

	LogInfo    string = "INFO"
	LogWarning        = "WARNING"
	LogError          = "ERROR"
	LogFatal          = "FATAL"
)

type Process struct {
	orm.DataModelBase `json:"-" bson:"-"`
	ID                string    `json:"_id" bson:"_id"`
	Ref               string    `json:"ref"`
	Source            string    `json:"source"`
	Owner             string    `json:"owner"`
	Progress          int       `json:"progress"`
	Start             time.Time `json:"start"`
	End               time.Time `json:"end"`
	LastMsg           string    `json:"lastmsg"`
	Status            string    `json:"status"`
}

func (o *Process) TableName() string {
	return "procs"
}

func (o *Process) GetID(_ dbflex.IConnection) ([]string, []interface{}) {
	return []string{"_id"}, []interface{}{o.ID}
}

func (o *Process) SetID(keys ...interface{}) {
	o.ID = keys[0].(string)
}

func (o *Process) PreSave(c dbflex.IConnection) error {
	if o.ID == "" {
		o.ID = uuid.New().String()
	}
	return nil
}

func (o *Process) AddLog(h *datahub.Hub, logType string, msg string, data toolkit.M) error {
	l := new(ProcessLog)
	l.LogType = logType
	l.Message = msg
	if data == nil {
		data = toolkit.M{}
	}
	l.Data = data
	l.PID = o.ID
	l.TS = time.Now()

	if e := h.Save(l); e != nil {
		return fmt.Errorf("fail to save log. %s", e.Error())
	}

	return nil
}

func CreateProcess(h *datahub.Hub, source, ref, owner string, fn func(*Process) error) (string, error) {
	p := new(Process)
	p.Source = source
	p.Ref = ref
	p.Owner = owner
	p.Start = time.Now()
	p.Status = ProcRunning
	if e := h.Save(p); e != nil {
		return "", fmt.Errorf("fail to record process. %s", e.Error())
	}

	go func() {
		e := fn(p)
		pupd := new(Process)
		eget := h.GetByID(pupd, p.ID)
		if eget == nil {
			if e != nil {
				p.AddLog(h, LogError, e.Error(), nil)
				pupd.End = time.Now()
				pupd.Status = ProcFail
				pupd.LastMsg = e.Error()
				h.Save(pupd)
				return
			}

			p.AddLog(h, LogInfo, "Done", nil)
			pupd.End = time.Now()
			pupd.Status = ProcDone
			pupd.LastMsg = "Done"
			h.Save(pupd)
		}
	}()

	go p.AddLog(h, LogInfo, "Process has been started", nil)
	return p.ID, nil
}
