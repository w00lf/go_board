package main

import (
  "log"
	"net/http"
	"w00lf/go_board/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
)

func LogRequest(handler httprouter.Handle) httprouter.Handle {
  return func(w http.ResponseWriter, r *http.Request, params httprouter.Params){
    log.Print(r)
    handler(w, r, params)
  }
}

func initializeRouter() http.Handler {
	router := httprouter.New()
	router.GET("/", LogRequest(handlerBoardsIndex))
	router.POST("/", LogRequest(handlerBoardSave))

  router.GET("/boards/:board_id", LogRequest(handlerPostsIndex))
  router.GET("/boards/:board_id/posts", LogRequest(handlerPostsIndex))

	router.GET("/boards/:board_id/posts/:id", LogRequest(handlerShow))

  router.POST("/boards/:board_id/posts", LogRequest(handlerPostSave))
  router.POST("/boards/:board_id/posts/:id", LogRequest(handlerPostSave))
  router.ServeFiles("/assets/*filepath", http.Dir("assets"))
	return router
}
