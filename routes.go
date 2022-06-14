package main

import (
	"api-go/entity"
	"api-go/repository"
	"encoding/json"
	"math/rand"
	"net/http"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func getPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		response.WriteHeader((http.StatusInternalServerError))
		response.Write([]byte(`{"error: "Error getting posts"}`))
	}

	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(posts)

}

func addPosts(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-type", "application/json")
	var post entity.Post
	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		response.WriteHeader((http.StatusInternalServerError))
		response.Write([]byte(`{"error: "Error marshalling the request"}`))
		return
	}

	post.ID = rand.Int63()
	repo.Save(&post)
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(post)

}
