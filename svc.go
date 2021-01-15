package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jmoiron/sqlx"
	"github.com/zs5460/my"
)

var db *sqlx.DB

func connect() {
	db = sqlx.MustConnect("mssql", "server=172.16.1.11;port=1433;database=icp;user id=icpread;password=readonly;encrypt=disable")
}

func createTopic(id int) {

	sql := "SELECT cateid,catename,sort FROM topic WHERE cateid=?"
	var topic topic
	err := db.Get(&topic, sql, id)
	if err != nil {
		panic(err)
	}

	cates := []category{}
	sql = "SELECT cateid,catename,sort FROM topic WHERE parentid=?"
	err = db.Select(&cates, sql, id)
	if err != nil {
		panic(err)
	}

	items := []item{}
	sql = "Select a.id,a.title,a.addtime,a.image,a.imagesm,a.resume,a.content,aoc.topicid as cateid From ArticleOfTopic aoc INNER JOIN article a on aoc.articleID=a.id where aoc.topicID in (select cateid from topic where parentid=?) order by a.id desc"

	err = db.Select(&items, sql, id)
	if err != nil {
		panic(err)
	}

	result := result{
		Base:  topic,
		Cates: cates,
		Items: items,
	}

	var b []byte
	if indent {
		b, err = json.MarshalIndent(&result, "  ", "  ")
	} else {
		b, err = json.Marshal(&result)
	}
	if err != nil {
		panic(err)
	}
	err = my.WriteText("data.json", my.BytesToString(b))
	if err != nil {
		panic(err)
	}

	fmt.Println("生成成功！")
}
