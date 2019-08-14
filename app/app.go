package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Post is a model for a blog post
type Post struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

var posts []Post

type App struct {
	Router *mux.Router
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(posts)
}

func createPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var newPost Post
	json.NewDecoder(req.Body).Decode(&newPost)
	newPost.ID = strconv.Itoa(len(posts) + 1)
	posts = append(posts, newPost)
	json.NewEncoder(res).Encode(newPost)
}

func (a *App) Init() {
	a.Router = mux.NewRouter()
	a.setRouters()
}

func (a *App) setRouters() {

	a.Router.HandleFunc("/post", createPost).Methods("POST")
	a.Router.HandleFunc("/post", getPosts).Methods("GET")
}

func (a *App) Run(host string) {
	fmt.Printf("Listening on port %s....", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}
