package fetchdetail

import (
	database "api/database"
	entity "api/entity"
	email "api/fetchemail"
	role "api/fetchrole"
	"encoding/json"
	"net/http"
	"strings"
)

func Fetchdetail(w http.ResponseWriter, r *http.Request) {

	// here we can give the token and decode the token
	token := strings.Split(r.Header["Token"][0], " ")[1]
	// here we can store the email value in c which is present in token
	c := email.Task_creator(token)
	// here we can store the role value in a which is present in token
	a := role.Is_manager(token)
	ass := c
	// here we can comapre the role
	if a == "employee" {
		var detail = []entity.Task{}
		// here we can fetch the details of email value which is present in token
		database.Database.Where("assign = ?", ass).Find(&detail)
		json.NewEncoder(w).Encode(detail)
	} else {
		b := "you can't the see data !!!"
		json.NewEncoder(w).Encode(b)
	}

}
