module api

go 1.19

require (
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/gorilla/mux v1.8.0
	golang.org/x/crypto v0.4.0
	gorm.io/driver/mysql v1.4.4
	gorm.io/gorm v1.24.2
)

require (
	github.com/go-co-op/gocron v1.18.0 // indirect
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/robfig/cron/v3 v3.0.1 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
)

replace api/database => ./database

replace api/changebyemployee => ./changebyemployee

replace api/fetchemail => ./fetchemail

replace api/fetchrole => ./fetchrole
