package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/orm"
	"beeblog/models"
	"time"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	check := checkAccount(this.Ctx)
	fmt.Println(check)
	this.TplName = "article_edit.html"
	if check{
		this.TplName = "article_edit.html"
	}else {
		this.Redirect("/login", 301)
		//this.TplName = "login.html"
	}
	return
}

func (this *ArticleController) Post() {

}

func (this *ArticleController) SaveArticle()  {
	ck, err := this.Ctx.Request.Cookie("uname")
	if err != nil {
		this.Data["json"] = "err1"
		this.ServeJSON()
		return
	}
	mdcode := this.Input().Get("mdcode")
	htmlcode := this.Input().Get("htmlcode")
	category := this.Input().Get("category")
	title := this.Input().Get("title")
	uname := ck.Value;

	o := orm.NewOrm()
	article :=new(models.Article)
	article.Author = uname
	article.Html_content = htmlcode
	article.Md_content = mdcode
	article.Title = title
	article.Created = time.Now().String()
	article.Updated = time.Now().String()
	article.Views = 0
	article.Category_id = 1
	fmt.Println(category)
	//创建标签
	//category_bean := models.Category{Title: category}
	//
	//if created, id, err := o.ReadOrCreate(&category_bean, category); err == nil{
	//	if created {
	//		fmt.Println("New Insert an object. Id:", id)
	//	} else {
	//		fmt.Println("Get an object. Id:", id)
	//	}
	//	article.Category_id = id
	//}
	o.Insert(article)


	this.Data["json"] = "okok"
	this.ServeJSON()
}
