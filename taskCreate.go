package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

func Task_Create(w http.ResponseWriter, r *http.Request) {
	// here we can give the token in header for decode using bearer
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := is_manager(token)
	c := task_creator(token)
	// fmt.Println(a)
	// fmt.Println(b)
	rep := c
	if a == "manager" {
		var task Task
		json.NewDecoder(r.Body).Decode(&task)
		task.Reportor = rep
		task.Created_At = time.Now().String()
		Database.Create(&task)
		json.NewEncoder(w).Encode(task)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}

}
