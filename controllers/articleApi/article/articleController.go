package article

import (
	"beegoBlog/libs/myUtils"
	"beegoBlog/models/articleModel"
	"github.com/astaxie/beego"
	"strconv"
	"time"
)

type ApiArticleController struct {
	beego.Controller
}

type Response struct {
	Status int    `json:status`
	Msg    string `json:msg`
}

func (this *ApiArticleController) Post() {
	title := this.Input().Get("title")
	content := this.Input().Get("content")
	curTimStamp := time.Now().Unix()
	curDate := time.Now().Format(myUtils.TIME_FORMAT_STRING)
	articleObj := &articleModel.BlogArticlesCopy1{
		0, 0, 0, title, curTimStamp, curTimStamp, 1, 0, curDate, 1, 0, "",
	}
	contentObj := &articleModel.BlogContent{
		0, 0, content, curDate, curDate,
	}
	articleId, err := articleModel.InsertArticle(*articleObj, *contentObj, nil)
	this.Data["status"] = 1
	this.Data["msg"] = "success!" + strconv.FormatInt(articleId, 11)
	if err != nil {
		this.Data["status"] = 2
		this.Data["msg"] = err.Error()
	}
	this.ServeJSON()
}
