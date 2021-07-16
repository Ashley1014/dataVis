package controllers

import (
	"dataVis/models"
	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type RegController struct {
	beego.Controller
}

func (c *RegController) Get() {
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
	gender := ra.Gender
	dob := ra.Dob
	o := orm.NewOrm()
	if !models.UserExists(o, "user", name) {
		err := models.InsertUser(o, name, password, gender, dob)
		if err != nil {
			err.Error()
		} else {c.Ctx.Output.Body([]byte("We're glad to have you here!"))}
	} else {
		c.Ctx.Output.Body([]byte("User already exists!"))
	}
}




