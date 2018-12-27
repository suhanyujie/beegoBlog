package articleController

import (
	"github.com/astaxie/beego"
	"html/template"
	models "beegoBlog/models/articleModel"
	"time"
	"strconv"
	"beegoBlog/libs/myUtils"
	"fmt"
)

type ArticleController struct {
	beego.Controller
}

func (this *ArticleController) Get() {
	var (
		totalPage int64;
		pageSize int64 = 10;
		nextPage int64;
		page     int64 = 1;
	)
	p := this.Input().Get("p")
	page,_ = strconv.ParseInt(p,10,10)
	this.Data["xsrfField"] = template.HTML(this.XSRFFormHTML())
	this.Data["xsrfToken"] = this.XSRFToken()
	filters := make([]interface{}, 0)
	filters = append(filters, "is_del", 0)
	articles, total := models.GetList(int(page), int(pageSize), filters...)
	//filters = append(filters,"is_count",1)
	article := models.BlogArticles{}
	article.Title = "文章首页"
	this.Data["Articles"] = articles
	this.Data["Total"] = total
	this.Data["page"] = page
	this.Data["pageSize"] = pageSize
	totalPage = total / pageSize
	if total % pageSize > 0 {
		totalPage++
	}
	nextPage = page + 1
	if nextPage >= totalPage {
		nextPage = totalPage
	}
	this.Data["nextPage"] = nextPage
	pagi := myUtils.NewPaginator(this.Ctx.Request, int(pageSize), total)
	this.Data["paginator"] = pagi
	this.Data["Lang"] = "zh-CN"

	this.TplName = "article/index.html"
}

//展示一个文章
func (_this *ArticleController) Show() {
	articleId := _this.Ctx.Input.Param(":id")
	condi := &models.ConditionType{
		"id",
		"=",
		articleId,
	}
	var condiArr []models.ConditionType
	condiArr = append(condiArr, *condi)
	err,article := models.GetRaw(condiArr)
	if err!=nil {
		fmt.Println(err)
	}
	if article.Content != "" {
		article.Content = myUtils.MarkdownToHtml(article.Content)
	}
	_this.Data["article"] = article
	_this.Data["articlePublishTime"] = time.Unix(article.PublishDate, 0).Format("2006-01-02 15:04:05")
	//time.Now().Format("2006-01-02 15:04:05")
	_this.Data["year"] = time.Now().Format("2006")
	_this.TplName = "article/details.html"
}

//新增文章
func (_this *ArticleController) Post() {
	var (
		newId       int64
		error       error
		responseMap map[string]string
		currentTime string
	)
	currentTime = time.Now().Format("2006-01-02 15:04:05")
	post := _this.Input()
	newArticle := &models.BlogArticles{
		Title:       post.Get("title"),
		Content:     post.Get("content"),
		Date:        time.Now().Unix(),
		PublishDate: 0,
		CreatedAt:   currentTime,
		UserId:      1,
	}
	newId, error = models.Add(newArticle);
	if error != nil {
		responseMap["error"] = error.Error()
	}
	responseMap = map[string]string{
		"name":  "suhanyu",
		"newId": strconv.FormatInt(newId, 10),
	}
	_this.Data["json"] = responseMap

	_this.ServeJSON()
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
