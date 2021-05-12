package controller

import (
	"go-api/model"
	"net/http"
)


func create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodGet {
			//fetch data from request
			if err:= model.CreateTODO(); err != nil{
				w.Write([]byte ("Some Error"))
				return 
			}
		}
	}
}