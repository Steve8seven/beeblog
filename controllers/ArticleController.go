package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"beeblog/models"
	"time"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	check := checkAccount(this.Ctx)
	if check{
		this.TplName = "article_edit.html"
	}else {
		this.Redirect("/login", 301)
	}
	return
}

func (this *ArticleController) Post() {

}

func (this *ArticleController) SaveArticle()  {
	ck, err := this.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error("无法获取cookie里的信息。")
		this.Data["json"] = "登录过期了吧，cookie里没有你的信息/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}
	mdcode := this.Input().Get("mdcode")
	htmlcode := this.Input().Get("htmlcode")
	category := this.Input().Get("category")
	title := this.Input().Get("title")
	sammed := this.Input().Get("sammed")
	uname := ck.Value;

	now_time := time.Now().Format("2006-01-02 15:04:05")

	o := orm.NewOrm()
	article :=new(models.Article)
	article.Author = uname
	article.Html_content = htmlcode
	article.Md_content = mdcode
	article.Title = title
	article.Created = now_time
	article.Updated = now_time
	article.Views = 0
	article.Status = 1
	article.Category = category
	article.Sammed =sammed
	//创建标签
	category_bean := models.Category{Title: category}
	category_bean.Views = 0
	category_bean.Created = now_time

	user := models.BeegoBlogUser{UserName: uname}

	err0 := o.Read(&user, "user_name");
	if err0 == orm.ErrNoRows {
		beego.Error("查找用户时报错",err)
		this.Data["json"] = "查找用户时报错了/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}else if err0 == orm.ErrMissPK {
		beego.Error("找不到主键",err)
		this.Data["json"] = "找不到主键了/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	if created, id, err := o.ReadOrCreate(&category_bean,"Title"); err == nil{
		if created {
			beego.Informational("创建一个新categoryId:", id)
		} else {
			beego.Informational("已经存在的categoryId:", id)
		}
		article.Category_id = id
	}
	 _, err1 := o.Insert(article)
	if err1 != nil {
		beego.Error("插入时错误",err)
		this.Data["json"] = "插入时错误/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	this.Data["json"] = "ok"
	this.ServeJSON()
}

func (this *ArticleController) MyArticle(){
	ck, err := this.Ctx.Request.Cookie("uname")
	pageNum := this.Input().Get("p")

	pageNumInt,err0 := strconv.Atoi(pageNum);

	if err0 != nil {
		beego.Error("页面获取失败")
		this.Data["json"] = "页面获取失败/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	if err != nil {
		beego.Error("无法获取cookie里的信息。")
		this.Data["json"] = "登录过期了吧，cookie里没有你的信息/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}
	uname := ck.Value
	o := orm.NewOrm()
	var Articles []*models.Article
	o.QueryTable("article").
		Filter("author", uname).
		Filter("status", 1).
		Limit(pageNumInt*5).
		OrderBy("-updated").
		All(&Articles, "id","title","updated","views")
	this.Data["Article"] = Articles
	this.Data["pageNum"] = pageNum
	this.TplName = "my_article.html"

}

func (this *ArticleController) GotoSuccess(){
	this.TplName = "success.html"
}

func (this *ArticleController) GoteUpdatePage(){
	idstr :=this.Input().Get("Id")
	o :=orm.NewOrm()

	id, err0 := strconv.ParseInt(idstr, 10, 64)

	if err0 != nil {
		beego.Error("转换失败")
		this.Data["json"] = "转换失败 /(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	articles  := []models.ArticleDO{}
	qb, _ := orm.NewQueryBuilder("mysql")

	qb.Select("article.title","sammed","md_content","updated","category.title ctitle","category.id cid","article.id").
		From("article").LeftJoin("category").
		On("category.id = category_id").
		Where("article.id = ?").Limit(1)
	sql := qb.String()

	// 执行 SQL 语句
	o.Raw(sql, id).QueryRows(&articles)

	article := articles[0]
	this.Data["article_md"] = article.Md_content
	this.Data["title"] = article.Title
	this.Data["category"] = article.Ctitle
	this.Data["sammed"] = article.Sammed
	this.Data["updated"] = article.Updated
	this.Data["Id"] = id

	this.TplName = "my_update_article.html";
}

//更新数据
func (this *ArticleController) UpdateArticle(){
	_, err := this.Ctx.Request.Cookie("uname")
	if err != nil {
		beego.Error("无法获取cookie里的信息。")
		this.Data["json"] = "登录过期了吧，cookie里没有你的信息/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}
	mdcode := this.Input().Get("mdcode")
	htmlcode := this.Input().Get("htmlcode")
	category := this.Input().Get("category")
	title := this.Input().Get("title")
	sammed := this.Input().Get("sammed")
	idstr := this.Input().Get("id")

	now_time := time.Now().Format("2006-01-02 15:04:05")

	id, err0 := strconv.ParseInt(idstr, 10, 64)

	if err0 != nil {
		beego.Error("转换失败")
		this.Data["json"] = "转换失败 /(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	o := orm.NewOrm()
	article := models.Article{Id: id}
	err = o.Read(&article)

	if err == orm.ErrNoRows {
		beego.Error("没有可以更新的数据", id)
		this.Data["json"] = "没有可以更新的数据 /(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	article.Html_content = htmlcode
	article.Md_content = mdcode
	article.Title = title
	article.Updated = now_time
	article.Sammed = sammed
	//创建标签
	category_bean := models.Category{Title: category}
	category_bean.Views = 0
	category_bean.Created = now_time

	if created, id, err := o.ReadOrCreate(&category_bean,"Title"); err == nil{
		if created {
			beego.Informational("创建一个新categoryId:", id)
		} else {
			beego.Informational("已经存在的categoryId:", id)
		}
		article.Category_id = id
	}
	_, err1 := o.Update(&article, "title","md_content","html_content","updated","category_id","sammed")
	if err1 != nil {
		beego.Error("插入时错误",err)
		this.Data["json"] = "插入时错误/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	this.Data["json"] = "ok"
	this.ServeJSON()
}

func (this *ArticleController)DeleteArticlez()  {
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
		beego.Error("没有可以删除的数据", id)
		this.Data["json"] = "没有可以删除的数据 /(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}

	article.Status = 0

	_, err1 := o.Update(&article,"status")

	if err1 != nil {
		beego.Error("删除时错误",err)
		this.Data["json"] = "删除时错误/(ㄒoㄒ)/~~"
		this.ServeJSON()
		return
	}
	this.Redirect("/admin/myArticle?p=1", 301)
}