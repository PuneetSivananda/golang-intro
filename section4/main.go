package main

import (
	"fmt"
	"go-api/controller"
	"log"
	"net/http"

	"go-api/model"
)

func main(){
	mux:= controller.Register()
	db:= model.Connect()
	fmt.Println("Serving the API...")
	defer db.Close()
	log.Fatal(http.ListenAndServe("localhost:3000", mux))
}