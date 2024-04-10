package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	var v any
	var s int
	var err error

	switch r.Method {
	case http.MethodGet:
		v, s, err = getRoot(r)
	default:
		err = fmt.Errorf("method not allowed: %s", r.Method)
	}

	if err != nil {
		http.Error(w, err.Error(), s)
	} else {
		json.NewEncoder(w).Encode(&v)
	}

}

func getRoot(r *http.Request) (v any, s int, err error) {
	s = http.StatusOK

	err = json.NewDecoder(r.Body).Decode(&v)
	if err != nil {
		s = http.StatusInternalServerError
	}

	return v, s, err
}
