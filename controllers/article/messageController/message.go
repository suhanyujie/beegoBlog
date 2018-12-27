package messageController

import "github.com/astaxie/beego"

type MessageController struct {
	beego.Controller
}

// todo
func (_this *MessageController) Index()  {
	_this.TplName = "article/message.html"
}
