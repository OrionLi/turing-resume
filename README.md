# 图灵简历(后端)

## 项目简介

本项目是一个基于Go语言和Gin框架开发的Web应用，用于管理创新组和创业组的简历信息。该系统提供了简历信息的创建和更新功能。

## 技术栈

- Go 1.22
- Gin 1.10.0
- GORM 1.25.12
- MySQL

## 项目结构

```
turing-resume/
├── config/
│ └── db.go
├── controllers/
│ ├── cx_controller.go
│ ├── cy_controller.go
│ └── response.go
├── models/
│ ├── cx.go
│ └── cy.go
├── routes/
│ └── router.go
├── .env
├── go.mod
├── go.sum
├── main.go
└── README.md
```

## 功能特性

- 创建和更新创新组（Cx）简历信息
- 创建和更新创业组（Cy）简历信息
- 数据验证
- 自动数据库迁移

## 安装和运行

1. 克隆项目到本地：

   ```
   git clone https://github.com/your-username/turing-resume.git
   cd turing-resume
   ```

2. 安装依赖：

   ```
   go mod tidy
   ```

3. 配置数据库连接：

   编辑 `.env` 文件，修改数据库连接信息。

4. 运行项目：

   ```
   go run main.go
   ```

   服务器将在 `http://localhost:30300` 上运行。

## API 接口

- POST `/fortec`: 创建或更新创新组简历
- POST `/forpla`: 创建或更新创业组简历
