package main

import (
	_ "beegoApi/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

