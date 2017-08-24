package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Category struct {
	Id 			int64
	Title 		string
	Created 	time.Time 	`orm:"index"`
	Views		int64		`orm:"index"`
}


type Article struct {
	Id 			int64
	Title		string
	Md_content 	string		`orm:"size(10000)"`
	Html_content	string	`orm:"size(10000)"`
	Created		string	`orm:"index"`
	Updated		string	`orm:"index"`
	Views		int64		`orm:"index"`
	Author 		string
	Category_id 	int64
}


func RegisterModels()  {
	fmt.Println("start register model")
	orm.RegisterModel(new(Category), new(Article))
}