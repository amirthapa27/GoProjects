package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rest-api/entity"
	"rest-api/errors"
	"rest-api/service"
	"strconv"

	"github.com/gorilla/mux"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(w http.ResponseWriter, r *http.Request)
	CreatePost(w http.ResponseWriter, r *http.Request)
	DeletePost(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	UpdatePost(w http.ResponseWriter, r *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

// func init() {
// 	posts = []Post{{Id: 1, Title: "title 1", Text: "text1"}}

// }

func (*controller) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error in gettting the posts"})
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(posts)
}

func (*controller) CreatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error in decoding data"})
		return
	}

	err1 := postService.Validate(&post)
	if err1 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}

	result, err2 := postService.AddPost(&post)
	if err2 != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error while creating post"})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func (*controller) DeletePost(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	stringID := vars["id"]
	id, err := strconv.ParseInt(stringID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error in parsing id"})

		return
	}
	err = postService.DeleteOne(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error in deleting the post"})
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Post successfully deleted")

}

func (*controller) GetOne(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	stringID := vars["id"]
	id, err := strconv.ParseInt(stringID, 0, 0)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error in parsing id"})
	}
	post, err := postService.GetOne(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(errors.ServiceError{Message: "Error in getting post"})
	}
	json.NewEncoder(w).Encode(post)
}

func (*controller) UpdatePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var posts []entity.Post

	vars := mux.Vars(r)
	stringID := vars["id"]
	id, err := strconv.ParseInt(stringID, 0, 0)
	// err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		fmt.Println("Error is parsing ID")
	}
	for index, item := range posts {
		if item.ID == id {
			posts = append(posts[:index], posts[index+1:]...)
		}
	}
	var post entity.Post
	_ = json.NewDecoder(r.Body).Decode(&post)
	post.ID = id
	posts = append(posts, post)
	postService.UpdatePost(&post)
	json.NewEncoder(w).Encode(&post)

}
