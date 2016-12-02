package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

var _ = strconv.IntSize

type nameRequest struct {
	Name string
}
type nameResponse struct {
	Name string `json:"name"`
}
type nameHandler func(nameRequest) (nameResponse, error)

func generateNameHandler(h nameHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		nameStr := r.FormValue("hoge_id")
		name := nameStr // <- !?
		req := nameRequest{name}
		res, _ := h(req)
		enc.Encode(res)
	}
}
func registerNameHandler(h nameHandler) {
	http.HandleFunc("/name", generateNameHandler(h))
}
