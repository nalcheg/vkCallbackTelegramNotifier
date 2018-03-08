package controllers

import (
    "github.com/astaxie/beego"
    "vknotifier/models"
    "strconv"
    "github.com/astaxie/beego/orm"
)

type CommunitiesAddController struct {
    beego.Controller
}

func (c *CommunitiesAddController) Get() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        c.Data["action"] = "/communities/add"
        c.Data["ButtonText"] = "Добавить"
        c.TplName = "pages/communitiesAddEdit.html"
    }
}

func (c *CommunitiesAddController) Post() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        var o = orm.NewOrm()
        newCommunity := new(models.Community)
        newCommunity.VkId, _ = strconv.Atoi(c.GetString("communityId"))
        newCommunity.Name = c.GetString("communityName")
        newCommunity.CheckResponse = c.GetString("responseCode")
        newCommunity.Type = c.GetString("communityType")
        newCommunity.Key = c.GetString("secretKey")
        newCommunity.CheckResponse = c.GetString("responseCode")
        o.Insert(newCommunity)
        c.Redirect("/communities", 302)
    }
}
