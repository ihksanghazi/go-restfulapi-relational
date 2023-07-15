package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/ihksanghazi/go-restfulapi-relational/database"
	"github.com/ihksanghazi/go-restfulapi-relational/helpers"
	"github.com/ihksanghazi/go-restfulapi-relational/models"
)

func GetAllPosts(w http.ResponseWriter, r *http.Request) {
	var posts []models.Post

	if err:= database.DB.Preload("User").Find(&posts).Error; err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusInternalServerError,response)
		return
	}

	helpers.ResponseJSON(w,http.StatusOK,posts)
}

func CreatePost(w http.ResponseWriter, r *http.Request) {
	// mengambil input dari json
	var postInput models.Post
	decoder:= json.NewDecoder(r.Body)
	if err:= decoder.Decode(&postInput); err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusBadRequest,response)
		return
	}
	defer r.Body.Close()

	if err:= database.DB.Create(&postInput).Error; err != nil {
		response:= map[string]string{"msg":err.Error()}
		helpers.ResponseJSON(w,http.StatusInternalServerError,response)
		return
	}

	helpers.ResponseJSON(w,http.StatusCreated,postInput)
}
