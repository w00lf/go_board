package main

import (
	"log"
	"net/http"
	"strconv"
	"w00lf/go_board/Godeps/_workspace/src/github.com/eknkc/amber"
	"w00lf/go_board/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
)

var db = inititalizeDb()

func handlerIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print(r)
	renderIndex(w)
}

func handlerShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print(r)
	renderShow(w)
}

func handlerSave(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	log.Print(params)
	post := Post{Title: r.FormValue("title"), Body: r.FormValue("body")}
	db.Create(&post)

	http.Redirect(w, r, "/posts/" + strconv.Itoa(post.ID), 302)
}

//template.ParseFiles("tmpl/header.html", "tmpl/form.html", "tmpl/index.html", "tmpl/footer.html") //open and parse a template text file
// t.ExecuteTemplate(w, "header", nil)
// t.ExecuteTemplate(w, "form", nil)
// t.ExecuteTemplate(w, "index", nil)
// t.ExecuteTemplate(w, "footer", nil)

func renderIndex(w http.ResponseWriter) {
	t, err := amber.CompileFile("tmpl/index.amber", amber.DefaultOptions)
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
		log.Fatal(err)
	}
	t.Execute(w, data)
}

func renderShow(w http.ResponseWriter) {
	t, err := amber.CompileFile("tmpl/show.amber", amber.DefaultOptions)
	var post Post
	db.First(&post, 10)

	data := struct {
		Title string
		Post  Post
	}{
		Title: "My page",
		Post:  post,
	}
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}
