package routers

import (
	"beeblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/login", &controllers.LoginController{})
	//用户管理界面
    beego.Router("/admin", &controllers.AdminController{})
    beego.Router("/admin/gotoArticle", &controllers.ArticleController{})
    beego.Router("/admin/artcle/save_articlez", &controllers.ArticleController{}, "post:SaveArticle")
    beego.Router("/admin/artcle/update_articlez", &controllers.ArticleController{}, "post:UpdateArticle")
    beego.Router("/admin/myArticle", &controllers.ArticleController{}, "get:MyArticle")
    beego.Router("/admin/gotoSuccess", &controllers.ArticleController{}, "get:GotoSuccess")
    beego.Router("/admin/update_page", &controllers.ArticleController{}, "get:GoteUpdatePage")
    beego.Router("/admin/artcle/delete_articlez", &controllers.ArticleController{}, "get:DeleteArticlez")
    //主页控制
    beego.Router("/goto_my_blog", &controllers.MainController{}, "get:GotoBlogPage")
}
