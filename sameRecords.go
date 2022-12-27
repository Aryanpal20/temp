package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Manager(w http.ResponseWriter, r *http.Request) {

	token := r.FormValue("Token")
	a := is_manager(token)
	if a == "manager" {
		var users = []User{}
		Database.Where("role IN ?", []string{"manager", "employee"}).Find(&users)
		fmt.Println(users)
		json.NewEncoder(w).Encode(users)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}

func Employee(w http.ResponseWriter, r *http.Request) {

	// here we can use database for fetching all same records in role field
	var users = []User{}
	Database.Where("role = ?", "employee").Find(&users)
	fmt.Println(users)
	json.NewEncoder(w).Encode(users)
}
func admin(w http.ResponseWriter, r *http.Request) {

	token := r.FormValue("Token")
	a := is_manager(token)
	if a == "admin" {
		var users = []User{}
		Database.Where("role IN ?", []string{"admin", "manager", "employee"}).Find(&users)
		fmt.Println(users)
		json.NewEncoder(w).Encode(users)
	} else {
		b := "access denied"
		json.NewEncoder(w).Encode(b)
	}
}
