package main

import (
	"flag"
	"fmt"
)

var topicID int
var indent bool

func main() {
	flag.IntVar(&topicID, "id", 0, "专题ID")
	flag.BoolVar(&indent, "i", false, "是否缩进")
	flag.Parse()

	if topicID == 0 {
		fmt.Println("必须指定专题ID")
		return
	}

	connect()
	defer db.Close()

	createTopic(topicID)
}
