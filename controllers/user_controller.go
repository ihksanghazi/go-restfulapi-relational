package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ihksanghazi/go-restfulapi-relational/database"
	"github.com/ihksanghazi/go-restfulapi-relational/helpers"
	"github.com/ihksanghazi/go-restfulapi-relational/models"
)

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	var users []models.User

	if err:= database.DB.Preload("Locker").Find(&users).Error; err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusInternalServerError,response)
		return
	}

	helpers.ResponseJSON(w,http.StatusOK,users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	// mengambil input dari json
	var userInput models.User
	decoder:= json.NewDecoder(r.Body)
	if err:= decoder.Decode(&userInput); err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusBadRequest,response)
		return
	}
	defer r.Body.Close()

	if err:= database.DB.Create(&userInput).Error; err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusInternalServerError,response)
		return
	}

	helpers.ResponseJSON(w,http.StatusCreated,userInput)
}
