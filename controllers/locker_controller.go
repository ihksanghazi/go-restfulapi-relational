package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ihksanghazi/go-restfulapi-relational/database"
	"github.com/ihksanghazi/go-restfulapi-relational/helpers"
	"github.com/ihksanghazi/go-restfulapi-relational/models"
)

func GetAllLockers(w http.ResponseWriter, r *http.Request) {
	var users []models.Locker

	if err:= database.DB.Preload("User").Find(&users).Error; err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusInternalServerError,response)
		return
	}

	helpers.ResponseJSON(w,http.StatusOK,users)
}

func CreateLocker(w http.ResponseWriter, r *http.Request) {
	// mengambil input dari json
	var lockerInput models.Locker
	decoder:= json.NewDecoder(r.Body)
	if err:= decoder.Decode(&lockerInput); err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusBadRequest,response)
		return
	}
	defer r.Body.Close()

	if err:= database.DB.Create(&lockerInput).Error; err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusInternalServerError,response)
		return
	}

	helpers.ResponseJSON(w,http.StatusCreated,lockerInput)
}
