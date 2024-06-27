package filestorehelper

import (
	"encoding/base64"
	"errors"
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
