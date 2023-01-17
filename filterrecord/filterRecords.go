package filterrecord

import (
	database "api/database"
	entity "api/entity"
	"encoding/json"
	"net/http"
)

func Filter_Records(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var tasks = []entity.Task{}
	// here user give the assign, id, status in params
	assign := r.URL.Query().Get("assign")
	status := r.URL.Query().Get("status")
	id := r.URL.Query().Get("id")
	if assign != "" {
		database.Database.Where("assign = ?", assign).Find(&tasks)
	} else if id != "" {
		database.Database.Where("id = ?", id).Find(&tasks)
	} else if status != "" {
		database.Database.Where("status = ?", status).Find(&tasks)
	}
	// here we can use and operator for both params are right then it will give data otherwise no data provide.
	if assign != "" && id != "" {
		database.Database.Where("assign = ? and id = ?", assign, id).Find(&tasks)
	} else if assign != "" && status != "" {
		database.Database.Where("assign = ? and status = ?", assign, status).Find(&tasks)
	} else if id != "" && status != "" {
		database.Database.Where("id = ? and status = ?", id, status).Find(&tasks)
	} else { // here we can use and operator for all the params are right then it will give data otherwise no data provide.
		database.Database.Where("id = ? and status = ? and params = ?", id, status, assign).Find(&tasks)
	}
	json.NewEncoder(w).Encode(tasks)

}
