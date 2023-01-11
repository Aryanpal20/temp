package login

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	database "api/database"
	entity "api/entity"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// here we can creating jwt token struct
type jwtToken struct {
	Token string `json:"token"`
}

// here we can creating jwt key
var JwtKey = []byte(os.Getenv("Jwt_Key"))

func Login(w http.ResponseWriter, r *http.Request) {

	// here we give the data from (form-data)
	email := r.FormValue("email")
	password := r.FormValue("password")

	var users = entity.User{}
	// here we will search the data from database
	database.Database.Where("email = ?", email).First(&users)
	// here we will compare the password with hash password
	err := bcrypt.CompareHashAndPassword([]byte(users.Password), []byte(password))
	fmt.Println(err) // nil means match

	if err == nil {

		// here we can create the token for see the values of email, username, phone, address, expire time of token.
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"email":       users.Email,
			"username":    users.Username,
			"phone":       users.Phone,
			"address":     users.Address,
			"role":        users.Role,
			"hourly_rate": users.Hourly_Rate,
			// if we put the password here it means the password will also show with all the data.
			// "password": student.Password,
			"exp": time.Now().Add(time.Hour * time.Duration(1)).Unix(),
		})
		tokenString, error := token.SignedString(JwtKey)
		json.NewEncoder(w).Encode(jwtToken{Token: tokenString})
		json.NewEncoder(w).Encode(users)
		if error != nil {
			fmt.Println(error)
		}
	}
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}
