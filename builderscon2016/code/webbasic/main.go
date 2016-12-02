package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}

type errResponse struct {
	Error string `json:"error"`
}

type rootRequest struct {
	HogeId int `json:"hoge_id"`
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	enc := json.NewEncoder(w)
	hogeIdStr := r.FormValue("hoge_id")
	hogeId, err := strconv.Atoi(hogeIdStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		enc.Encode(errResponse{err.Error()})
		return
	}
	req := rootRequest{}
	req.HogeId = hogeId
	err = enc.Encode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		enc.Encode(errResponse{err.Error()})
		return
	}
}
