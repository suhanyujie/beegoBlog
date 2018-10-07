package main

import (
	_ "beegoBlog/routers"
	"github.com/astaxie/beego"
)

func main() {
	blogInit()
	beego.Run()
}

func blogInit() {
	//增加静态文件访问的支持
	beego.SetStaticPath("/layui","views/layui")
	beego.SetStaticPath("/static","views/static")
	//增加html后缀的模板支持
	beego.AddTemplateExt("html")
}

