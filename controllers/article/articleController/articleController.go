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
	err,title := models.GetOne(param,"Title")
	if err!= nil {

		this.TplName = "common/404.html"
	}
	article := models.BlogArticles{}
	article.Title = title
	this.Data["articleData"] = article

	this.TplName = "article/index.html"
}
