package routers

import (
    "vknotifier/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/login", &controllers.LoginController{})
    beego.Router("/", &controllers.MainController{})
    beego.Router("/communities", &controllers.CommunitiesController{})
    beego.Router("/communities/add", &controllers.CommunitiesAddController{})
    beego.Router("/communities/edit", &controllers.CommunitiesEditController{})
    beego.Router("/communities/delete", &controllers.CommunitiesDelController{})
    beego.Router("/vkCallback", &controllers.VkCallbackController{})
}
