package alltaskdetail

import (
	database "api/database"
	entity "api/entity"
	role "api/fetchrole"
	"encoding/json"
	"net/http"
	"strings"
)

func DetailByManager(w http.ResponseWriter, r *http.Request) {
	token := strings.Split(r.Header["Token"][0], " ")[1]
	a := role.Is_manager(token)
	var email string
	// here we can compare with manager
	if a == "manager" {
		var task = []entity.Task{}
		var user entity.User
		var c, z, d, e, f int
		var m map[string]int
		var p, q, s, t string
		var x string = "task"
		if user.Email == email {
			json.NewDecoder(r.Body).Decode(&user)
			database.Database.Where("assign = ?", user.Email).Find(&task)
			for _, k := range task {
				if k.Assign == user.Email {
					z = z + 1
				}
				if k.Status == 1 {
					c = c + 1
					p = k.Comment
				}
				if k.Status == 0 {
					d = d + 1
					q = k.Comment
				}
				if k.Status == 2 {
					e = e + 1
					s = k.Comment
				}
				if k.Status == 3 {
					f = f + 1
					t = k.Comment
				}
				m = map[string]int{x: z, p: c, q: d, s: e, t: f}
				// here we can remove default value
				delete(m, "")
			}

		}
		json.NewEncoder(w).Encode(user.Email)
		json.NewEncoder(w).Encode(m)

	}

}
