package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"../views"
)


func ping() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			data := views.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			fmt.Println("Request Recieved")
			json.NewEncoder(w).Encode(data)
		}
	}
}