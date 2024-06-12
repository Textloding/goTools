package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

type IpInfo struct {
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"` // Geographical coordinates (latitude, longitude)
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
}

func main() {
	var mw *walk.MainWindow
	var ipLabel *walk.Label
	// 加载嵌入的资源图标
	icon, _ := walk.NewIconFromResourceId(2) // IDI_ICON1 默认为资源ID 1

	err := MainWindow{
		AssignTo: &mw,
		Title:    "获取地址",
		Icon:     icon, // 设置图标
		Size:     Size{Width: 300, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Label{AssignTo: &ipLabel, Text: "地址显示区域"},
			PushButton{
				Text: "获取地址",
				OnClicked: func() {
					ip, _ := getAddress()
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

func getAddress() (string, error) {
	url := "http://ipinfo.io/json" // 使用JSON格式的响应来获取详细地址信息
	resp, err := http.Get(url)
	if err != nil {
		return "Error: " + err.Error(), err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "Error: " + err.Error(), err
	}

	var info IpInfo
	err = json.Unmarshal(body, &info)
	if err != nil {
		return "Error: Unable to parse response", err
	}

	// 根据需要自定义返回的地址格式，这里以城市、国家为例
	address := fmt.Sprintf("%s, %s", info.City, info.Country)
	return strings.TrimSpace(address), nil
}
