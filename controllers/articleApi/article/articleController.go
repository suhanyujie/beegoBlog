package article

import (
	"beegoBlog/libs/myUtils"
	"beegoBlog/models/articleModel"
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
)

type ApiArticleController struct {
	beego.Controller
}

type Response struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
}

func (this *ApiArticleController) Post() {
	inputByte := this.Ctx.Input.CopyBody(1024000)
	var inputMap = make(map[string]string, 0)
	err := json.Unmarshal(inputByte, &inputMap)
	if err != nil {
		this.Data["status"] = 2
		this.Data["msg"] = err.Error()
	}
	title := inputMap["title"]
	content := inputMap["content"]
	curTimStamp := time.Now().Unix()
	curDate := time.Now().Format(myUtils.TIME_FORMAT_STRING)
	articleObj := &articleModel.BlogArticlesCopy1{
		0, 0, 0, title, curTimStamp, curTimStamp, 1, 0, curDate, 1, 0, "",
	}
	contentObj := &articleModel.BlogContent{
		0, 0, content, curDate, curDate,
	}
	_, err = articleModel.InsertArticle(articleObj, contentObj, nil)
	res := &Response{1, "success!"}
	if err != nil {
		res.Status = 2
		res.Msg = err.Error()
	}
	//jsonData, err := json.Marshal(res)
	this.Data["json"] = res
	this.ServeJSON(true)
	//_, err = this.Ctx.ResponseWriter.Write(jsonData)
}
