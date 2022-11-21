package main

import (
	"fmt"
	"net/http"

	"rest-api/controller"
	router "rest-api/http"
	"rest-api/repository"
	"rest-api/service"
)

var (
	postRepo       repository.PostRepo       = repository.NewPSQLRepo() //replace NewFireStoreRepo with the datbase repo you want
	postService    service.PostService       = service.NewPostService(postRepo)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewMuxRouter() //change router by replacing NewMuxRouter
)

func main() {

	const port = ":8000"

	httpRouter.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "helloo")
	})

	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.CreatePost)
	httpRouter.DELETE("/post/{id}", postController.DeletePost)
	httpRouter.GETONE("/post/{id}", postController.GetOne)
	httpRouter.UPDATE("/post/{id}", postController.UpdatePost)
	httpRouter.SERVE(port)

}
