# Go-template
Golang 模板项目

该模板包含以下内容：
* Prometheus metrics接口`pkg/server`
* 优雅停止服务`pkg/server`
* 支持Mysql、Postgresql、SQLite数据库`pkg/middleware/database`
* 封装基本的数据库增删改查等操作`pkg/middleware/database`
* 支持自定义日志格式`pkg/log`
* 封装返回码`pkg/response`
* controller层接口`pkg/controller`
* service层接口`pkg/service`
* swag文档`pkg/server` `/api/v1/swagger/index.html`

> 如果没有swag，可以通过以下命令安装: `go install github.com/swaggo/swag/cmd/swag@latest`