package controllers

import (
    "github.com/astaxie/beego"
)

type MainController struct {
    beego.Controller
}

func (c *MainController) Get() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        c.Data["user"] = user
        c.TplName = "pages/index.html"
    }
}
