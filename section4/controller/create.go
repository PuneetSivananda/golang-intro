package controller

import (
	"encoding/json"
	"fmt"
	"go-api/model"
	"go-api/views"
	"net/http"
)


func create() http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		if r.Method == http.MethodPost {  
			//fetch data from request
			data:= views.PostRequest{}
			json.NewDecoder(r.Body).Decode(&data) 
			fmt.Println(data)
			if err:= model.CreateTODO(data.Name, data.Todo); err != nil{
				w.Write([]byte ("Some Error"))
				return 
			}
			w.WriteHeader(http.StatusCreated)
			json.NewEncoder(w).Encode(data)
		}
	}
}