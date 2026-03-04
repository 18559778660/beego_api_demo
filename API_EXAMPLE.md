# Beego CRUD API 使用示例

## 启动服务
```bash
go run main.go
```

服务默认在 8080 端口运行

## API 接口说明

### 1. 创建用户 (Create)
**请求：** POST /api/users
**Content-Type:** application/json

```bash
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"张三","email":"zhangsan@example.com","age":25}'
```

**响应示例：**
```json
{
  "code": 201,
  "message": "创建成功",
  "data": {
    "id": 1,
    "username": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }
}
```

---

### 2. 获取所有用户 (Read All)
**请求：** GET /api/users

```bash
curl http://localhost:8080/api/users
```

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": [
    {
      "id": 1,
      "username": "张三",
      "email": "zhangsan@example.com",
      "age": 25
    },
    {
      "id": 2,
      "username": "李四",
      "email": "lisi@example.com",
      "age": 30
    }
  ]
}
```

---

### 3. 获取单个用户 (Read One)
**请求：** GET /api/users/:id

```bash
curl http://localhost:8080/api/users/1
```

**响应示例：**
```json
{
  "code": 200,
  "message": "获取成功",
  "data": {
    "id": 1,
    "username": "张三",
    "email": "zhangsan@example.com",
    "age": 25
  }
}
```

**用户不存在时：**
```json
{
  "code": 404,
  "message": "用户不存在",
  "data": null
}
```

---

### 4. 更新用户 (Update)
**请求：** PUT /api/users/:id
**Content-Type:** application/json

```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username":"张三更新","age":26}'
```

**响应示例：**
```json
{
  "code": 200,
  "message": "更新成功",
  "data": {
    "id": 1,
    "username": "张三更新",
    "email": "zhangsan@example.com",
    "age": 26
  }
}
```

---

### 5. 删除用户 (Delete)
**请求：** DELETE /api/users/:id

```bash
curl -X DELETE http://localhost:8080/api/users/1
```

**响应示例：**
```json
{
  "code": 200,
  "message": "删除成功",
  "data": null
}
```

---

## 快速测试脚本

### 创建两个用户
```bash
# 创建用户 1
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"张三","email":"zhangsan@example.com","age":25}'

# 创建用户 2
curl -X POST http://localhost:8080/api/users \
  -H "Content-Type: application/json" \
  -d '{"username":"李四","email":"lisi@example.com","age":30}'
```

### 获取所有用户
```bash
curl http://localhost:8080/api/users
```

### 获取 ID 为 1 的用户
```bash
curl http://localhost:8080/api/users/1
```

### 更新 ID 为 1 的用户
```bash
curl -X PUT http://localhost:8080/api/users/1 \
  -H "Content-Type: application/json" \
  -d '{"username":"张三更新","age":26}'
```

### 删除 ID 为 1 的用户
```bash
curl -X DELETE http://localhost:8080/api/users/1
```

---

## 代码要点说明

### 1. 结构体定义
```go
// 用户数据模型
type User struct {
    ID       int    `json:"id"`
    Username string `json:"username"`
    Email    string `json:"email"`
    Age      int    `json:"age"`
}

// 统一响应格式
type Response struct {
    Code    int         `json:"code"`    // 状态码
    Message string      `json:"message"` // 消息
    Data    interface{} `json:"data"`    // 数据（可以是任何类型）
}
```

### 2. 控制器方法
- `Post()` - 处理 POST 请求（创建）
- `Get()` - 处理 GET 请求（获取单个）
- `GetAll()` - 处理 GET 请求（获取所有）
- `Put()` - 处理 PUT 请求（更新）
- `Delete()` - 处理 DELETE 请求（删除）

### 3. 常用操作
- 解析 JSON：`json.Unmarshal(c.Ctx.Input.RequestBody, &user)`
- 获取路径参数：`c.Ctx.Input.Param(":id")`
- 返回 JSON：`c.Data["json"] = responseData; c.ServeJSON()`
