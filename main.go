package main

import (
	database "api/database"
	rating "api/rating"
	routing "api/routing"
)

// creating Exception struct
// type Exception struct {
// 	Message string `json:"message"`
// }

// creating Reponse struct
// type Reponse struct {
// 	Data string `json:"data"`
// }

func main() {

	database.DataMigration()
	rating.RunCronJobs()
	routing.HandlerRouting()

}

// func Filter_Records(w http.ResponseWriter, r *http.Request) {

// 	var tasks Task
// 	params := r.URL.Query().Get("assign")

// 	id := r.URL.Query().Get("id")
// 	Database.Where("assign = ? or id = ?", params, id).Find(&tasks)
// 	// status := r.URL.Query().Get("status")

// 	// if status == "status" {
// 	// 	Database.Where("status = ?", status).Find(&tasks)
// 	// }
// 	// fmt.Println("Query string key value", params)

// 	// fmt.Println("Query string key value", id)
// 	// Database.Where("id = ?", id).Find(&tasks)

// 	// fmt.Println("Query string key value", status)
// 	// Database.Where("status = ?", status).Find(&tasks)
// 	json.NewEncoder(w).Encode(tasks)

// }

// func Profile(w http.ResponseWriter, r *http.Request) {

// 	// here we can define token(which is created by login) was entered by user from (form-data)
// 	// token := r.FormValue("Token")
// 	// a := is_manager(token)
// 	// fmt.Println(a)
// 	// if a == "admin" {
// 	// 	var users = []User{}
// 	// 	Database.Where("role IN ?", []string{"admin", "Manager", "Employee"}).Find(&users)
// 	// // 	fmt.Println(users)
// 	// // 	json.NewEncoder(w).Encode(users)
// 	// // } else {
// 	// // 	b := "access denied"
// 	// // 	json.NewEncoder(w).Encode(b)
// 	// // }
// 	// if a == "Manager" {
// 	// 	var users = []User{}
// 	// 	Database.Where("role IN ?", []string{"Manager", "Employee"}).Find(&users)
// 	// 	fmt.Println(users)
// 	// 	json.NewEncoder(w).Encode(users)
// 	// } else {
// 	// 	b := "access denied !!!"
// 	// 	json.NewEncoder(w).Encode(b)
// 	// }

// }
