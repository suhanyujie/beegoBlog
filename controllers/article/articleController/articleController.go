package articleController

import (
	"github.com/astaxie/beego"
	"html/template"
	"beegoBlog/models"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get(){
	this.Data["xsrfField"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrfToken"] = this.XSRFToken()
	var condition  = &models.ConditionType{
		Column:"id",
		Operater:"=",
		Value:"2",
	}
	var param []models.ConditionType
	param = append(param, *condition)
	err,data := models.GetRow(param)
	if err!= nil {

		this.TplName = "common/404.html"
	}
	this.Data["articleData"] = data

	this.TplName = "article/index.html"
}
