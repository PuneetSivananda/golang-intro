// API Application Programming interface;
// MVC Model View Controller

package main

import (
	"net/http"
)

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		w.Write([]byte("Hello World"))
	})
	http.ListenAndServe("localhost:3000", mux)
}