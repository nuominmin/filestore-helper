package filestorehelper

import (
	"encoding/base64"
	"errors"
	"os"
	"path/filepath"
	"strings"
)

// ExtractDataFromBase64 从 base64 编码的数据 URL 提取 MIME 类型和纯数据部分
func ExtractDataFromBase64(encodedData string) (mimeType string, data []byte, err error) {
	if !strings.HasPrefix(encodedData, "data:") {
		return "", nil, errors.New("invalid base64 header")
	}

	// 分割 MIME 类型和纯数据
	commaIndex := strings.Index(encodedData, ",")
	if commaIndex == -1 {
		return "", nil, errors.New("invalid base64 data")
	}

	encodedPart := encodedData[commaIndex+1:]

	// 解码 base64 数据
	if data, err = base64.StdEncoding.DecodeString(encodedPart); err != nil {
		return "", nil, err
	}

	return strings.TrimSuffix(encodedData[5:commaIndex], ";base64"), data, nil
}

// SaveDataToFile 保存数据到文件，文件路径基于 MIME 类型自动确定扩展名
func SaveDataToFile(data string, baseDirectory, baseFilename string) (string, error) {
	// 解析 MIME 类型并获取扩展名
	mimeType, decodedData, err := ExtractDataFromBase64(data)
	if err != nil {
		return "", err
	}

	var ext string
	if ext, err = GetExtensionFromMimeType(mimeType); err != nil {
		return "", err
	}

	// 构建完整的文件路径
	filename := filepath.Join(baseDirectory, baseFilename+ext)

	// 确保目录存在
	if err = os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
		return "", err
	}

	// 写入文件
	return filename, os.WriteFile(filename, decodedData, 0644)
}
