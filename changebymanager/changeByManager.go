package changebymanager

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	database "api/database"
	entity "api/entity"
	role "api/fetchrole"

	"github.com/gorilla/mux"
)

func ChangeByManager(w http.ResponseWriter, r *http.Request) {
	// here we can split the token and decode the token
	token := strings.Split(r.Header["Token"][0], " ")[1]
	// here we can compare with employee
	a := role.Is_manager(token)
	if a == "manager" {
		var tasks entity.Task
		// here we can find the from database by id
		database.Database.First(&tasks, mux.Vars(r)["id"])
		json.NewDecoder(r.Body).Decode(&tasks)
		// here we can update the assign by manager
		if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("assign", tasks.Assign).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}
		// here we can update the estimate time work by manager
		if err := database.Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("estimate_time_work", tasks.Estimate_time_work).Error; err != nil {
			fmt.Printf("update err != nil; %v\n", err)
		}

		// Database.Create(&task)
		json.NewEncoder(w).Encode(tasks)
	} else {
		b := "access denied !!!"
		json.NewEncoder(w).Encode(b)
	}

}
