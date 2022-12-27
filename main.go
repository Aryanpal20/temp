package main

// type User struct {
// 	Email       string `json:"email"`
// 	Password    string `json:"password"`
// 	Username    string `json:"username"`
// 	Phone       int    `json:"phone"`
// 	Address     string `json:"address"`
// 	Role        string `json:"role"`
// 	Hourly_Rate int    `json:"hourly_rate"`
// }

// type Task struct {
// 	ID                 int    `json:"id"`
// 	Assign             string `json:"assign"`
// 	Reportor           string `json:"reportor"`
// 	Title              string `json:"title"`
// 	Status             int    `default:"0"`
// 	Description        string `json:"description"`
// 	Created_At         string `json:"created_at"`
// 	Comment            string `json:"comment"`
// 	Working_Hours      int    `json:"working_hours"`
// 	Estimate_time_work string `json:"estimate_time_work"`
// 	Work_Done_time     string `json:"work_done_time"`
// }

// creating jwt token struct
// type jwtToken struct {
// 	Token string `json:"token"`
// }

// creating Exception struct
// type Exception struct {
// 	Message string `json:"message"`
// }

// creating Reponse struct
// type Reponse struct {
// 	Data string `json:"data"`
// }

// creating jwt key
// var JwtKey = []byte(os.Getenv("Jwt_Key"))

// here we want to create Database
// var Database *gorm.DB

// // Username:Password@tcp(127.0.0.1:3306)/Database_Name
// var urlDSN = "root:Java1234!@#$@tcp(127.0.0.1:3306)/detail"
// var err error

// func DataMigration() {

// 	Database, err = gorm.Open(mysql.Open(urlDSN), &gorm.Config{})

// 	if err != nil {
// 		fmt.Println(err.Error())

// 		panic("connection failed")
// 	}
// 	Database.AutoMigrate(User{}, Task{})
// }

func main() {
	DataMigration()
	HandlerRouting()

}

// func Login(w http.ResponseWriter, r *http.Request) {

// 	// here we give the data from (form-data)
// 	email := r.FormValue("email")
// 	password := r.FormValue("password")

// 	var users = User{}
// 	// here we will search the data from database
// 	Database.Where("email = ?", email).First(&users)
// 	// here we will compare the password with hash password
// 	err = bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
// 	fmt.Println(err) // nil means match

// 	if err == nil {

// 		// here we can create the token for see the values of email, username, phone, address, expire time of token.
// 		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 			"email":       users.Email,
// 			"username":    users.Username,
// 			"phone":       users.Phone,
// 			"address":     users.Address,
// 			"role":        users.Role,
// 			"hourly_rate": users.Hourly_Rate,
// 			// if we put the password here it means the password will also show with all the data.
// 			// "password": student.Password,
// 			"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
// 		})
// 		tokenString, error := token.SignedString(JwtKey)
// 		json.NewEncoder(w).Encode(jwtToken{Token: tokenString})
// 		json.NewEncoder(w).Encode(users)
// 		if error != nil {
// 			fmt.Println(error)
// 		}
// 	}
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// }

// func Create(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	var user User
// 	json.NewDecoder(r.Body).Decode(&user)

// 	password := []byte(string(user.Password))

// 	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
// 	if err != nil {
// 		panic(err)
// 	}
// 	// Comparing the password with the hash
// 	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
// 	fmt.Println(err) // nil means it is a match

// 	user.Password = string(hashedPassword)

// 	Database.Create(&user)

// }

// func Task_Create(w http.ResponseWriter, r *http.Request) {
// 	// here we can give the token in header for decode using bearer
// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	a := is_manager(token)
// 	c := task_creator(token)
// 	// fmt.Println(a)
// 	// fmt.Println(b)
// 	rep := c
// 	if a == "manager" {
// 		var task Task
// 		json.NewDecoder(r.Body).Decode(&task)
// 		task.Reportor = rep
// 		task.Created_At = time.Now().String()
// 		Database.Create(&task)
// 		json.NewEncoder(w).Encode(task)
// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}

// }

// func ChangeByManager(w http.ResponseWriter, r *http.Request) {
// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	a := is_manager(token)
// 	if a == "manager" {
// 		var tasks Task
// 		Database.First(&tasks, mux.Vars(r)["id"])
// 		json.NewDecoder(r.Body).Decode(&tasks)
// 		if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("assign", tasks.Assign).Error; err != nil {
// 			fmt.Printf("update err != nil; %v\n", err)
// 		}
// 		if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("estimate_time_work", tasks.Estimate_time_work).Error; err != nil {
// 			fmt.Printf("update err != nil; %v\n", err)
// 		}

// 		// Database.Create(&task)
// 		json.NewEncoder(w).Encode(tasks)
// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}

// }
// func ChangeByEmployee(w http.ResponseWriter, r *http.Request) {
// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	a := is_manager(token)
// 	if a == "employee" {
// 		var tasks Task
// 		Database.First(&tasks, mux.Vars(r)["id"])
// 		c := task_creator(token)
// 		rep := c
// 		if rep == tasks.Assign {
// 			// Database.First(&tasks, mux.Vars(r)["id"])
// 			json.NewDecoder(r.Body).Decode(&tasks)
// 			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("comment", tasks.Comment).Error; err != nil {
// 				fmt.Printf("update err != nil; %v\n", err)
// 			}
// 			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("status", tasks.Status).Error; err != nil {
// 				fmt.Printf("update err != nil; %v\n", err)
// 			}
// 			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("working_hours", tasks.Working_Hours).Error; err != nil {
// 				fmt.Printf("update err != nil; %v\n", err)
// 			}
// 			if err := Database.Model(&tasks).Where("reportor = ?", tasks.Reportor).Update("work_done_time", tasks.Work_Done_time).Error; err != nil {
// 				fmt.Printf("update err != nil; %v\n", err)
// 			}

// 			// Database.Create(&task)
// 			json.NewEncoder(w).Encode(tasks)
// 		} else {
// 			d := "you can't change"
// 			json.NewEncoder(w).Encode(d)
// 		}

// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}

// }

// func fetch_details(w http.ResponseWriter, r *http.Request) {

// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	c := task_creator(token)
// 	a := is_manager(token)
// 	ass := c
// 	if a == "employee" {
// 		var detail = []Task{}
// 		Database.Where("assign = ?", ass).Find(&detail)
// 		json.NewEncoder(w).Encode(detail)
// 	} else {
// 		b := "you can't the see data !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}

// }

// func Filter_Records(w http.ResponseWriter, r *http.Request) {

// 	var tasks Task
// 	params := r.URL.Query().Get("assign")

// 	id := r.URL.Query().Get("id")
// 	Database.Where("assign = ? or id = ?", params, id).Find(&tasks)
// 	// status := r.URL.Query().Get("status")

// 	// if status == "status" {
// 	// 	Database.Where("status = ?", status).Find(&tasks)
// 	// }
// 	// fmt.Println("Query string key value", params)

// 	// fmt.Println("Query string key value", id)
// 	// Database.Where("id = ?", id).Find(&tasks)

// 	// fmt.Println("Query string key value", status)
// 	// Database.Where("status = ?", status).Find(&tasks)
// 	json.NewEncoder(w).Encode(tasks)

// }

// func detailByManager(w http.ResponseWriter, r *http.Request) {

// }

// func Salary(w http.ResponseWriter, r *http.Request) {
// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	a := is_manager(token)
// 	var email string
// 	if a == "manager" {
// 		var user User
// 		var z int = 0
// 		var e int
// 		var task Task
// 		var tasks = []Task{}
// 		if user.Email == email {
// 			json.NewDecoder(r.Body).Decode(&user)
// 			Database.Where("email = ?", user.Email).Find(&user)
// 			if task.Assign == email {
// 				Database.Where("assign = ?", user.Email).Find(&tasks)
// 				// json.NewEncoder(w).Encode(tasks)
// 				for _, k := range tasks {
// 					Database.Where("assign = ?", user.Email).Find(&task)
// 					if k.Status == 1 {
// 						c := user.Hourly_Rate * k.Working_Hours
// 						if k.Estimate_time_work >= k.Work_Done_time {
// 							d := user.Hourly_Rate * k.Working_Hours
// 							e = d * 5 / 100
// 							fmt.Println("the bonous is : ", e)
// 						}
// 						z = z + c + e
// 					}

// 				}

// 			}
// 		}
// 		json.NewEncoder(w).Encode(z)
// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}
// }

// func ManagerFixRate(w http.ResponseWriter, r *http.Request) {
// 	token := strings.Split(r.Header["Token"][0], " ")[1]
// 	a := is_manager(token)
// 	var email string
// 	if a == "manager" {
// 		var user User
// 		if user.Email == email {
// 			json.NewDecoder(r.Body).Decode(&user)
// 			if err := Database.Model(&user).Where("email = ?", user.Email).Update("hourly_rate", user.Hourly_Rate).Error; err != nil {
// 				fmt.Printf("update err != nil; %v\n", err)
// 			}
// 		}
// 		Database.Save(user)
// 		json.NewEncoder(w).Encode(user)
// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}
// }

// func Profile(w http.ResponseWriter, r *http.Request) {

// 	// here we can define token(which is created by login) was entered by user from (form-data)
// 	// token := r.FormValue("Token")
// 	// a := is_manager(token)
// 	// fmt.Println(a)
// 	// if a == "admin" {
// 	// 	var users = []User{}
// 	// 	Database.Where("role IN ?", []string{"admin", "Manager", "Employee"}).Find(&users)
// 	// // 	fmt.Println(users)
// 	// // 	json.NewEncoder(w).Encode(users)
// 	// // } else {
// 	// // 	b := "access denied"
// 	// // 	json.NewEncoder(w).Encode(b)
// 	// // }
// 	// if a == "Manager" {
// 	// 	var users = []User{}
// 	// 	Database.Where("role IN ?", []string{"Manager", "Employee"}).Find(&users)
// 	// 	fmt.Println(users)
// 	// 	json.NewEncoder(w).Encode(users)
// 	// } else {
// 	// 	b := "access denied !!!"
// 	// 	json.NewEncoder(w).Encode(b)
// 	// }

// }

// func Manager(w http.ResponseWriter, r *http.Request) {

// 	token := r.FormValue("Token")
// 	a := is_manager(token)
// 	if a == "manager" {
// 		var users = []User{}
// 		Database.Where("role IN ?", []string{"manager", "employee"}).Find(&users)
// 		fmt.Println(users)
// 		json.NewEncoder(w).Encode(users)
// 	} else {
// 		b := "access denied !!!"
// 		json.NewEncoder(w).Encode(b)
// 	}
// }

// func Employee(w http.ResponseWriter, r *http.Request) {

// 	// here we can use database for fetching all same records in role field
// 	var users = []User{}
// 	Database.Where("role = ?", "employee").Find(&users)
// 	fmt.Println(users)
// 	json.NewEncoder(w).Encode(users)
// }
// func admin(w http.ResponseWriter, r *http.Request) {

// 	token := r.FormValue("Token")
// 	a := is_manager(token)
// 	if a == "admin" {
// 		var users = []User{}
// 		Database.Where("role IN ?", []string{"admin", "manager", "employee"}).Find(&users)
// 		fmt.Println(users)
// 		json.NewEncoder(w).Encode(users)
// 	} else {
// 		b := "access denied"
// 		json.NewEncoder(w).Encode(b)
// 	}
// }

// func HandlerRouting() {

// 	r := mux.NewRouter()

// 	r.HandleFunc("/users", Create).Methods("POST")
// 	r.HandleFunc("/user", Login).Methods("POST")
// 	r.HandleFunc("/user", Profile).Methods("GET")
// 	r.HandleFunc("/users", Manager).Methods("GET")
// 	r.HandleFunc("/uses", Employee).Methods("GET")
// 	r.HandleFunc("/use", admin).Methods("GET")
// 	r.HandleFunc("/task", Task_Create).Methods("POST")
// 	r.HandleFunc("/task", fetch_details).Methods("GET")
// 	r.HandleFunc("/task/{id}", ChangeByManager).Methods("PATCH")
// 	r.HandleFunc("/tasks/{id}", ChangeByEmployee).Methods("PATCH")
// 	r.HandleFunc("/filter", Filter_Records).Methods("GET")
// 	r.HandleFunc("/userss", ManagerFixRate).Methods("POST")
// 	r.HandleFunc("/salary", Salary).Methods("GET")
// 	r.HandleFunc("/detail", detailByManager).Methods("GET")

// 	log.Fatal(http.ListenAndServe(":8000", r))
// }

// // here we can use normal function and get the token by login
// func is_manager(token string) string {
// 	// here we can define a role variable
// 	var role string
// 	// here we can use for loop for split the token by dot(.) and i will store index value and part will store the value
// 	for i, part := range strings.Split(token, ".") {
// 		// here we can decode the string.
// 		decoded, err := base64.RawURLEncoding.DecodeString(part)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// here if i == 1 then it work on payload
// 		if i != 1 {
// 			continue // i == 1 is the payload
// 		}
// 		// here we can declare a M variable were we can store the decode value.
// 		var m map[string]interface{}
// 		if err := json.Unmarshal(decoded, &m); err != nil {
// 			fmt.Println("json decoding failed:", err)
// 			continue
// 		}
// 		// if email, ok := m["email"]; ok {
// 		// 	fmt.Println("email:", email)
// 		// }
// 		// fmt.Println(m)
// 		// fmt.Println(m["role"])
// 		// here we can save the value in role variable for return the string value
// 		role = m["role"].(string)

// 	}
// 	return role
// 	// fmt.Println(role)
// }
// func task_creator(token string) string {
// 	// here we can define a role variable
// 	var role1 string
// 	// here we can use for loop for split the token by dot(.) and i will store index value and part will store the value
// 	for i, part := range strings.Split(token, ".") {
// 		// here we can decode the string.
// 		decoded, err := base64.RawURLEncoding.DecodeString(part)
// 		if err != nil {
// 			panic(err)
// 		}
// 		// here if i == 1 then it work on payload
// 		if i != 1 {
// 			continue // i == 1 is the payload
// 		}
// 		// here we can declare a M variable were we can store the decode value.
// 		var m map[string]interface{}
// 		if err := json.Unmarshal(decoded, &m); err != nil {
// 			fmt.Println("json decoding failed:", err)
// 			continue
// 		}

// 		// here we can save the value in role variable for return the string value
// 		role1 = m["email"].(string)

// 	}
// 	return role1
// 	// fmt.Println(role)
// }
