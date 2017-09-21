package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
)

type Category struct {
	Id 			int64
	Title 		string
	Created 	string 	`orm:"index"`
	Views		int64		`orm:"index"`
	UserId	 	int64
}


type Article struct {
	Id 			int64
	Title		string
	Md_content 	string
	Html_content	string
	Sammed	string
	Category	string
	Created		string	`orm:"index"`
	Updated		string	`orm:"index"`
	Views		int64		`orm:"index"`
	Author 		string
	Category_id 	int64
	Status 			int8
}

type ArticleDO struct {
	Id 			int64
	Cid			int64
	Title		string
	Ctitle		string
	Sammed		string
	Md_content 	string
	Updated		string	`orm:"index"`
	Category_id 	int64
}

type BeegoBlogUser struct {
	Id      int64
	UserName	string
	PassWord	string
	Created		string
}

func RegisterModels()  {
	fmt.Println("start register model")
	orm.RegisterModel(new(Category), new(Article), new(BeegoBlogUser))
}