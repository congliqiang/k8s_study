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

> 参考文档：https://github.com/kubernetes/client-go

### k8s集成

```bash
go get k8s.io/client-go@v0.20.4
```

## 项目接口开发

### Pod管理接口开发
- [x] 命名空间列表接口
- [x] Pod创建
- [x] Pod编辑（更新/升级）
---
- [x] Pod查看-详情
    展示podrequest 数据 用于重新创建
- [x] Pod查看-列表
- [x] Pod删除

接口调优：
1. pod更新会多出一个卷挂载（ serviceAccount ）
- 计算哪些是emptydir volume mount 进行非emptydir过滤
2. 更新 pod 超时 -- pod删除等待时间不确定 改为强制删除
3. pod列表支持关键字搜索

