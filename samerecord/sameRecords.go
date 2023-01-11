package samerecord

import (
	database "api/database"
	entity "api/entity"
	role "api/fetchrole"
	"encoding/json"
	"net/http"
)

func Manager(w http.ResponseWriter, r *http.Request) {

	// here we can give the token
	token := r.FormValue("Token")
	// here we can use is_manager fuction where we can the decode the token the fetch the role value
	a := role.Is_manager(token)
	// here we can compare the role is equal to manager
	if a == "manager" {
		var users = []entity.User{}
		// here we can use  query for fetch the role is equal to manager and employee
		database.Database.Where("role IN ?", []string{"manager", "employee"}).Find(&users)
		json.NewEncoder(w).Encode(users)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}

func Employee(w http.ResponseWriter, r *http.Request) {

	// here we can use database for fetching all same records in role field
	var users = []entity.User{}
	// here we can use  query for fetch the role is equal to employee
	database.Database.Where("role = ?", "employee").Find(&users)
	json.NewEncoder(w).Encode(users)
}

func Admin(w http.ResponseWriter, r *http.Request) {

	// here we can give the token
	token := r.FormValue("Token")
	// here we can use is_manager fuction where we can the decode the token the fetch the role value
	a := role.Is_manager(token)
	// here we can compare the role is equal to admin
	if a == "admin" {
		var users = []entity.User{}
		// here we can use  query for fetch the role is equal to admin and manager and employee.
		database.Database.Where("role IN ?", []string{"admin", "manager", "employee"}).Find(&users)
		json.NewEncoder(w).Encode(users)
	} else {
		b := "access denied"
		json.NewEncoder(w).Encode(b)
	}
}
