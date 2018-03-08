package controllers

import (
    "github.com/astaxie/beego"
    "vknotifier/models"
)

type CommunitiesController struct {
    beego.Controller
}

func (c *CommunitiesController) Get() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        c.Data["communities"] = models.GetCommunities()
        c.TplName = "pages/communities.html"
    }
}
