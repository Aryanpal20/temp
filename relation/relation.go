// here we can use preload

package relation

import (
	"api/database"
	"api/entity"
	"encoding/json"
	"fmt"
	"net/http"
)

func Relation(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	var tasks []entity.Task
	email := r.FormValue("email")
	database.Database.Select("email", "username", "id").Where("email = ?", email).Find(&user)
	fmt.Println(user)
	database.Database.Where("assign = ?", email).Preload("Users").Find(&tasks)
	user.Tasks = append(user.Tasks, tasks...)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}
