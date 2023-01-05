package rating

import (
	"api/database"
	"api/entity"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func Is_Rating() {

	var users = []entity.User{}
	var task = []entity.Task{}
	// here we want fetch the data from database
	database.Database.Find(&users)
	database.Database.Find(&task)

	n := 4
	// here we use for loop for in users
	for _, i := range users {
		// here we can compare a email in task table.
		database.Database.Where("assign = ?", i.Email).Find(&task)
		fmt.Println(i.Email)
		sum := 0.0
		for _, k := range task {
			if k.Rating != 0.0 {
				c := k.Rating
				sum = sum + c
			}
		}

		avg := (sum) / (float64(n))
		fmt.Println("Sum = ", sum, "\nAverage = ", avg)
	}
}

func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}
func RunCronJobs() {
	s := gocron.NewScheduler(time.UTC)
	s.Every(1).Seconds().Do(func() {
		// hello("chonu")
		Is_Rating()

	})
	s.StartBlocking()
}
