package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/extrame/xls"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/tealeg/xlsx"
)

func main() {
	mw, err := createMainWindow()
	if err != nil {
		fmt.Println("创建主窗口失败:", err)
		return
	}

	mw.Run()
}

func createMainWindow() (*walk.MainWindow, error) {
	var mw *walk.MainWindow
	var inColumn, inSheetName, inStartRow *walk.LineEdit

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "选择Excel文件",
		Size:     Size{Width: 400, Height: 300},
		Layout:   VBox{},
		Children: []Widget{
			Label{Text: "输入工作表名称:"},
			LineEdit{AssignTo: &inSheetName},
			Label{Text: "预定义输入列 (如 A, B, ... ; * 表示所有列):"},
			LineEdit{AssignTo: &inColumn},
			Label{Text: "定义起始行数 (如 1 表示第一行):"},
			LineEdit{AssignTo: &inStartRow},
			PushButton{
				Text: "选择文件",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Filter = "Excel 文件 (*.xlsx;*.xls)|*.xlsx;*.xls"
					if ok, err := dlg.ShowOpen(mw); err != nil {
						walk.MsgBox(mw, "错误", "无法打开文件对话框: "+err.Error(), walk.MsgBoxIconError)
					} else if ok {
						startRow, _ := strconv.Atoi(inStartRow.Text())
						processFile(dlg.FilePath, inSheetName.Text(), inColumn.Text(), startRow)
					}
				},
			},
		},
	}).Create(); err != nil {
		return nil, err
	}

	return mw, nil
}

func processFile(filePath, sheetName, columns string, startRow int) {
	fmt.Printf("处理文件: %s 在工作表 %s 从行 %d 开始\n", filePath, sheetName, startRow)
	if strings.HasSuffix(filePath, ".xlsx") {
		processXLSX(filePath, sheetName, columns, startRow)
	} else if strings.HasSuffix(filePath, ".xls") {
		processXLS(filePath, sheetName, columns, startRow)
	}
}

func processXLSX(filename, sheetName, columns string, startRow int) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("无法打开文件: %v\n", err)
		return
	}
	sheet, ok := xlFile.Sheet[sheetName]
	if !ok {
		fmt.Printf("未找到名为 %s 的工作表\n", sheetName)
		return
	}

	data := map[string]interface{}{
		"province": map[string]string{},
		"city":     map[string]map[string]string{},
		"area":     map[string]map[string]string{},
	}

	for i, row := range sheet.Rows {
		if i < startRow-1 || len(row.Cells) < 9 {
			continue
		}
		provinceName := row.Cells[3].String()
		provinceCode := row.Cells[4].String()
		cityName := row.Cells[5].String()
		cityCode := row.Cells[6].String()
		areaName := row.Cells[7].String()
		areaCode := row.Cells[8].String()

		if _, exists := data["province"].(map[string]string)[provinceName]; !exists {
			data["province"].(map[string]string)[provinceName] = provinceCode
		}
		if _, exists := data["city"].(map[string]map[string]string)[cityName]; !exists {
			data["city"].(map[string]map[string]string)[cityName] = map[string]string{
				"code":   cityCode,
				"father": provinceCode,
			}
		}
		data["area"].(map[string]map[string]string)[areaCode] = map[string]string{
			"name":   areaName,
			"father": cityCode,
		}
	}
	formattedData := formatDataForClipboard(data)
	fmt.Println("Formatted data ready for clipboard:")
	fmt.Println(formattedData)
	copyToClipboard(formattedData)
}

func processXLS(filename, sheetName string, cloumns string, startRow int) {
	xlFile, err := xls.Open(filename, "utf-8")
	if err != nil {
		fmt.Printf("无法打开文件: %v\n", err)
		return
	}

	sheet := getSheetByName(xlFile, sheetName)
	if sheet == nil {
		fmt.Printf("未找到名为 '%s' 的工作表\n", sheetName)
		return
	}

	// 使用切片存储省市区数据
	provinces := []map[string]string{}
	cities := []map[string]string{}
	areas := []map[string]string{}

	for i := startRow - 1; i <= int(sheet.MaxRow); i++ {

		row := sheet.Row(i)
		if row == nil || row.LastCol() < 9 {
			fmt.Printf("行 %d 是空数据, skipping\n", i+1)
			continue
		}
		if row.Col(3) == "" || row.Col(4) == "" {
			fmt.Printf("行数据为空 %d\n", i+1)
		}

		// 打印每一行的数据，确认数据读取无误
		fmt.Printf("行 %d: 省名称: '%s', 省代码: '%s'\n", i+1, row.Col(3), row.Col(4))

		// 收集数据
		provinces = append(provinces, map[string]string{"name": row.Col(3), "code": row.Col(4)})
		cities = append(cities, map[string]string{"name": row.Col(5), "code": row.Col(6), "father": row.Col(4)})
		areas = append(areas, map[string]string{"name": row.Col(7), "code": row.Col(8), "father": row.Col(6)})
	}
	println("省份数据.")
	println(provinces)

	// 数据去重
	uniqueProvinces := uniqueEntries(provinces)
	uniqueCities := uniqueEntries(cities)
	uniqueAreas := uniqueEntries(areas)

	// 格式化数据并复制到剪贴板
	formattedData := formatDataForClipboardXLS(uniqueProvinces, uniqueCities, uniqueAreas)
	copyToClipboard(formattedData)

	fmt.Println("Data copied to clipboard successfully.")
}

// 去重函数
func uniqueEntries(entries []map[string]string) map[string]map[string]string {
	unique := make(map[string]map[string]string)
	//fmt.Println("初始值:")
	//for index, entry := range entries {
	//	fmt.Printf("Entry %d: Name: '%s', Code: '%s'\n", index+1, entry["name"], entry["code"])
	//}
	for _, entry := range entries {
		// 标准化数据
		code := strings.TrimSpace(entry["code"])
		name := strings.TrimSpace(entry["name"])

		// 以名称和代码的组合作为键，确保唯一性
		key := fmt.Sprintf("%s|%s", name, code)
		if _, exists := unique[key]; !exists {
			unique[key] = map[string]string{"name": name, "code": code}
		}
	}
	return unique
}

// 数据格式化函数
func formatDataForClipboardXLS(provinces, cities, areas map[string]map[string]string) string {
	var provinceParts, cityParts, areaParts []string
	for _, p := range provinces {
		provinceParts = append(provinceParts, fmt.Sprintf("\"%s\"=>\"%s\"", p["name"], p["code"]))
	}
	for _, c := range cities {
		cityParts = append(cityParts, fmt.Sprintf("\"%s\"=>[\"code\"=>\"%s\",\"father\"=>\"%s\"]", c["name"], c["code"], c["father"]))
	}
	for _, a := range areas {
		areaParts = append(areaParts, fmt.Sprintf("\"%s\"=>[\"name\"=>\"%s\",\"father\"=>\"%s\"]", a["code"], a["name"], a["father"]))
	}

	// 组合最终的字符串
	return fmt.Sprintf("$data = [\n    \"province\"=>[\n    %s,\n    ],\n    \"city\"=>[\n        %s,\n        ],\n    \"area\"=>[\n        %s,\n        ],\n    ];",
		strings.Join(provinceParts, ",\n    "),
		strings.Join(cityParts, ",\n        "),
		strings.Join(areaParts, ",\n        "))
}

func formatDataForClipboard(data map[string]interface{}) string {
	// 省份数据
	provinceParts := []string{}
	for name, code := range data["province"].(map[string]string) {
		provinceParts = append(provinceParts, fmt.Sprintf("\"%s\"=>\"%s\"", name, code))
	}
	provinceStr := strings.Join(provinceParts, ",\n    ")

	// 城市数据
	cityParts := []string{}
	for name, details := range data["city"].(map[string]map[string]string) {
		cityParts = append(cityParts, fmt.Sprintf("\"%s\"=>[\"code\"=>\"%s\",\"father\"=>\"%s\"]", name, details["code"], details["father"]))
	}
	cityStr := strings.Join(cityParts, ",\n        ")

	// 地区数据
	areaParts := []string{}
	for code, details := range data["area"].(map[string]map[string]string) {
		areaParts = append(areaParts, fmt.Sprintf("\"%s\"=>[\"name\"=>\"%s\",\"father\"=>\"%s\"]", code, details["name"], details["father"]))
	}
	areaStr := strings.Join(areaParts, ",\n        ")

	return fmt.Sprintf("$data = [\n    \"province\"=>[\n    %s,\n    ],\n    \"city\"=>[\n        %s,\n        ],\n    \"area\"=>[\n        %s,\n        ],\n    ];", provinceStr, cityStr, areaStr)
}

func copyToClipboard(data string) {
	fmt.Println("Data copied to clipboard:")
	fmt.Println(data)

	// 尝试复制到剪贴板
	err := clipboard.WriteAll(data)
	if err != nil {
		showMessage("失败", "无法写入到剪贴板")
	} else {
		showMessage("成功", "表格内容已写入剪贴板，请前往目的地粘贴")
	}
}

func showMessage(title, message string) {
	walk.MsgBox(nil, title, message, walk.MsgBoxIconInformation)
}

func getSheetByName(xlFile *xls.WorkBook, name string) *xls.WorkSheet {
	for i := 0; i < xlFile.NumSheets(); i++ {
		sheet := xlFile.GetSheet(i)
		if sheet.Name == name {
			return sheet
		}
	}
	return nil
}
