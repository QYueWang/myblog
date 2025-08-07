# Kratos练习项目
参考Kratos的示例项目 [Blog](https://github.com/go-kratos/examples/tree/main/blog)做的Kratos框架练习项目。实现了Kratos示例项目中没有完成的评论和标签模块，使之成为一个完整完善的博客项目，另外，本项目还使用了Kratos的配置、序列化、错误处理、日志、元信息传递、对接监控、服务注册发现、路由负载、中间件等，完整把Kratos框架的内容点融入，供学习参考。

## 相关配置
安装Kratos：
```
go install github.com/go-kratos/kratos/cmd/kratos/v2@latest
```

生成error代码（由于Widonws难安装make，就只能执行执行了）：
```
protoc --proto_path=. --proto_path=./third_party --go_out=paths=source_relative:. --go-errors_out=paths=source_relative:. .\api\v1\error\error.proto
```
若go mod tidy导致go版本号变更，可以查看是哪个依赖导致版本号变更的，更改依赖版本，同时在VSCode中修改GOTOOLCHAIN为当前go版本号：
```
go list -m -f '{{if .GoVersion}}{{.Path}} => Go {{.GoVersion}}{{end}}' all
```

