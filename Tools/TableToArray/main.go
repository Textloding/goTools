package main

import (
	"fmt"
	"github.com/extrame/xls"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"github.com/tealeg/xlsx"
)

func main() {
	mw, err := createMainWindow()
	if err != nil {
		fmt.Println("Failed to create main window:", err)
		return
	}

	mw.Run()

}

func createMainWindow() (*walk.MainWindow, error) {
	var mw *walk.MainWindow
	var inColumn *walk.LineEdit

	// 加载嵌入的资源图标
	icon, err := walk.NewIconFromResourceId(2) // IDI_ICON1 默认为资源ID 1
	if err != nil {
		// 处理错误
		panic(err)
	}

	if err := (MainWindow{
		AssignTo: &mw,
		Title:    "选择Excel文件",
		Icon:     icon, // 设置图标
		Size:     Size{Width: 400, Height: 200},
		Layout:   VBox{},
		Children: []Widget{

			Label{
				TextColor: walk.RGB(218, 8, 9),
				Text:      "注意:需先定义输入列然后选择文件，否则结果为空",
				Font: Font{
					Family:    "微软雅黑",
					PointSize: 12, // 设置文字大小为12
				},
			},
			Label{Text: "预定义输入列 (如 A, B, ... ; * 表示所有列;不选即为空):"},
			LineEdit{AssignTo: &inColumn},
			PushButton{
				Text: "选择文件",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Filter = "Excel files (*.xlsx;*.xls)|*.xlsx;*.xls"
					if ok, err := dlg.ShowOpen(mw); err != nil {
						walk.MsgBox(mw, "错误", "无法打开文件对话框: "+err.Error(), walk.MsgBoxIconError)
					} else if ok {
						processFile(dlg.FilePath, inColumn.Text())
					}
				},
			},
		},
	}).Create(); err != nil {
		return nil, err
	}

	return mw, nil
}

func processFile(filePath, columns string) {
	fmt.Printf("Processing file: %s\n", filePath)
	if strings.HasSuffix(filePath, ".xlsx") {
		processXLSX(filePath, columns)
	} else if strings.HasSuffix(filePath, ".xls") {
		processXLS(filePath, columns)
	}
}

func processXLSX(filename, columns string) {
	xlFile, err := xlsx.OpenFile(filename)
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}

	columnIndexes := parseColumns(columns, xlFile.Sheets[0].MaxCol)
	var data []string
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			for _, colIndex := range columnIndexes {
				if colIndex < len(row.Cells) {
					text := row.Cells[colIndex].String()
					data = append(data, fmt.Sprintf("'%s'", text))
				}
			}
		}
	}
	copyToClipboard(data)
}

func processXLS(filename, columns string) {
	xlFile, err := xls.Open(filename, "utf-8")
	if err != nil {
		fmt.Printf("Failed to open file: %v\n", err)
		return
	}

	sheet := xlFile.GetSheet(0)
	if sheet == nil {
		fmt.Println("No sheet found in file")
		return
	}

	maxCols := findMaxCols(sheet)
	columnIndexes := parseColumns(columns, maxCols)

	var data []string
	for i := 0; i <= int(sheet.MaxRow); i++ {
		row := sheet.Row(i)
		if row != nil {
			for _, colIndex := range columnIndexes {
				if colIndex < row.LastCol() {
					text := row.Col(colIndex)
					data = append(data, fmt.Sprintf("'%s'", text))
				}
			}
		}
	}
	copyToClipboard(data)
}

func findMaxCols(sheet *xls.WorkSheet) int {
	maxCols := 0
	for i := 0; i <= int(sheet.MaxRow); i++ {
		row := sheet.Row(i)
		if row.LastCol() > maxCols {
			maxCols = row.LastCol()
		}
	}
	return maxCols
}

func copyToClipboard(data []string) {
	arrayString := fmt.Sprintf("[\n%s\n]", strings.Join(data, ",\n"))
	fmt.Println("Data copied to clipboard:")
	fmt.Println(arrayString)

	// 清理字符串，移除NUL字符
	cleanArrayString := cleanString(arrayString)

	// 尝试复制到剪贴板
	err := clipboard.WriteAll(cleanArrayString)
	if err != nil {
		showMessage("失败", "无法写入到剪贴板")
	}
	showMessage("成功", "表格内容已写入剪贴板请前往目的地粘贴")
}

func parseColumns(cols string, maxCols int) []int {
	if cols == "*" {
		var allCols []int
		for i := 0; i < maxCols; i++ {
			allCols = append(allCols, i)
		}
		return allCols
	}

	var indexes []int
	for _, col := range strings.Split(cols, ",") {
		if idx, err := columnToIndex(strings.TrimSpace(col)); err == nil {
			if idx < maxCols {
				indexes = append(indexes, idx)
			}
		}
	}
	return indexes
}

func columnToIndex(col string) (int, error) {
	if len(col) != 1 {
		return 0, fmt.Errorf("invalid column: %s", col)
	}
	return int(col[0] - 'A'), nil
}

func cleanString(str string) string {
	return strings.ReplaceAll(str, "\x00", "")
}
func showMessage(title, message string) {
	walk.MsgBox(nil, title, message, walk.MsgBoxIconInformation)
}
