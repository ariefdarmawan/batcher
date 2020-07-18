package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"git.kanosolution.net/kano/appkit"
	"git.kanosolution.net/kano/dbflex"
	"github.com/ariefdarmawan/batcher"
	"github.com/ariefdarmawan/byter"
	"github.com/eaciit/toolkit"
)

var (
	btr = byter.NewByter("")
)

func handlingMux(mux *http.ServeMux) {
	mux.HandleFunc("/v1/process/gets", getProcess)
	mux.HandleFunc("/v1/process/add", addProcess)
}

func getProcess(w http.ResponseWriter, r *http.Request) {
	ps := []batcher.Process{}
	e := h.Gets(new(batcher.Process), dbflex.NewQueryParam(), &ps)
	if e != nil {
		writeError(w, e)
		return
	}

	writeJson(w, ps)
}

func addProcess(w http.ResponseWriter, r *http.Request) {
	addPayload := struct {
		S int
	}{}

	e := getPayload(r, &addPayload)
	if e != nil {
		writeError(w, e)
		return
	}
	if addPayload.S == 0 {
		addPayload.S = toolkit.RandInt(10) + 2
	}

	id, e := batcher.CreateProcess(h, "fake.process", appkit.MakeID(fmt.Sprintf("%d.", addPayload.S), 20), "fake-user", func(*batcher.Process) error {
		time.Sleep(time.Duration(addPayload.S) * time.Second)
		//time.Sleep(5 * time.Second)
		return nil
	})
	if e != nil {
		writeError(w, e)
	}

	writeJson(w, toolkit.M{}.Set("id", id))
}

func getPayload(r *http.Request, target interface{}) error {
	defer r.Body.Close()
	bs, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = btr.DecodeTo(bs, target, nil)
	if err != nil {
		return err
	}
	return nil
}

func writeError(w http.ResponseWriter, e error) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(e.Error()))
}

func writeJson(w http.ResponseWriter, any interface{}) {
	bs, e := btr.Encode(any)
	if e != nil {
		writeError(w, fmt.Errorf("fail to encode result. %s", e.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(bs)
}
