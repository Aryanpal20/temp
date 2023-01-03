package feedback

import (
	"api/database"
	entity "api/entity"
	role "api/fetchrole"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

// here we want to  create a get api for fetch the deatils by status and email.
func FetchDetailByStatus(w http.ResponseWriter, r *http.Request) {
	// here we can split the token and decode the token
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	// here we can compare with manager
	if a == "manager" {

		var tasks = []entity.Task{}
		// here we get the data from params
		assign := r.URL.Query().Get("assign")
		// here we get the from params
		status := r.URL.Query().Get("status")
		// here we give the condition assign not equal too nil
		if assign != "" {
			database.Database.Where("assign = ?", assign).Find(&tasks)
		} else if status != "" { // here we give the condition status not equal too nil
			database.Database.Where("status = ?", status).Find(&tasks)
		}
		// here we give the condition assign not equal too nil and status not equal ton nil
		if assign != "" && status != "" {
			database.Database.Where("assign = ? and status = ?", assign, status).Find(&tasks)
		}
		// here we can use for loop for see all the data ion beautifull manner.
		for _, k := range tasks {
			json.NewEncoder(w).Encode(k)
		}
		// here we can use else condition for doesn't access by anyone except manager.
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}

// here we want to create a post api for only manager access
func PostFeedback(w http.ResponseWriter, r *http.Request) {
	// here we can split the token and decode the token
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	// here we can compare with manager
	if a == "manager" {
		var tasks entity.Task
		// here we can find the data from database by id
		database.Database.First(&tasks, mux.Vars(r)["id"])
		json.NewDecoder(r.Body).Decode(&tasks)
		// here we can update the feedback by manager
		if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("feedback", tasks.Feedback).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		// here we can update the status by manager
		if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("status", tasks.Status).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		// here we can update the comment by manager
		if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("comment", tasks.Comment).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}

		json.NewEncoder(w).Encode(tasks)
		// here we can use else condition for doesn't access by anyone except manager.
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}
