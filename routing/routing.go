package routing

import (
	"log"
	"net/http"

	Admin "api/adminaccess"
	alltaskdetail "api/alltaskdetail"
	changebyemployee "api/changebyemployee"
	changebymanager "api/changebymanager"
	"api/cron"
	feedback "api/feedback"
	fetch "api/fetchdetail"
	filterrecord "api/filterrecord"
	login "api/login"
	fixrate "api/managerfixrate"
	create "api/register"
	rel "api/relations"
	salary "api/salary"
	same "api/samerecord"
	task "api/taskcreater"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()

	r.HandleFunc("/users", create.Create).Methods("POST")
	r.HandleFunc("/user", login.Login).Methods("POST")
	// r.HandleFunc("/user", Profile).Methods("GET")
	r.HandleFunc("/users", same.Manager).Methods("GET")
	r.HandleFunc("/uses", same.Employee).Methods("GET")
	r.HandleFunc("/use", same.Admin).Methods("GET")
	r.HandleFunc("/task", task.Task_Create).Methods("POST")
	r.HandleFunc("/task", fetch.Fetchdetail).Methods("GET")
	r.HandleFunc("/task/{id}", changebymanager.ChangeByManager).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", changebyemployee.ChangeByEmployee).Methods("PATCH")
	r.HandleFunc("/filter", filterrecord.Filter_Records).Methods("GET")
	r.HandleFunc("/userss", fixrate.ManagerFixRate).Methods("POST")
	r.HandleFunc("/salary", salary.Salary).Methods("GET")
	r.HandleFunc("/detail", alltaskdetail.DetailByManager).Methods("GET")
	r.HandleFunc("/feedback", feedback.FetchDetailByStatus).Methods("GET")
	r.HandleFunc("/feedback/{id}", feedback.PostFeedback).Methods("POST")
	r.HandleFunc("/data", cron.Data).Methods("GET")
	r.HandleFunc("/admin", Admin.Admin).Methods("GET")
	r.HandleFunc("/relation", rel.GetRelation).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
