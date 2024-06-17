package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
)

// downloadZip 远程下载zip文件
// @param url    远程zip文件地址
// @param file   本地下载目录文件名
// @param timeout 超时时间（秒）
func downloadZip(url string, file string, timeout int) (string, error) {
	// 检查文件是否存在，如果存在则删除
	if _, err := os.Stat(file); err == nil {
		if err := os.Remove(file); err != nil {
			return "", err
		}
	}

	// 如果没有提供文件名，则使用URL中的文件名
	if file == "" {
		file = filepath.Base(url)
	}

	// 确保目录存在
	dir := filepath.Dir(file)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		if err := os.MkdirAll(dir, 0755); err != nil {
			return "", err
		}
	}

	// 替换URL中的空格为%20
	url = strings.ReplaceAll(url, " ", "%20")

	// 创建HTTP客户端
	client := &http.Client{
		Timeout: timeout,
	}

	// 获取远程文件
	resp, err := client.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// 创建文件并写入内容
	out, err := os.Create(file)
	if err != nil {
		return "", err
	}
	defer out.Close()

	// 将响应内容写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return "", err
	}

	return file, nil
}
