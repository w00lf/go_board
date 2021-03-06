package main

import (
	"log"
	"net/http"
	"html/template"
	"strconv"
	"w00lf/go_board/Godeps/_workspace/src/github.com/jinzhu/gorm"
	"w00lf/go_board/Godeps/_workspace/src/github.com/eknkc/amber"
	"w00lf/go_board/Godeps/_workspace/src/github.com/julienschmidt/httprouter"
)

var db = inititalizeDb()

func handlerBoardsIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	renderBoardsIndex(w)
}

func handlerPostsIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	boardID, _ := strconv.Atoi(params.ByName("board_id"))
	var board Board
	request := db.First(&board, boardID)
	if request.Error == gorm.RecordNotFound {
		http.NotFound(w, r)
		return
	}
	renderPostsIndex(w, board)
}

func handlerShow(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	postID, _ := strconv.Atoi(params.ByName("id"))
	renderShow(w, postID)
}

func handlerBoardSave(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	board := Board{Title: r.FormValue("title"), Body: r.FormValue("body")}
	db.Create(&board)

	http.Redirect(w, r, "/boards/" + strconv.Itoa(board.ID), 302)
}

func handlerPostSave(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	boardID, _ := strconv.Atoi(params.ByName("board_id"))
	postID, _ := strconv.Atoi(params.ByName("id"))
	post := Post{
		Title: r.FormValue("title"),
		Body: r.FormValue("body"),
		BoardID: boardID,
		PostID: postID,
	}
	db.Create(&post)
	log.Print(post)
	http.Redirect(w, r, ("/boards/" + strconv.Itoa(boardID) + "/posts/" + strconv.Itoa(postID)), 302)
}

func renderBoardsIndex(w http.ResponseWriter) {
	t, err := amber.CompileFile("tmpl/boards_index.amber", amber.DefaultOptions)
	var boards []Board
	db.Order("created_at desc").Find(&boards)

	data := struct {
		Title string
		Boards []Board
	}{
		Title: "My page",
		Boards: boards,
	}
	if err != nil {
		log.Fatal(err)
	}
	t.Execute(w, data)
}

func createTemplate(templateName string) *template.Template {
	template, err := amber.CompileFile("tmpl/posts_index.amber", amber.DefaultOptions)
	if err != nil {
		log.Fatal(err)	
	}
	return template
}

func renderPostsIndex(w http.ResponseWriter, board Board) {
	var recentPosts []Post
	db.Order("created_at desc").Where("board_id = ?", board.ID).Find(&recentPosts)
	createURL := ("/boards/" + strconv.Itoa(board.ID) + "/posts")

	data := struct {
		Title 		string
		CreateURL string
		Board 		Board
		Posts 		[]Post
	}{
		Title: 			"My page",
		CreateURL: 	createURL,
		Board: 			board,
		Posts: 			recentPosts,
	}

	createTemplate("tmpl/posts_index.amber").Execute(w, data)
}

func renderShow(w http.ResponseWriter, id int) {
	var post Post
	var posts []Post
	db.First(&post, id).Association("Posts").Find(&posts)

	createURL := ("/boards/" + strconv.Itoa(post.BoardID) + "/posts/" + strconv.Itoa(post.ID))

	data := struct {
		Title string
		CreateURL string
		Post  Post
		Posts 	[]Post
	}{
		Title: "My page",
		CreateURL: createURL,
		Post:  post,
		Posts: posts,
	}

	createTemplate("tmpl/posts_show.amber").Execute(w, data)
}
