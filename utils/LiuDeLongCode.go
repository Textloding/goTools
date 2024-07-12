package utils

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"reflect"
	"errors"
	"unicode/utf8"
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

// 假设我们有一个函数来获取请求中的IP地址
func getIPFromRequest() string {
	if ip := os.Getenv("HTTP_CLIENT_IP"); ip != "" {
		return ip
	} else if ip := os.Getenv("HTTP_X_FORWARDED_FOR"); ip != "" {
		return ip
	} else if ip, _, err := net.SplitHostPort(os.Getenv("REMOTE_ADDR")); err == nil {
		return ip
	}
	return "Unknow"
}

// matchIP 检查给定的IP是否在允许的IP列表中，或者是否与通配符模式匹配
func matchIP(ip string, allowedIPs []string) bool {
	ipParts := strings.Split(ip, ".")

	for _, allowedIP := range allowedIPs {
		allowedIPParts := strings.Split(allowedIP, ".")
		match := true
		for i := range allowedIPParts {
			if allowedIPParts[i] != "*" && allowedIPParts[i] != ipParts[i] {
				match = false
				break
			}
		}
		if match {
			return true
		}
	}
	return false
}

// ipTest 函数检查客户端IP是否在允许的IP列表中
func ipTest(allowedIPs []string) bool {
	clientIP := getIPFromRequest()
	return matchIP(clientIP, allowedIPs)
}
// 对象转数组
func objectArray(v interface{}) interface{} {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Slice, reflect.Array:
		result := make([]interface{}, rv.Len())
		for i := 0; i < rv.Len(); i++ {
			result[i] = objectArray(rv.Index(i).Interface())
		}
		return result
	case reflect.Map:
		result := make(map[interface{}]interface{})
		for _, key := range rv.MapKeys() {
			result[key.Interface()] = objectArray(rv.MapIndex(key).Interface())
		}
		return result
	case reflect.Ptr:
		return objectArray(rv.Elem().Interface())
	case reflect.Struct:
		result := make(map[string]interface{})
		t := rv.Type()
		for i := 0; i < rv.NumField(); i++ {
			field := t.Field(i)
			result[field.Name] = objectArray(rv.Field(i).Interface())
		}
		return result
	default:
		return v
	}
}

// mbStrSplit 分割字符串到指定长度的切片。
// 如果splitLength小于1，函数将返回错误。
func mbStrSplit(str string, splitLength int, charset string) ([]string, error) {
	// 默认字符集为UTF-8，目前Go语言只支持UTF-8编码，因此charset参数被忽略。
	if splitLength < 1 {
		return nil, errors.New("split length should be greater than 0")
	}

	var arr []string
	for i := 0; i < len(str); {
		_, size := utf8.DecodeRuneInString(str[i:])
		if size > splitLength {
			return nil, errors.New("rune size is larger than split length")
		}
		r, _ := utf8.DecodeRuneInString(str[i:])
		i += size
		if r == utf8.RuneError {
			continue // 跳过无效的UTF-8编码
		}
		end := i + splitLength
		if end > len(str) {
			end = len(str)
		}
		arr = append(arr, str[i:end])
		i = end
	}
	return arr, nil
}


