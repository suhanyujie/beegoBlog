package routers

import (
	"beegoBlog/controllers"
	"github.com/astaxie/beego"
	"beegoBlog/controllers/article/articleController"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/article", &articleController.ArticleController{})
}
