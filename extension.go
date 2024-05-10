package filestorehelper

import (
	"errors"
	"mime"
)

// MimeTypeToExtension 映射表，优先使用自定义的扩展名
var MimeTypeToExtension = map[string]string{
	"image/jpeg":         ".jpg",  // 常用于图片文件
	"image/png":          ".png",  // 常用于图片文件
	"image/gif":          ".gif",  // 常用于动图
	"image/svg+xml":      ".svg",  // 矢量图文件
	"image/webp":         ".webp", // 现代网络图片格式
	"text/plain":         ".txt",  // 纯文本文件
	"text/html":          ".html", // HTML 文件
	"text/css":           ".css",  // 样式表文件
	"text/javascript":    ".js",   // JavaScript 文件
	"application/json":   ".json", // JSON 数据文件
	"application/xml":    ".xml",  // XML 数据文件
	"application/zip":    ".zip",  // 压缩文件
	"application/pdf":    ".pdf",  // PDF 文件
	"application/msword": ".doc",  // 微软 Word 文档
	"application/vnd.openxmlformats-officedocument.wordprocessingml.document": ".docx", // Word 文档 (较新版本)
	"application/vnd.ms-excel": ".xls", // 微软 Excel 文档
	"application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":         ".xlsx", // Excel 文档 (较新版本)
	"application/vnd.ms-powerpoint":                                             ".ppt",  // 微软 PowerPoint 演示文稿
	"application/vnd.openxmlformats-officedocument.presentationml.presentation": ".pptx", // PowerPoint 演示文稿 (较新版本)
	"audio/mpeg":      ".mp3", // MP3 音频文件
	"audio/ogg":       ".ogg", // Ogg 音频文件
	"video/mp4":       ".mp4", // MP4 视频文件
	"video/x-msvideo": ".avi", // AVI 视频文件
}

// GetExtensionFromMimeType 从 MIME 类型获取文件扩展名
func GetExtensionFromMimeType(mimeType string) (ext string, err error) {
	var ok bool
	if ext, ok = MimeTypeToExtension[mimeType]; ok {
		return ext, nil
	}

	var ret []string
	if ret, err = mime.ExtensionsByType(mimeType); err != nil {
		return "", err
	}
	if len(ret) == 0 {
		return "", errors.New("no extension found for MIME type")
	}
	return ret[0], nil // 返回第一个扩展名
}
