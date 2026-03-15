package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

type U struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

type R struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type Uc struct {
	beego.Controller
}

var us = make(map[int]U)
var next = 1

func (c *Uc) Post() {
	var user U
	// 方式 1：使用 c.Ctx.Input.RequestBody（Beego 封装的方法）
	bodyBytes := c.Ctx.Input.RequestBody

	// 方式 2：使用 io.ReadAll 直接读取（原生方法，更可靠）
	//log.Println(string(bodyBytes))
	// bodyBytes, err := io.ReadAll(c.Ctx.Request.Body)
	// if err != nil {
	// 	c.Data["json"] = R{
	// 		Code:    400,
	// 		Message: "Bad Request",
	// 		Data:    nil,
	// 	}
	// 	c.ServeJSON()
	// 	return
	// }

	err := json.Unmarshal(bodyBytes, &user)
	if err != nil {
		c.Data["json"] = R{
			Code:    400,
			Message: "Bad Request",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	if user.Username == "" || user.Email == "" {
		c.Data["json"] = R{
			Code:    400,
			Message: "用户名和邮箱不能为空",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	user.ID = next
	us[next] = user
	next++

	c.Data["json"] = R{
		Code:    200,
		Message: "User created successfully",
		Data:    user,
	}
	c.ServeJSON()
}
