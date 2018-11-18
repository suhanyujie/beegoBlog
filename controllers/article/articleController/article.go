package articleController

import (
	"github.com/astaxie/beego"
	"html/template"
	models "beegoBlog/models/articleModel"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	this.Data["xsrfField"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrfToken"] = this.XSRFToken()
	filters := make([]interface{},0)
	filters = append(filters,"is_del", 0)
	articles,total := models.GetList(1,10,filters...)
	article := models.BlogArticles{}
	article.Title = "文章首页"
	this.Data["Articles"] = articles
	this.Data["Total"] = total

	this.TplName = "article/index.html"
}

func (_this *ArticleController) Post() {

}

func (this *ArticleController) GetOneColumn() {
	this.Data["xsrfField"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrfToken"] = this.XSRFToken()
	var condition = &models.ConditionType{
		Column:   "id",
		Operater: "=",
		Value:    "2",
	}
	var param []models.ConditionType
	param = append(param, *condition)
	err, title := models.GetOne(param, "Title")
	if err != nil {
		this.TplName = "common/404.html"
	}
	article := models.BlogArticles{}
	article.Title = title
	this.Data["articleData"] = article

	this.TplName = "article/index.html"
}
