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
    beego.Router("/gotoArticle", &controllers.ArticleController{})
}
