package main

import (
	"net/http"

	"./controller"
)

func main(){
	mux:= controller.Register()
	http.ListenAndServe("localhost:3000", mux)
}