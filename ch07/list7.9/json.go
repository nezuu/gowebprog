package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// 1.データに対応する構造体を定義
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
	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println("Error reading JSON data:", err)
		return
	}

	fmt.Println(string(jsonData))
	var post Post
	// 3.json.Unmarshalは、構造体のjsonタグがあればその値を対応するフィールドにマッピングする
	json.Unmarshal(jsonData, &post)
	fmt.Println(post.Id)
	fmt.Println(post.Content)
	fmt.Println(post.Author.Id)
	fmt.Println(post.Author.Name)
	fmt.Println(post.Comments[0].Id)
	fmt.Println(post.Comments[0].Content)
	fmt.Println(post.Comments[0].Author)
}
