# kubeimooc

imooc课程：go语言k8s集群管理工具

开发环境说明：
- go语言版本：go1.18.5 darwin/arm64
- 编译环境：mac
- 开发工具：goland

## 项目的初始化

### web框架的选型
```bash
go get -u github.com/gin-gonic/gin@v1.8.1
```

### 把配置参数分离

```bash
go get github.com/spf13/viper@v1.13.0
```

> 参考文档：https://github.com/spf13/viper

### k8s集成

```bash
go get k8s.io/client-go@v0.20.4
```

> 参考文档：go get k8s.io/client-go@v0.20.4
