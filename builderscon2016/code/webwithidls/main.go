package main

import "net/http"

func main() {
	registerRootHandler(func(req rootRequest) (rootResponse, error) {
		res := rootResponse{}
		res.HogeId = req.HogeId
		return res, nil
	})
	http.ListenAndServe(":8080", nil)
}
