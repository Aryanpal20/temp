package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandlerRouting() {

	r := mux.NewRouter()

	r.HandleFunc("/users", Create).Methods("POST")
	r.HandleFunc("/user", Login).Methods("POST")
	// r.HandleFunc("/user", Profile).Methods("GET")
	r.HandleFunc("/users", Manager).Methods("GET")
	r.HandleFunc("/uses", Employee).Methods("GET")
	r.HandleFunc("/use", admin).Methods("GET")
	r.HandleFunc("/task", Task_Create).Methods("POST")
	r.HandleFunc("/task", fetch_details).Methods("GET")
	r.HandleFunc("/task/{id}", ChangeByManager).Methods("PATCH")
	r.HandleFunc("/tasks/{id}", ChangeByEmployee).Methods("PATCH")
	r.HandleFunc("/filter", Filter_Records).Methods("GET")
	r.HandleFunc("/userss", ManagerFixRate).Methods("POST")
	r.HandleFunc("/salary", Salary).Methods("GET")
	r.HandleFunc("/detail", detailByManager).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
