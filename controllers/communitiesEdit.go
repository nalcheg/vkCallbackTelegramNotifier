package controllers

import (
    "github.com/astaxie/beego"
    "vknotifier/models"
    "strconv"
)

type CommunitiesEditController struct {
    beego.Controller
}

func (c *CommunitiesEditController) Get() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        id, _ := strconv.Atoi(c.GetString("id"))
        community := models.GetCommunity(id)

        c.Data["action"] = "/communities/edit"
        c.Data["Id"] = community.Id
        c.Data["VkId"] = community.VkId
        c.Data["Name"] = community.Name
        c.Data["Type"] = community.Type
        c.Data["Key"] = community.Key
        c.Data["CheckResponse"] = community.CheckResponse
        c.Data["ButtonText"] = "Редактировать"
        c.TplName = "pages/communitiesAddEdit.html"
    }
}

func (c *CommunitiesEditController) Post() {
    user := c.GetSession("login")
    if user == nil {
        c.Redirect("/login", 302)
    } else {
        communityRow := make(map[string]string)
        communityRow["Id"] = c.GetString("Id")
        communityRow["communityId"] = c.GetString("communityId")
        communityRow["communityName"] = c.GetString("communityName")
        communityRow["secretKey"] = c.GetString("secretKey")
        communityRow["responseCode"] = c.GetString("responseCode")
        communityRow["communityType"] = c.GetString("communityType")

        models.EditCommunity(communityRow)

        c.Redirect("/communities", 302)
    }
}
