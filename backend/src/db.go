package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// TODO: refactor with declarative generic functions

type PostSign struct {
	Id   int
	Name string
}

type Post struct {
	PostSign
	Content string
}

func initPostsDb() *sql.DB {
	db, _ := sql.Open("sqlite3", "../ms.db")
	db.Exec("create table if not exists posts (id integer primary key,name text,content text)")
	return db
}

func addPostName(db *sql.DB, name string) int {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into posts (name) values (?)")
	data, err := stmt.Exec(name)
	var id, _ = data.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
	return int(id)
}

func addPostContent(db *sql.DB, id int, content string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare(
		fmt.Sprintf("update posts set content = ? where id=%d", id),
	)
	_, err := stmt.Exec(content)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func addPost(db *sql.DB, name string, content string) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into posts (name, content) values (?,?)")
	_, err := stmt.Exec(name, content)
	if err != nil {
		log.Fatal(err)
	}
	tx.Commit()
}

func getAllPostSigns(db *sql.DB) []PostSign {
	rows, err := db.Query("select id,name from posts")
	if err != nil {
		log.Fatal(err)
	}

	var posts []PostSign

	for rows.Next() {
		var p PostSign
		err = rows.Scan(&p.Id, &p.Name)
		posts = append(posts, p)
	}

	return posts
}

func getPost(db *sql.DB, id int) Post {
	rows, err := db.Query("select id,name from posts")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var p Post
		err = rows.Scan(&p.Id, &p.Name)
		if p.Id == id {
			return p
		}
	}

	return Post{}
}

func getAllPosts(db *sql.DB) []Post {
	rows, err := db.Query("select * from posts")
	if err != nil {
		log.Fatal(err)
	}

	var posts []Post
	for rows.Next() {
		var p Post
		err = rows.Scan(&p.Id, &p.Name, &p.Content)
		posts = append(posts, p)
	}
	return posts
}
