package rating

import (
	"api/database"
	"api/entity"
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

// here er can create a function for Calculate the Monthly average rating of employee
func Is_Rating() {

	var users = []entity.User{}
	var task = []entity.Task{}
	// here we want fetch the data from database
	database.Database.Find(&users)
	database.Database.Find(&task)

	n := 4
	// here we use for loop for in users
	for _, i := range users {
		// here we can fetch the email in task table.
		database.Database.Where("assign = ?", i.Email).Find(&task)
		sum := 0.0
		// here we can compare the role.
		if i.Role == "employee" {
			fmt.Println(i.Email)
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
}

// here caan create a hello function
func hello(name string) {
	message := fmt.Sprintf("Hi, %v", name)
	fmt.Println(message)
}
func RunCronJobs() {
	// here we can create a cron job scheduler.
	s := gocron.NewScheduler(time.Local)
	// here we can set the time for is rating function.
	job, _ := s.Every(1).Day().At("12:15").Do(func() {
		// hello("chonu")
		Is_Rating()

	})
	// here we can set the time for hello function.
	job1, _ := s.Every(1).Day().At("12:17").Do(func() {
		hello("chonu")
	})
	s.StartAsync()
	fmt.Println(job.ScheduledTime())
	fmt.Println(job1.ScheduledTime())
}
