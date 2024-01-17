# Go-template
Golang 模板项目，用于快速构建一个Web项目

## 特性清单

### Prometheus metrics
默认集成Prometheus metrics接口，对外暴露地址为`/api/v1/metrics`

### Swagger
默认集成Swagger，对外暴露地址为`/api/v1/swagger/index.html`
首次运行前先执行`make docs`或者`swag init`生成`docs`目录
> 如果没有swag，可以通过以下命令安装: `go install github.com/swaggo/swag/cmd/swag@latest`

### Database
#### 数据库驱动
默认使用Gorm连接数据库，当前已经集成`MySql`、`Postgresql`和`SQLite`，可在生成`config`时通过`WithDatabase`方法进行配置，如
```go
cfg := config.New(config.WithDatabase("postgres", "host=postgresql-hl user=elune password=elune dbname=elune port=5432 sslmode=disable TimeZone=Asia/Shanghai"))
```
#### 数据库处理
已经封装了基本的数据库增删改查、表大小、分页查询等操作，可参考`pkg/domain/demo`和`pkg/middleware/database`

### Dockerfile
包含一份最小化Dockerfile

### Makefile
包含一份Makefile，可通过`make help`查看支持的操作

### 示例代码
包含一份示例代码，参考`pkg/domain/demo`

### 服务相关
支持优雅停止服务

### 支持自定义日志格式
参考`pkg/log`

### 支持自定义返回码
参考`pkg/response`

### 支持代理前端静态文件
支持代理前端静态文件，将编译好的前端项目防止在`pkg/server/static`目录下即可

