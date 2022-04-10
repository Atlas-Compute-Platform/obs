package main

import (
	"io"
	"net/http"
	"os"

	"github.com/Atlas-Compute-Platform/lib"
)

func apiLoad(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		objId string = r.URL.Query().Get(lib.KEY_ID)
		obj   []byte
		err   error
	)

	if obj, err = os.ReadFile(objId); err != nil {
		lib.LogError(os.Stderr, "main.apiStore", err)
		http.NotFound(w, r)
		return
	}

	if _, err = w.Write(obj); err != nil {
		lib.LogError(os.Stderr, "main.apiStore", err)
	}
}

func apiStore(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		objId string = r.URL.Query().Get(lib.KEY_ID)
		obj   []byte
		err   error
	)

	if obj, err = io.ReadAll(r.Body); err != nil {
		lib.LogError(os.Stderr, "main.apiStore", err)
		return
	}

	if err = os.WriteFile(objId, obj, lib.FD_MODE); err != nil {
		lib.LogError(os.Stderr, "main.apiStore", err)
	}
}

func apiRemove(w http.ResponseWriter, r *http.Request) {
	lib.SetCors(&w)
	var (
		objId string = r.URL.Query().Get(lib.KEY_ID)
		err   error
	)

	if err = os.Remove(objId); err != nil {
		lib.LogError(os.Stderr, "main.apiStore", err)
		http.NotFound(w, r)
	}
}
