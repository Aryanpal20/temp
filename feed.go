package main

// import (
// 	database "api/database"
// 	"api/entity"
// 	role "api/fetchrole"
// 	"encoding/json"
// 	"net/http"
// 	"strings"
// )

// func FetchDetailByStatus(w http.ResponseWriter, r *http.Request) {
// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	a := role.Is_manager(token)
// 	var email string
// 	var status entity.Task

// 	json.NewDecoder(r.Body).Decode(&status)

// 	if a == "manager" {
// 		var user entity.User
// 		var task entity.Task
// 		var tasks = []entity.Task{}

// 		if user.Email == email {
// 			json.NewDecoder(r.Body).Decode(&user)
// 			database.Database.Where("email = ?", user.Email).Find(&user)
// 		}
// 		if task.Assign == email {
// 			// fmt.Println("USER_EMAIL: ", user.Email)
// 			database.Database.Where("assign = ? or status = ?", user.Email, status.Status).Find(&tasks)
// 			// fmt.Println(tasks, "HGHGHGU")
// 			for _, k := range tasks {
// 				database.Database.Where("assign = ?", user.Email).Find(&task)
// 				if k.Status == 0 {
// 					c := k.ID
// 					json.NewEncoder(w).Encode(c)
// 					json.NewEncoder(w).Encode(k)
// 				}
// 				if k.Status == 1 {
// 					c := k.ID
// 					json.NewEncoder(w).Encode(c)
// 					json.NewEncoder(w).Encode(k)
// 				}
// 				if k.Status == 2 {
// 					c := k.ID
// 					json.NewEncoder(w).Encode(c)
// 					json.NewEncoder(w).Encode(k)
// 				}
// 				if k.Status == 3 {
// 					c := k.ID
// 					json.NewEncoder(w).Encode(c)
// 					json.NewEncoder(w).Encode(k)
// 				}

// 			}

// 		}
// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}

// }
