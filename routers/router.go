package routers

import (
	"beegoBlog/controllers/article/articleController"
	"beegoBlog/controllers/article/messageController"
	"beegoBlog/controllers/articleApi/article"
	"github.com/astaxie/beego"
)

func init() {
	//beego.Router("/", &controllers.MainController{})
	//访问根路径也是访问文章首页
	beego.Router("/", &articleController.ArticleController{})
	beego.Router("/article/index", &articleController.ArticleController{})
	beego.Router("/article", &articleController.ArticleController{}, "post:Post")
	beego.Router("/article/detail/:id", &articleController.ArticleController{}, "get:Show")
	beego.Router("/article/message/index", &messageController.MessageController{}, "get:Index")
	beego.Router("/api/article", &article.ApiArticleController{}, "post:Post")
}
