package controllers

import (
	"dataVis/models"
	"fmt"
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
			c.LoginSuccess(name)
			return
		}
		c.Ctx.Output.Body([]byte("密码不正确！"))
	} else {
		c.Redirect("../register",http.StatusFound)
	}
}

func (c *LoginController) LoginSuccess(name string) {
	o := orm.NewOrm()
	la := models.LoginAccount{Username: name}
	ma := models.AgeMap{}
	female := models.GetNumber(o,"gender", "female")
	male := models.GetNumber(o, "gender", "male")
	agelist := models.GetAges(o)
	ma.CreateAgeMap(agelist)
	fmt.Println(ma)
	c.Data["user"] = &la
	c.Data["female"] = female
	c.Data["male"] = male
	c.Data["age"] = &ma
	c.TplName = "view.html"
}
