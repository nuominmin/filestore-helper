# filestore-helper

一个轻量级的、可扩展的 Go 语言包，旨在为多种消息通知服务提供统一的接口。通过实现 Notifier 接口，用户可以轻松地集成多个第三方消息通知服务（例如 Lark）到他们的应用程序中
## 安装
你可以使用' go get '来安装这个包:

```sh
go get github.com/nuominmin/filestore-helper
```

## 示例
```go

// base64 转存文件
assetPath, err := filestorehelper.SaveBase64DataToFile(base64Data, "assets", "product")
if err != nil {
    log.Errorf("save asset error: %s", err.Error())
    return
}

// 转存文件
assetPath, err = filestorehelper.SaveDataToFile([]byte("hello world"), "assets", "helloworld", ".txt")
if err != nil {
    log.Errorf("save asset error: %s", err.Error())
    return
}
```