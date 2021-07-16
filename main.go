package main

import (
	_ "dataVis/routers"
	beego "github.com/beego/beego/v2/server/web"
	_ "github.com/beego/beego/v2/server/web/session/redis"
)

func main() {
	beego.Run()
}

