package main

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
				w.Header().Set("Access-Control-Allow-Origin", "*")
				d, e := ioutil.ReadFile("../../data.json")
				if e != nil {
					log.Fatal(e)
				}
				println("IN ROOT")
				fmt.Fprintf(w, string(d))
			},
		},
		{
			path: "/posts",
			cb: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
				var postSigns = getAllPostSigns(db)
				d, err := json.Marshal(postSigns)
				if err != nil {
					log.Fatal(err)
				}
				println("IN POSTS")
				fmt.Fprintf(w, string(d))
			},
		},
		{
			path: "/posts/{id}",
			cb: func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Access-Control-Allow-Origin", "*")
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
				println("IN POST")
				fmt.Fprintf(w, string(d))
			},
		},
	}

	go botInit(db)

	var router = mux.NewRouter().StrictSlash(true)

	routeAll(ctx, router, &ROUTES)

	var cert, _ = tls.LoadX509KeyPair(
		"/etc/letsencrypt/live/ashencloud.xyz/fullchain.pem",
		"/etc/letsencrypt/live/ashencloud.xyz/privkey.pem",
	)

	var server = &http.Server{
		Addr:    PORT,
		Handler: router,
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	log.Fatal(server.ListenAndServeTLS("", ""))

	// log.Fatal(http.ListenAndServeTLS(PORT,
	// 	,
	// 	,
	// 	router))
	// log.Fatal(http.ListenAndServe(PORT, router))
}

func routeAll(ctx context.Context, router *mux.Router, routes *[]Route) {
	for _, r := range *routes {
		router.HandleFunc(r.path, r.cb)
	}
}
