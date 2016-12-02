package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type rootRequest struct {
	HogeId int
}
type rootResponse struct {
	HogeId int `json:"hoge_id"`
}
type rootHandler func(rootRequest) (rootResponse, error)

func generateRootHandler(h rootHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		hogeIdStr := r.FormValue("hoge_id")
		hogeId, _ := strconv.Atoi(hogeIdStr)
		req := rootRequest{hogeId}
		res, _ := h(req)
		enc.Encode(res)
	}
}
func registerRootHandler(h rootHandler) {
	http.HandleFunc("/", generateRootHandler(h))
}
