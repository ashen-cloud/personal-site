package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// TODO: read port from envvar
// TODO: abstract json to somewhere

var PORT = ":8080"

type Route struct {
	path string
	cb   func(http.ResponseWriter, *http.Request)
}

func main() {
	var ctx = context.Background()

	var db = initPostsDb()

	var ROUTES = []Route{
		{
			path: "/",
			cb: func(w http.ResponseWriter, r *http.Request) {
				println("in root")
			},
		},
		{
			path: "/posts",
			cb: func(w http.ResponseWriter, r *http.Request) {
				var postSigns = getAllPostSigns(db)
				d, err := json.Marshal(postSigns)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(w, string(d))
			},
		},
		{
			path: "/posts/{id}",
			cb: func(w http.ResponseWriter, r *http.Request) {
				var id, ok = mux.Vars(r)["id"]
				if !ok {
					log.Fatal("No post id")
				}
				var nID, _ = strconv.Atoi(id)
				var post = getPost(db, nID)

				d, err := json.Marshal(post)
				if err != nil {
					log.Fatal(err)
				}
				fmt.Fprintf(w, string(d))
			},
		},
	}

	go botInit(db)

	var router = mux.NewRouter().StrictSlash(true)

	routeAll(ctx, router, &ROUTES)

	log.Fatal(http.ListenAndServe(PORT, router))
}

func routeAll(ctx context.Context, router *mux.Router, routes *[]Route) {
	for _, r := range *routes {
		router.HandleFunc(r.path, r.cb)
	}
}
