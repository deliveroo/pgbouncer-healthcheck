package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
)

type requestHandlerWithContext func(context.Context, http.ResponseWriter, httprouter.Params) error
type requestHandler func(http.ResponseWriter, httprouter.Params) error

func makeRequestHandlerWithContext(h requestHandlerWithContext) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
		defer cancel()
		err := h(ctx, w, ps)
		if err != nil {
			errMsg := fmt.Sprintf("An error occured: %s", err)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
	}
}

func makeRequestHandler(h requestHandler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		err := h(w, ps)
		if err != nil {
			errMsg := fmt.Sprintf("An error occured: %s", err)
			http.Error(w, errMsg, http.StatusInternalServerError)
			return
		}
	}
}

func makeRequestHandlerCommand(desc string, cmdName string, cmdArg ...string) httprouter.Handle {
	return makeRequestHandler(func(w http.ResponseWriter, ps httprouter.Params) error {
		cmd := exec.Command(cmdName, cmdArg...)
		output, err := cmd.Output()
		if err != nil {
			return errors.Wrapf(err, "Error fetching %s", desc)
		}
		w.Write(output)
		return nil
	})
}

func makeRequestHandlerFile(desc string, path string) httprouter.Handle {
	return makeRequestHandler(func(w http.ResponseWriter, ps httprouter.Params) error {
		output, err := ioutil.ReadFile(path)
		if err != nil {
			return errors.Wrapf(err, "Error fetching %s", desc)
		}
		w.Write(output)
		return nil
	})
}

func returnJSON(w http.ResponseWriter, item interface{}) error {

	resp, err := json.Marshal(item)
	if err != nil {
		return errors.Wrap(err, "Error converting response to JSON")
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(resp)
	return nil
}
