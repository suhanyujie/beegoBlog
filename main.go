package main

import (
	_ "beegoBlog/routers"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

const TIME_FORMAT_STRING = "2006-01-02 15:04:05"

func main() {
	blogInit()
	beego.Run()
}

func blogInit() {
	//增加静态文件访问的支持
	beego.SetStaticPath("/layui", "views/layui")
	beego.SetStaticPath("/static", "static")
	beego.SetStaticPath("/vstatic", "views/static")
	//增加html后缀的模板支持
	beego.AddTemplateExt("html")
	beego.AddFuncMap("i18n", i18n.Tr)
}
