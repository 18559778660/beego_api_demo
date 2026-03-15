package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "beego.vip"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Get() {
	// 方法 1：直接写字符串到响应
	c.Ctx.WriteString("Hello, World!")
}

// ==================== 数据模型定义 ====================

// User 用户结构体 - 用于存储数据
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

// Response 统一响应结构体
type Response struct {
	Code    int         `json:"code"`    // 状态码
	Message string      `json:"message"` // 消息
	Data    interface{} `json:"data"`    // 数据
}

// UserController 用户控制器
type UserController struct {
	beego.Controller
}

// ==================== 内存数据存储（模拟数据库） ====================

var users = make(map[int]User)
var nextID = 1

// ==================== CRUD 接口实现 ====================

// Create 创建用户 - POST /api/users
func (c *UserController) Post() {
	var user User
	// 解析请求体中的 JSON 数据
	// 注意：使用 io.ReadAll 直接从 Request.Body 读取，兼容性更好
	bodyBytes, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		fmt.Println("读取 Body 失败:", err)
		c.Data["json"] = Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		c.Data["json"] = Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 简单验证
	if user.Username == "" || user.Email == "" {
		c.Data["json"] = Response{
			Code:    400,
			Message: "用户名和邮箱不能为空",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 分配 ID 并保存
	user.ID = nextID
	users[nextID] = user
	nextID++

	// 返回创建成功的响应
	c.Data["json"] = Response{
		Code:    201,
		Message: "创建成功",
		Data:    user,
	}
	c.ServeJSON()
}

// GetAll 获取所有用户 - GET /api/users
func (c *UserController) GetAll() {
	userList := make([]User, 0)
	for _, user := range users {
		userList = append(userList, user)
	}

	c.Data["json"] = Response{
		Code:    200,
		Message: "获取成功",
		Data:    userList,
	}
	c.ServeJSON()
}

// Get 获取单个用户 - GET /api/users/:id
func (c *UserController) Get() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = Response{
			Code:    400,
			Message: "无效的用户 ID",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	user, exists := users[id]
	if !exists {
		c.Data["json"] = Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	c.Data["json"] = Response{
		Code:    200,
		Message: "获取成功",
		Data:    user,
	}
	c.ServeJSON()
}

// Put 更新用户 - PUT /api/users/:id
func (c *UserController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = Response{
			Code:    400,
			Message: "无效的用户 ID",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 检查用户是否存在
	_, exists := users[id]
	if !exists {
		c.Data["json"] = Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 解析更新数据
	var updateUser User
	err = json.Unmarshal(c.Ctx.Input.RequestBody, &updateUser)
	if err != nil {
		c.Data["json"] = Response{
			Code:    400,
			Message: "请求参数错误",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 更新用户信息
	user := users[id]
	if updateUser.Username != "" {
		user.Username = updateUser.Username
	}
	if updateUser.Email != "" {
		user.Email = updateUser.Email
	}
	if updateUser.Age != 0 {
		user.Age = updateUser.Age
	}
	users[id] = user

	c.Data["json"] = Response{
		Code:    200,
		Message: "更新成功",
		Data:    user,
	}
	c.ServeJSON()
}

// Delete 删除用户 - DELETE /api/users/:id
func (c *UserController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.Data["json"] = Response{
			Code:    400,
			Message: "无效的用户 ID",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 检查用户是否存在
	_, exists := users[id]
	if !exists {
		c.Data["json"] = Response{
			Code:    404,
			Message: "用户不存在",
			Data:    nil,
		}
		c.ServeJSON()
		return
	}

	// 删除用户
	delete(users, id)

	c.Data["json"] = Response{
		Code:    200,
		Message: "删除成功",
		Data:    nil,
	}
	c.ServeJSON()
}
