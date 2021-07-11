package controllers

import (
	"dataVis/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type RegController struct {
	beego.Controller
}

func (c *RegController) Get() {
	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "register.html"
}

func (c *RegController) Post() {
	c.TplName = "register.html"
	ra := models.RegAccount{}
	if err := c.ParseForm(&ra); err != nil {
		err.Error()
	}
	name := ra.Username
	password := orm.CharField(ra.Password)
	o := orm.NewOrm()
	if !models.UserExists(o, name) {
		err := models.InsertUser(o, name, password)
		if err != nil {
			return 
		}
	}
	c.Redirect("../login",http.StatusFound)
}




