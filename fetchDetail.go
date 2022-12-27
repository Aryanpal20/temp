package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func fetch_details(w http.ResponseWriter, r *http.Request) {

	token := strings.Split(r.Header["Token"][0], " ")[1]
	c := task_creator(token)
	a := is_manager(token)
	ass := c
	if a == "employee" {
		var detail = []Task{}
		Database.Where("assign = ?", ass).Find(&detail)
		json.NewEncoder(w).Encode(detail)
	} else {
		b := "you can't the see data !!!"
		json.NewEncoder(w).Encode(b)
	}

}
