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

func FetchDetailByStatus(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)

	if a == "manager" {

		var tasks = []entity.Task{}
		assign := r.URL.Query().Get("assign")

		status := r.URL.Query().Get("status")
		if assign != "" {
			database.Database.Where("assign = ?", assign).Find(&tasks)
		} else if status != "" {
			database.Database.Where("status = ?", status).Find(&tasks)
		}
		if assign != "" && status != "" {
			database.Database.Where("assign = ? and status = ?", assign, status).Find(&tasks)
		}
		for _, k := range tasks {
			json.NewEncoder(w).Encode(k)
		}
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}

func PostFeedback(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
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
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}
