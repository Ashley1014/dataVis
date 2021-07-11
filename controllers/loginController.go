package controllers

import (
	"dataVis/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"net/http"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
}

func (c *LoginController) Post() {
	c.TplName = "login.html"
	la := models.LoginAccount{}
	if err := c.ParseForm(&la); err != nil {
		err.Error()
	}
	name := la.Username
	password := orm.CharField(la.Password)
	o := orm.NewOrm()
	if models.UserExists(o,name) {
		correct := orm.CharField(models.GetPassword(o,name))
		if password == correct {
			err := models.UpdateLoginTime(o, name)
			if err != nil {
				return
			}
			c.Ctx.Output.Body([]byte("登录成功！"))
		}
		c.Ctx.Output.Body([]byte("密码不正确！"))
	} else {
		c.Redirect("../register",http.StatusFound)
	}
}
