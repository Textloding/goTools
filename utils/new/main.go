package new

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"os"
	"strings"

	"github.com/lxn/walk"
)

func main() {
	var mw *walk.MainWindow
	var ipLabel *walk.Label
	// 加载嵌入的资源图标
	icon, _ := walk.NewIconFromResourceId(2) // IDI_ICON1 默认为资源ID 1

	err := MainWindow{
		AssignTo: &mw,
		Title:    "获取IP地址",
		Icon:     icon, // 设置图标
		Size:     Size{Width: 300, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Label{AssignTo: &ipLabel, Text: "IP地址显示区域"},
			PushButton{
				Text: "获取公网IP地址",
				OnClicked: func() {
					ip := getPublicIP()
					ipLabel.SetText(ip)
				},
			},
			PushButton{
				Text: "获取局域网IP地址",
				OnClicked: func() {
					ip := getLocalIP()
					ipLabel.SetText(ip)
				},
			},
			PushButton{
				Text: "获取代理IP地址",
				OnClicked: func() {
					ip := getProxyIP()
					ipLabel.SetText(ip)
				},
			},
			PushButton{
				Text: "获取IPv6地址",
				OnClicked: func() {
					ip := getIPv6()
					ipLabel.SetText(ip)
				},
			},
		},
	}.Create()

	if err != nil {
		fmt.Println("创建主窗口失败:", err)
		return
	}

	mw.Run()
}

func getPublicIP() string {
	resp, err := http.Get("http://ipinfo.io/ip")
	if err != nil {
		return "Error: " + err.Error()
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error: " + err.Error()
	}
	return strings.TrimSpace(string(body))
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Error: " + err.Error()
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return "No IP found"
}

func getProxyIP() string {
	proxyIP := os.Getenv("HTTP_PROXY")
	if proxyIP == "" {
		return "No Proxy IP found"
	}
	return proxyIP
}

func getIPv6() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Error: " + err.Error()
	}

	for _, addr := range addrs {
		if ipNet, ok := addr.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To16() != nil && ipNet.IP.To4() == nil {
				return ipNet.IP.String()
			}
		}
	}
	return "No IPv6 found"
}
