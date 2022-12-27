## 技术栈

gin+gorm+ini

## 支持自动热重载
```
go install github.com/cosmtrek/air@latest
```

在控制台用`air`命令启动项目

## 正常启动

本地: `go run main.go`    
测试: `go run main.go stg`    
正式: `go run main.go prod`

## 端口

28088

## 发布版本

需要修改`VERSION`文件

## 安装swag
`go get -u github.com/swaggo/swag/cmd/swag`

## 提交代码

运行以下命令重新生成swagger api并整理mod依赖
`swag init -g cmd/api/main.go && go mod tidy`

## 提交规范

| 前缀         | 说明                                        |
|------------|-------------------------------------------|
| `chore`    | 杂項、如build/library/package/config/release等 |
| `ci`       | ci相关修正                                    |
| `docs`     | README等文件修改                               |
| `feat`     | 新功能增加、对应branch feature                    |
| `fix`      | bug修正、对应branch hotfix                     |
| `perf`     | 效能改善相关修正                                  |
| `refactor` | improve/clean/重构 等相关task                  |
| `revert`   | 修正取消                                      |
| `style `   | 空白, 换行, 注解等相关代码排版                         |
| `test`     | 测试相关                                      |