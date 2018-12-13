package routers

import (
	"github.com/astaxie/beego"
	"beegoBlog/controllers/article/articleController"
)

func init() {
    //beego.Router("/", &controllers.MainController{})
    //访问根路径也是访问文章首页
    beego.Router("/", &articleController.ArticleController{})
	beego.Router("/article/index", &articleController.ArticleController{})
	beego.Router("/article", &articleController.ArticleController{},"post:Post")
    beego.Router("/article/detail/:id", &articleController.ArticleController{},"get:Show")
}
