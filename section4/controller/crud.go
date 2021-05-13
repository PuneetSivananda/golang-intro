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
		} else if r.Method == http.MethodGet {  
			name := r.URL.Query().Get("name")
			data, err:= model.ReadByName(name)
			if err!=nil{
				w.Write([]byte(err.Error()))
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(data)
		} else if r.Method == http.MethodDelete{
			name := r.URL.Path[1:]
			err:= model.DeleteByName(name)
			if err!=nil{
				w.Write([]byte("DELETE Error"))
				return
			}
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(struct{
				Status string `json:status`
			}{"Item Deleted"})
		}
	}
}