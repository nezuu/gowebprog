//  1. Postgresを起動
//  2. createuser -P -d gwp
//     （passwd gwp）
//  3. creatdb gwp
//  4. psql -U gwp -f setup.sql -d gwp
//    （一番最初は ERRORが表示されるが、そのままでOK)
//  5. go get "github.com/lib/pq"
//  6. go run store.go

package main

import (
	"database/sql"
	"fmt"

	// このパッケージがPostgreSQLのドライバであり、インポートされた時に関数initが実行されて自分を登録する。
	_ "github.com/lib/pq"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var Db *sql.DB

// データベースの接続する
func init() {
	var err error
	Db, err = sql.Open("postgres", "user=gwp dbname=gwp password=gwp sslmode=disable")
	if err != nil {
		panic(err)
	}
}

// get all posts
func Posts(limit int) (posts []Post, err error) {
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil {
		return
	}
	for rows.Next() {
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil {
			return
		}
		posts = append(posts, post)
	}
	rows.Close()
	return
}

// 投稿1件の取得
func GetPost(id int) (post Post, err error) {
	post = Post{}
	err = Db.QueryRow("select id, content, author from posts where id = $1", id).Scan(&post.Id, &post.Content, &post.Author)
	return
}

// 新規投稿の作成
func (post *Post) Create() (err error) {
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	// 上記のステートメントをプリペアドステートメントとして作成している
	//（プリペアドステートメントとは、SQLを最初に用意しておいて、そのあとはクエリ内内のパラメータの値だけを変更してクエリを実行できる機能のこと）
	stmt, err := Db.Prepare(statement)
	if err != nil {
		return
	}
	defer stmt.Close()
	// ステートメントのメソッドQueryRowにレシーバからのデータを渡して、プリペアドステートメントを実行
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

// 投稿の更新
func (post *Post) Update() (err error) {
	_, err = Db.Exec("update posts set content = $2, author = $3 where id = $1", post.Id, post.Content, post.Author)
	return
}

// 投稿の削除
func (post *Post) Delete() (err error) {
	_, err = Db.Exec("delete from posts where id = $1", post.Id)
	return
}

// Delete all posts
func DeleteAll() (err error) {
	_, err = Db.Exec("delete from posts")
	return
}

func main() {
	post := Post{Content: "Hello World!", Author: "Sau Sheong"}

	// Create a post
	fmt.Println(post) // {0 Hello World! Sau Sheong}
	post.Create()
	fmt.Println(post) // {1 Hello World! Sau Sheong}

	// １つの投稿を取得
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost) // {1 Hello World! Sau Sheong}

	// Update the post
	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()

	// Get all posts
	posts, _ := Posts(10)
	fmt.Println(posts) // [{1 Bonjour Monde! Pierre}]

	// Delete the post
	readPost.Delete()

	// Get all posts
	posts, _ = Posts(10)
	fmt.Println(posts) // []

	// Delete all posts
	// DeleteAll()
}
