package main

import (
	"go-api/controller"
	"log"
	"net/http"

	"go-api/model"
)

func main(){
	mux:= controller.Register()
	db:= model.Connect()
	defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}