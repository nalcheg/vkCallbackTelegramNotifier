package controllers

import (
    "github.com/astaxie/beego"
    "vknotifier/models"
)

type CommunitiesDelController struct {
    beego.Controller
}

func (c *CommunitiesDelController) Get() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        var id int
        c.Ctx.Input.Bind(&id, "id")
        models.DelCommunity(id)
        c.Redirect("/communities", 302)
    }
}
