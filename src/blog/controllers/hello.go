package controllers

import (
	"github.com/astaxie/beego"
)

type HelloController struct {
	beego.Controller
}

func (hello *HelloController) Get()  {
	hello.Ctx.WriteString("hello go")
}

