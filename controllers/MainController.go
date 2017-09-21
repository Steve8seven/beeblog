package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beeblog/models"
	"strconv"
	"fmt"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Get() {

	categoryParam := this.Input().Get("category")

	fmt.Println(categoryParam)

	o := orm.NewOrm()
	var Articles []*models.Article

	if categoryParam == "" {
		o.QueryTable("article").
			Filter("status", 1).
			OrderBy("-updated").
			All(&Articles, "id","title","updated","views","sammed", "category")
	}else {
		o.QueryTable("article").
			Filter("status", 1).
			Filter("category_id", categoryParam).
			OrderBy("-updated").
			All(&Articles, "id","title","updated","views","sammed", "category")
	}

	this.Data["Article"] = Articles

	var category []*models.Category

	o.QueryTable("category").All(&category)

	fmt.Print(category)
	this.Data["Category"] = category

	this.TplName = "home.html"
}

func (this *MainController) GotoBlogPage() {

	idstr := this.Input().Get("id")

	id, err0 := strconv.ParseInt(idstr, 10, 64)

	if err0 != nil {
		beego.Error("转换失败")
		this.Data["json"] = "转换失败 /(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}
	o := orm.NewOrm()
	article := models.Article{Id: id}
	err := o.Read(&article)

	if err == orm.ErrNoRows {
		beego.Error("没有可以更新的数据", id)
		this.Data["json"] = "没有可以更新的数据 /(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	_, err1 := o.Raw("UPDATE article SET views = views+1 WHERE status = 1 AND id = ?", id).Exec()
	if err1 != nil {
		beego.Error("更新人数错误")
	}

	var maps1 []orm.Params
	var maps2 []orm.Params
	num, err := o.Raw("SELECT MIN(id) prid FROM article WHERE status = 1 and id > ?", id).Values(&maps1)
	var prid interface{}
	if err == nil && num > 0 {
		prid = maps1[0]["prid"]
	}

	num, err = o.Raw("SELECT  MAX(id) neid FROM article where status = 1 and id < ?", id).Values(&maps2)

	var neid interface{}
	if err == nil && num > 0 {
		neid = maps2[0]["neid"]
	}


	this.Data["article_html"] = article.Html_content
	this.Data["title"] = article.Title
	this.Data["category"] = article.Category
	this.Data["sammed"] = article.Sammed
	this.Data["author"] = article.Author
	this.Data["updated"] = article.Updated
	this.Data["prid"] = prid
	this.Data["neid"] = neid
	this.Data["Id"] = id
	this.TplName = "myblog.html"
}
