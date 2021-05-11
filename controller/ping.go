package controller

import (
	"../view"
	"encoding/json"
	"net/http"
)

func ping() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := view.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			json.NewEncoder(w).Encode(data)
		}
	}
}
