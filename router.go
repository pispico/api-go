package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{{Id: 1, Title: "titulo1", Text: "Texto 1"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader((http.StatusInternalServerError))
		resp.Write([]byte(`{"error: "Error marshalling posts array"}`))
		return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPosts(resp http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader((http.StatusInternalServerError))
		resp.Write([]byte(`{"error: "Error marshalling the request"}`))
		return
	}

	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	resp.Write(result)

}
