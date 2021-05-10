package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"./structs"
)

func main(){
	mux:=http.NewServeMux()

	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet{
			data:= structs.Response{
				Code: http.StatusOK,
				Body: "pong",
			}
			fmt.Println("Request Recieved")	
			json.NewEncoder(w).Encode(data)
		}
	})
	
	http.ListenAndServe("localhost:3000", mux)
}