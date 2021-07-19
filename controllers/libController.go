package controllers

import (
	"dataVis/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type LibraryController struct{
	beego.Controller
}

func (c *LibraryController) Get() {
	c.TplName = "musiclib.html"
	user:= c.GetSession("user")
	username := fmt.Sprintf("%v",user)
	if user == nil {
		c.Data["is_logged_in"]=false
		return
	}
	o := orm.NewOrm()
	u := models.GetUserInfo(o, username)
	c.Data["User"]=u
	c.Data["is_logged_in"]=true
}

func (c *LibraryController) Post() {
	user:= c.GetSession("user").(models.User)
	fmt.Println(user.Username)
	c.TplName = "musiclib.html"
}


