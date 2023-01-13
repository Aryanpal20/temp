package relations

import (
	"api/database"
	"api/entity"
	"encoding/json"
	"log"
	"net/http"
)

func GetRelation(w http.ResponseWriter, r *http.Request) {

	var user entity.User
	var tasks []entity.Task
	email := r.FormValue("email")
	database.Database.Where("email = ?", email).Find(&user)
	if err := database.Database.Joins(" tasks JOIN users on tasks.user_id=users.id").
		Where("users.email=?", email).
		Find(&tasks).Error; err != nil {
		log.Fatal(err)
	}
	user.Tasks = append(user.Tasks, tasks...)
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
	// json.NewEncoder(w).Encode(user.Email)
	// json.NewEncoder(w).Encode(user.Username)
	// json.NewEncoder(w).Encode(user.Tasks)
}
