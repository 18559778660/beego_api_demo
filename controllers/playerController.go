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

// EditPlayer 编辑玩家信息
func (c *PlayerController) EditPlayer() {
	// 1. 获取 URL 中的 id 参数，例如 /edit/player?id=1
	idStr := c.GetString("id") // 获取查询参数
	if idStr == "" {
		c.Data["json"] = PlayerResponse{
			Message: "缺少 id 参数",
			Data:    nil,
			Code:    400,
		}
		c.ServeJSON()
		return
	}

	// 2. 字符串转整数
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = PlayerResponse{
			Message: "无效的 ID",
			Data:    nil,
			Code:    400,
		}
		c.ServeJSON()
		return
	}

	// 3. 检查玩家是否存在
	player, exists := palyers[id]
	if !exists {
		c.Data["json"] = PlayerResponse{
			Message: "玩家不存在",
			Data:    nil,
			Code:    404,
		}
		c.ServeJSON()
		return
	}

	// 4. 解析请求体中的更新数据
	var updateData Player
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &updateData)
	if err != nil {
		c.Data["json"] = PlayerResponse{
			Message: "参数解析失败",
			Data:    nil,
			Code:    400,
		}
		c.ServeJSON()
		return
	}

	// 5. 更新数据（只更新非空字段）- 学习点：条件赋值
	if updateData.Name != "" {
		player.Name = updateData.Name // 变量赋值
	}
	if updateData.Age != 0 {
		player.Age = updateData.Age // 条件赋值
	}

	// 6. 保存更新后的数据 - 学习点：map 的赋值
	palyers[id] = player

	// 7. 返回成功响应
	c.Data["json"] = PlayerResponse{
		Message: "更新玩家成功",
		Data:    player, // 返回更新后的完整数据
		Code:    200,
	}
	c.ServeJSON()
}
