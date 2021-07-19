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

func (c *LoginController) fetchInfo() (string, orm.CharField)  {
	la := models.LoginAccount{}
	if err := c.ParseForm(&la); err != nil {
		err.Error()
	}
	name := la.Username
	password := orm.CharField(la.Password)
	return name, password
}

func checkPassword(o orm.Ormer, name string, password orm.CharField) bool {
	correct := orm.CharField(models.GetPassword(o, name))
	return correct==password
}

func (c *LoginController) Get() {
	c.TplName = "login.html"
	c.ManageUserSession()
}

func (c *LoginController) Post() {
	c.TplName = "login.html"
	c.SetUp()
}

func (c *LoginController) SetUp() {
	name, password := c.fetchInfo()
	o := orm.NewOrm()
	if models.UserExists(o, "user", name) {
		if checkPassword(o, name, password) {
			err := models.UpdateLoginTime(o, name)
			if err != nil {
				return
			}
			c.ManageUserSession()
			c.Redirect("../browse", http.StatusFound)
			return
		}
		c.Ctx.Output.Body([]byte("Incorrect password!"))
	} else {
		c.Ctx.Output.Body([]byte("User doesn't exist!"))
	}
}

func (c *LoginController) DataVisualize(name string) {
	o := orm.NewOrm()
	la := models.LoginAccount{Username: name}
	ma := models.AgeMap{}
	female := models.GetNumber(o,"gender", "female")
	male := models.GetNumber(o, "gender", "male")
	agelist := models.GetAges(o)
	ma.CreateAgeMap(agelist)
	c.Data["user"] = &la
	c.Data["female"] = female
	c.Data["male"] = male
	c.Data["age"] = &ma
	c.TplName = "view.html"
}

func (c *LoginController) ManageUserSession() {
	o := orm.NewOrm()
	name, _ := c.fetchInfo()
	if !models.UserExists(o, "user", name) {
		if err := c.DelSession("user"); err != nil {
			err.Error()
			return
		}
		if err := c.DestroySession(); err != nil {
			err.Error()
			return
		}
	} else {
		user := models.GetUserInfo(o, name)
		err := c.SetSession("user", user.Username)
		if err != nil {
			err.Error()
			return
		}
	}
}

func (c *LoginController) Logout() {
	if err := c.DelSession("user"); err != nil {
		err.Error()
		return
	}
	if err := c.DestroySession(); err != nil {
		err.Error()
		return
	}
	c.Redirect("../browse",302)
}







