package controllers

import (
	"encoding/json"
	"math/rand"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type PlayerController struct {
	beego.Controller
}

type Player struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type PlayerResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Code    int         `json:"code"`
}

var palyers = make(map[int]Player)
var num = 1

func (c *PlayerController) CreatePlayer() {
	var player Player
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &player)
	if err != nil {
		c.Data["json"] = PlayerResponse{
			Message: "创建用户失败,请检查参数",
			Data:    nil,
			Code:    500,
		}
		c.ServeJSON()
		return
	}

	player.ID = num

	if player.Name == "" {
		player.Name = "player" + strconv.Itoa(num)
	}

	if player.Age == 0 {
		player.Age = rand.Intn(100) // ✅ 自动种子，高质量，每次不同
	}

	palyers[num] = player
	num++

	c.Data["json"] = PlayerResponse{
		Message: "创建用户成功",
		Data:    palyers,
		Code:    200,
	}

	c.ServeJSON()

}
