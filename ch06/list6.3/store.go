// 目的
// posts.csvというファイルの中に変数allPosts内に投稿を読み込む

// 処理の流れ
// csvファイルの書き込み
// 1.関数NewWriterにcsvFile(ファイル)を渡して「ライター」を生成
// 2.各投稿ごとに文字列のスライスを生成
// 3.ライターのメソッドWriteを呼び出して、文字列のスライスをCSVファイルに書き込む
// 4.最後にFlushを呼び出して確実にファイルに読み込む

// csvファイルの読み込み
// 1.ファイルを開く
// 2.関数NewReaderにそのファイルを渡してリーダーを生成
// 3.リーダーのメソッドReadAllを呼び出して全レコードを一度に読み込む

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func main() {
	// CSVファイルの作成
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"},
		Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"},
		Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"},
		Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	// CSVファイルの読み込み
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}
	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)
}
