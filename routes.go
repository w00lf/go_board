package main

import (
	"net/http"
	"w00lf/go_board/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
)

func initializeRouter() http.Handler {
	router := httprouter.New()
	router.GET("/", handlerBoardsIndex)
	router.POST("/", handlerBoardSave)

  router.GET("/boards/:board_id", handlerPostsIndex)
  router.GET("/boards/:board_id/posts", handlerPostsIndex)

	router.GET("/boards/:board_id/posts/:id", handlerShow)

  router.POST("/boards/:board_id/posts", handlerPostSave)
  router.POST("/boards/:board_id/posts/:id", handlerPostSave)
	return router
}
