// ~スライス~
// 1.int型のスライス
// var s []int

// 2.要素数と容量が10であるint型のスライス
// s := make([]int, 10)

// 3. appendの使い方
// s := []int{1, 2, 3}
// s = append(s, 4)
// s = append(s, 5, 6, 7)
// => s == [1, 2, 3, 4, 5, 6, 7]

// ~マップ(配列)~
// 1.int型のキーとstring型の値を保持するマップ
// var m map[int]string

// 2.マップの代入例
// m := make(map[int]string)

// m[1]  = "US"
// m[81] = "Japan"
// m[86] = "China"

// fmt.Println(m)
// => map[1:US 81:Japan 86:China]

// ~解説~
// 1. PostsByAuthor[SauSheong] = append(Sau Sheong, [1, "Hello World!", "Sau Sheong"](post1のポインタ))
// 2. PostsByAuthor = [SauSheong: [1, "Hello World!", "Sau Sheong"], [4, "Greetings Earthlings!", "Sau Sheong"]]

package main

import (
	"fmt"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

var PostById map[int]*Post
var PostsByAuthor map[string][]*Post

func store(post Post) {
	PostById[post.Id] = &post
	PostsByAuthor[post.Author] = append(PostsByAuthor[post.Author], &post)
}

func main() {

	PostById = make(map[int]*Post)
	PostsByAuthor = make(map[string][]*Post)

	post1 := Post{Id: 1, Content: "Hello World!", Author: "Sau Sheong"}
	post2 := Post{Id: 2, Content: "Bonjour Monde!", Author: "Pierre"}
	post3 := Post{Id: 3, Content: "Hola Mundo!", Author: "Pedro"}
	post4 := Post{Id: 4, Content: "Greetings Earthlings!", Author: "Sau Sheong"}

	store(post1)
	store(post2)
	store(post3)
	store(post4)

	fmt.Println(PostById[1])
	fmt.Println(PostById[2])

	for _, post := range PostsByAuthor["Sau Sheong"] {
		fmt.Println(post)
	}
	for _, post := range PostsByAuthor["Pedro"] {
		fmt.Println(post)
	}
}
