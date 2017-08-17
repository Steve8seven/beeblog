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
	TopticTime 	time.Time 	`orm:"index"`
	TopticCount	int64
	TopticLastUserId	int64
}

type Toptic struct {
	Id 			int64
	Uid			int64
	Title		string
	Content 	string		`orm:"size(5000)"`
	Attachment	string
	Created		time.Time	`orm:"index"`
	Updated		time.Time	`orm:"index"`
	Views		int64		`orm:"index"`
	Author 		string
}

func RegisterModels()  {
	fmt.Println("start register model")
	orm.RegisterModel(new(Category), new(Toptic))
}