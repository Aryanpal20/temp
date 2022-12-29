package filterrecord

import (
	database "api/database"
	entity "api/entity"
	"encoding/json"
	"net/http"
)

func Filter_Records(w http.ResponseWriter, r *http.Request) {

	var tasks = []entity.Task{}
	params := r.URL.Query().Get("assign")

	id := r.URL.Query().Get("id")
	database.Database.Where("assign = ? or id = ?", params, id).Find(&tasks)
	// status := r.URL.Query().Get("status")

	// if status == "status" {
	// 	Database.Where("status = ?", status).Find(&tasks)
	// }
	// fmt.Println("Query string key value", params)

	// fmt.Println("Query string key value", id)
	// Database.Where("id = ?", id).Find(&tasks)

	// fmt.Println("Query string key value", status)
	// Database.Where("status = ?", status).Find(&tasks)
	json.NewEncoder(w).Encode(tasks)

}
