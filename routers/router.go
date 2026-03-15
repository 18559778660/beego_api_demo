package routers

import (
	"beegoApi/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/login", &controllers.LoginController{})

	// CRUD 路由
	// POST     /api/users    创建用户
	// GET      /api/users    获取所有用户
	// GET      /api/users/:id 获取单个用户
	// PUT      /api/users/:id 更新用户
	// DELETE   /api/users/:id 删除用户
	beego.Router("/api/users", &controllers.UserController{}, "post:Post;get:GetAll")
	beego.Router("/api/users/:id", &controllers.UserController{}, "get:Get;put:Put;delete:Delete")
	beego.Router("/api/us", &controllers.Uc{}, "post:Post")
}
