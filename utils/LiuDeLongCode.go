package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
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


// stringTrim 函数用于去除字符串中的非中文字符，只保留中文字符。
// 它使用正则表达式匹配所有非中文字符，并将它们替换为空字符串。
// 
// 参数:
// s string: 需要处理的原始字符串。
//
// 返回值:
// string: 经过处理，只包含中文字符的新字符串。
func stringTrim(s string) string {
	// 正则表达式匹配所有非中文字符
	// 这里的正则表达式[^\u4e00-\u9fa5]+匹配任何不在中文字符范围内的字符序列
	reg, err := regexp.Compile("[^\u4e00-\u9fa5]+")
	if err != nil {
		panic(err) // 编译正则表达式失败，通常不应该发生
	}
	// 使用正则表达式替换匹配到的非中文字符为空字符串
	return reg.ReplaceAllString(s, "")
}


// stringCard 处理身份证字符，去除所有非数字和非X的字符，并转换为大写
func stringCard(s string) (string, error) {
	// 转换为大写
	s = strings.ToUpper(s)

	// 编译正则表达式，匹配所有非数字和非X的字符
	reg, err := regexp.Compile(`[^0-9X]+`)
	if err != nil {
		return "", err // 如果正则表达式编译失败，返回错误
	}

	// 使用正则表达式替换匹配到的字符为空字符串
	s = reg.ReplaceAllString(s, "")

	return s, nil
}
