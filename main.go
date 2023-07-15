package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ihksanghazi/go-restfulapi-relational/controllers"
	"github.com/ihksanghazi/go-restfulapi-relational/database"
	"github.com/ihksanghazi/go-restfulapi-relational/models"
)

func main() {

	database.ConnectDatabase()

	database.DB.AutoMigrate(
		&models.User{},
		&models.Locker{},
		&models.Post{},
	)

	r := mux.NewRouter()

	r.HandleFunc("/users",controllers.GetAllUsers).Methods("GET")
	r.HandleFunc("/users",controllers.CreateUser).Methods("POST")

	r.HandleFunc("/lockers",controllers.GetAllLockers).Methods("GET")
	r.HandleFunc("/lockers",controllers.CreateLocker).Methods("POST")
	
	r.HandleFunc("/posts",controllers.GetAllPosts).Methods("GET")
	r.HandleFunc("/posts",controllers.CreatePost).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000",r))
}