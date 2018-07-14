package main

import "github.com/astaxie/beego"

type MainController struct{
	beego.Controller
}

func(this * MainController)Get(){
	this.Ctx.WriteString("hello xu!")
}

func main(){
	beego.Router("/hi",&MainController{})
	beego.Run()
}
