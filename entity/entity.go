package entity

// here we can create the struct as well as table User
type User struct {
	Email       string `json:"email"`
	Password    string `json:"password"`
	Username    string `json:"username"`
	Phone       string `json:"phone"`
	Address     string `json:"address"`
	Role        string `json:"role"`
	Hourly_Rate int    `json:"hourly_rate"`
}

// here we can create the struct as well as table Task
type Task struct {
	ID                 int     `json:"id"`
	Assign             string  `json:"assign"`
	Reportor           string  `json:"reportor"`
	Title              string  `json:"title"`
	Status             int     `default:"0"`
	Description        string  `json:"description"`
	Created_At         string  `json:"created_at"`
	Comment            string  `json:"comment"`
	Working_Hours      int     `json:"working_hours"`
	Estimate_time_work string  `json:"estimate_time_work"`
	Work_Done_time     string  `json:"work_done_time"`
	Feedback           string  `json:"feedback"`
	Rating             float64 `json:"rating"`
	Average_Rating     float64 ` json:"average_rating"`
}
