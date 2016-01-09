package main

import (
  "html/template"
  "net/http"
  "fmt"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
  if r.Method == "POST" {
    savePost(*r)
  }
  renderIndex(w)
}

func renderIndex(w http.ResponseWriter) {
  t, err := template.ParseFiles("tmpl/header.html", "tmpl/form.html", "tmpl/index.html", "tmpl/footer.html") //open and parse a template text file
  data := struct {
    Title string
    Items []string
  }{
    Title: "My page",
    Items: []string{
      "My photos",
      "My blog",
    },
  }
  if err != nil {
    fmt.Println(err)
  }
  t.ExecuteTemplate(w, "header", nil)
  t.ExecuteTemplate(w, "form", nil)
  t.ExecuteTemplate(w, "index", nil)
  t.ExecuteTemplate(w, "footer", nil)
  t.Execute(w, data)
}

func savePost(r http.Request) {
    title := r.FormValue("title")
    body := r.FormValue("body")
    fmt.Println(title, body)
}


func main() {
    http.HandleFunc("/", indexHandler)
    http.ListenAndServe(":8080", nil)
}