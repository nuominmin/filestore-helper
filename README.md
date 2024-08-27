# filestore-helper

该包提供了处理文件存储的辅助函数，将 Base64 编码的数据保存为文件。
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
