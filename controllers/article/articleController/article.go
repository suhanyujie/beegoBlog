package articleController

import (
	"github.com/astaxie/beego"
	"html/template"
	models "beegoBlog/models/articleModel"
	"time"
	"strconv"
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

//新增文章
func (_this *ArticleController) Post() {
	var (
		newId int64
		error error
		responseMap map[string]string
		currentTime string
	)
	currentTime = time.Now().Format("2006-01-02 15:04:05")
	post := _this.Input()
	newArticle := &models.BlogArticles{
		Title:post.Get("title"),
		Content:post.Get("content"),
		Date:time.Now().Unix(),
		PublishDate:0,
		CreatedAt:currentTime,
		UserId:1,
	}
	newId,error = models.Add(newArticle);
	if error!=nil {
		responseMap["error"] = error.Error()
	}
	responseMap = map[string]string{
		"name":"suhanyu",
		"newId":strconv.FormatInt(newId,10),
	}
	_this.Data["json"] = responseMap

	_this.ServeJSON()
}

//展示一个文章
func (_this *ArticleController) Show() {

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
