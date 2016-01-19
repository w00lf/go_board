package main

import (
  "net/http"
  "log"
)

func main() {
  fs := http.FileServer(http.Dir("assets"))
  http.Handle("/assets", fs)

  log.Fatal(http.ListenAndServe(":8080", initializeRouter()))
}