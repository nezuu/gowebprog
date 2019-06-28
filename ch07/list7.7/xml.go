package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	// 1.データを入れて構造体を作成する
	post := Post{
		Id:      "1",
		Content: "你好",
		Author: Author{
			Id:   "2",
			Name: "比嘉将吾",
		},
	}

	// 2.構造体を組み替えて(marshal)バイト列のXMLデータにする
	// 第二引数に各行の先頭に付けるプレフィックスで、第三引数はインデント(字下げ)文字です
	output, err := xml.MarshalIndent(&post, "", "\t")
	if err != nil {
		fmt.Println("Error marshalling to XML:", err)
		return
	}
	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(output)), 0644)
	if err != nil {
		fmt.Println("Error writing XML to file:", err)
		return
	}
}
