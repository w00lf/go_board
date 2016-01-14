package main

import (
	"net/http"
	"w00lf/go_board/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
)

func initializeRouter() http.Handler {
	router := httprouter.New()
	router.GET("/", handlerIndex)
	router.POST("/", handlerSave)
	router.GET("/posts/:id", handlerShow)
  router.POST("/posts/:id", handlerShow)
	return router
}
