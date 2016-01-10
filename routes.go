package main

import (
  "github.com/julienschmidt/httprouter"
  "net/http"
)

func initializeRouter() http.Handler {
  router := httprouter.New()
  router.GET("/", handlerIndex)
  router.GET("/threads/:name", handlerShow)
  return router
}