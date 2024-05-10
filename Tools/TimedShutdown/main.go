package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	var inTE *walk.LineEdit
	var cb *walk.ComboBox
	var mainWindow *walk.MainWindow
	// 加载嵌入的资源图标
	icon, err := walk.NewIconFromResourceId(2) // IDI_ICON1 默认为资源ID 1
	if err != nil {
		// 处理错误
		panic(err)
	}

	MainWindow{
		AssignTo: &mainWindow,
		Title:    "定时关机",
		Icon:     icon, // 设置图标
		Size:     Size{Width: 200, Height: 200},
		Layout:   VBox{},
		Children: []Widget{
			ComboBox{
				AssignTo: &cb,
				Model:    []string{"设置具体日期和时间", "设置持续时间后关机"},
			},
			Label{Text: "请输入时间（日期时间格式：2024-05-01 15:04 或 持续时间分钟）:"},
			LineEdit{AssignTo: &inTE},
			PushButton{
				Text: "确定",
				OnClicked: func() {
					var cmd *exec.Cmd
					if cb.CurrentIndex() == 0 { // 设置具体日期和时间
						cmd = exec.Command("shutdown", "-s", "-t", calculateSecondsToDate(inTE.Text()))
					} else { // 设置持续时间后关机
						minutes, err := strconv.Atoi(inTE.Text())
						if err != nil {
							walk.MsgBox(mainWindow, "错误", "请选择定时类型并输入有效的数字", walk.MsgBoxIconError)
							return
						}
						cmd = exec.Command("shutdown", "-s", "-t", strconv.Itoa(minutes*60))
					}
					executeCommand(cmd, mainWindow)
					walk.MsgBox(mainWindow, "成功", "关机时间已设置", walk.MsgBoxIconInformation)
				},
			},
			PushButton{
				Text: "取消定时关机",
				OnClicked: func() {
					cmd := exec.Command("shutdown", "-a")
					if err := cmd.Run(); err != nil {
						walk.MsgBox(mainWindow, "失败", "无法取消关机,请确定系统中是否存在定时关机", walk.MsgBoxIconError)
					} else {
						walk.MsgBox(mainWindow, "成功", "关机已取消", walk.MsgBoxIconInformation)
					}
				},
			},
		},
	}.Run()
}

func calculateSecondsToDate(dateTimeStr string) string {
	layout := "2006-01-02 15:04"
	t, err := time.Parse(layout, dateTimeStr)
	if err != nil {
		return "0"
	}
	duration := t.Sub(time.Now()).Seconds()
	return strconv.Itoa(int(duration))
}

func executeCommand(cmd *exec.Cmd, mainWindow *walk.MainWindow) {
	if err := cmd.Start(); err != nil {
		walk.MsgBox(mainWindow, "失败", "无法设置关机时间", walk.MsgBoxIconError)
	}
}
