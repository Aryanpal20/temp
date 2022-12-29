package register

import (
	"encoding/json"
	"fmt"
	"net/http"

	database "api/database"
	entity "api/entity"

	"golang.org/x/crypto/bcrypt"
)

func Create(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	var user entity.User
	json.NewDecoder(r.Body).Decode(&user)
	// here we can give the user Password in byte form in password variable
	password := []byte(string(user.Password))
	// here we can create hashed Passwword by bcrypt
	hashedPassword, err := bcrypt.GenerateFromPassword(password, 10)
	if err != nil {
		panic(err)
	}
	// Comparing the password with the hash
	err = bcrypt.CompareHashAndPassword(hashedPassword, password)
	fmt.Println(err) // nil means it is a match

	// here we can convert hashed password in string form and store in user Password
	user.Password = string(hashedPassword)

	// here we can create the data on the database but the password will be saved in hashed Password
	database.Database.Create(&user)

}
