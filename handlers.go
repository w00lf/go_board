package main

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/eknkc/amber"
  "github.com/julienschmidt/httprouter"
  "fmt"
)

var db gorm.DB = inititalizeDb()

func handlerIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  if r.Method == "POST" {
    savePost(*r)
  }
  renderIndex(w)
}

func handlerShow(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  renderShow(w)
}

//template.ParseFiles("tmpl/header.html", "tmpl/form.html", "tmpl/index.html", "tmpl/footer.html") //open and parse a template text file
// t.ExecuteTemplate(w, "header", nil)
// t.ExecuteTemplate(w, "form", nil)
// t.ExecuteTemplate(w, "index", nil)
// t.ExecuteTemplate(w, "footer", nil)

func renderIndex(w http.ResponseWriter) {
  t, err :=  amber.CompileFile("tmpl/index.amber", amber.DefaultOptions) 
  var recentThreads []Thread
  db.Find(&recentThreads)

  data := struct {
    Title string
    Threads []Thread
  }{
    Title: "My page",
    Threads: recentThreads,
  }
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, data)
}

func renderShow(w http.ResponseWriter) {
  t, err :=  amber.CompileFile("tmpl/show.amber", amber.DefaultOptions) 
  var thread Thread
  db.First(&thread, 10)

  data := struct {
    Title string
    Thread Thread
  }{
    Title: "My page",
    Thread: thread,
  }
  if err != nil {
    fmt.Println(err)
  }
  t.Execute(w, data)
}

func savePost(r http.Request) {
  post := Post{ Title: r.FormValue("title"), Body: r.FormValue("body") }
  db.Create(&post)
}