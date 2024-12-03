package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/fjanza/go-gorm-restapi/entity"
	"github.com/fjanza/go-gorm-restapi/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the posts"`))
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func GetPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := r.URL.Query().Get("id")
	if id == "" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error empty id"}`))
		return
	}
	intId, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error at convert id"}`))
		return
	}
	posts, err := repo.Find(intId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error getting the posts"}`))
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func AddPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var post entity.Post

	err := json.NewDecoder(r.Body).Decode(&post)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{"error": "Error unmarshalling data"}`))
		return
	}

	post.ID = int64(rand.Int())
	repo.Save(&post)
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(post)
}
