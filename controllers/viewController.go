package controllers

import beego "github.com/beego/beego/v2/server/web"

type ViewController struct {
	beego.Controller
}

func (c *ViewController) Get() {
	c.TplName = "view.tpl"
}
