package controllers

import (
    "github.com/astaxie/beego"
    "vknotifier/models"
)

type LoginController struct {
    beego.Controller
}

func (c *LoginController) Get() {
        c.TplName = "login.html"
}

func (c *LoginController) Post() {
    var login, password string
    c.Ctx.Input.Bind(&login, "login")
    c.Ctx.Input.Bind(&password, "password")
    result := models.CheckLoginPassword(login, password)

    if result == true {
        c.SetSession("login", login)
        c.Redirect("/", 302)
    } else {
        c.Redirect("/login", 302)
    }
}
