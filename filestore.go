package filestorehelper

import (
	"os"
	"path/filepath"
)

// SaveBase64DataToFile 保存 base64 数据到文件，文件路径基于 MIME 类型自动确定扩展名
func SaveBase64DataToFile(data string, baseDirectory, baseFilename string) (string, error) {
	// 解析 MIME 类型并获取扩展名
	mimeType, decodedData, err := ExtractDataFromBase64(data)
	if err != nil {
		return "", err
	}

	var ext string
	if ext, err = GetExtensionFromMimeType(mimeType); err != nil {
		return "", err
	}

	return SaveDataToFile(decodedData, baseDirectory, baseFilename, ext)
}

// SaveDataToFile 保存数据到文件，文件路径基于 MIME 类型自动确定扩展名
func SaveDataToFile(data []byte, baseDirectory, baseFilename, ext string) (string, error) {
	// 构建完整的文件路径
	filename := filepath.Join(baseDirectory, baseFilename+ext)

	// 确保目录存在
	if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return "", err
	}

	// 写入文件
	return filename, os.WriteFile(filename, data, 0644)
}
