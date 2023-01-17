package adminaccess

import (
	"api/database"
	"api/entity"
	role "api/fetchrole"
	"encoding/json"
	"net/http"
	"strings"
)

func Admin(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	var task = []entity.Task{}
	// here we can compare with admin
	if a == "admin" {
		// here we can find the data from task table
		database.Database.Find(&task)
		json.NewEncoder(w).Encode(task)

	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}
}
