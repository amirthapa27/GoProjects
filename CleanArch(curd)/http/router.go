package router

import "net/http"

type Router interface {
	// router.HandleFunc("/posts", GetPosts).Methods("GET")
	GET(uri string, f func(w http.ResponseWriter, r *http.Request))
	// router.HandleFunc("/post", CreatePost).Methods("POST")
	POST(uri string, f func(w http.ResponseWriter, r *http.Request))
	DELETE(uri string, f func(w http.ResponseWriter, r *http.Request))
	GETONE(uri string, f func(w http.ResponseWriter, r *http.Request))
	// http.ListenAndServe(":8000", router)
	UPDATE(uri string, f func(w http.ResponseWriter, r *http.Request))
	SERVE(port string)
}
