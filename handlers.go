package main

import (
  "net/http"
  "github.com/jinzhu/gorm"
  "github.com/eknkc/amber"
  "fmt"
)

var db gorm.DB = inititalizeDb()

func indexHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    savePost(*r)
  }
  renderIndex(w)
}

//template.ParseFiles("tmpl/header.html", "tmpl/form.html", "tmpl/index.html", "tmpl/footer.html") //open and parse a template text file
// t.ExecuteTemplate(w, "header", nil)
// t.ExecuteTemplate(w, "form", nil)
// t.ExecuteTemplate(w, "index", nil)
// t.ExecuteTemplate(w, "footer", nil)

func renderIndex(w http.ResponseWriter) {
  t, err :=  amber.CompileFile("tmpl/index.amber", amber.DefaultOptions) 
  var recentPosts []Post
  db.Find(&recentPosts)

  data := struct {
    Title string
    Posts []Post
  }{
    Title: "My page",
    Posts: recentPosts,
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