// 処理の流れ
// 1.JSONを格納する構造体を作成する
// 2.JSONをデコードするデコーダを生成する
// 3.JSONを順次処理して構造体にデコードする
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	jsonFile, err := os.Open("post.json")
	if err != nil {
		fmt.Println("Error opening JSON file:", err)
		return
	}
	defer jsonFile.Close()

	// 1.JSONデータからデコーダを生成する
	decoder := json.NewDecoder(jsonFile)
	// 2.EOFが検出されるまで繰り返す
	for {
		var post Post
		// 3.JSONデータをデコードし構造体に収納する
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error decoding JSON:", err)
			return
		}
		fmt.Println(post)
	}
}
