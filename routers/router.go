package routers

import (
	"dataVis/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/register", &controllers.RegController{})
    beego.Router("/login", &controllers.LoginController{})
}
