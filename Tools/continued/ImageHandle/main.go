package main

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"strconv"

	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"gocv.io/x/gocv"
)

var (
	imagePath       string
	targetWidth     int
	targetHeight    int
	backgroundColor color.RGBA = color.RGBA{255, 255, 255, 255}
	widthLineEdit   *walk.LineEdit
	heightLineEdit  *walk.LineEdit
	imageView       *walk.ImageView
	radioGroup      *walk.RadioButtonGroup
	sizeDropdown    *walk.ComboBox
	unitDropdown    *walk.ComboBox
	colorDialog     *walk.ColorDialog
)

func main() {
	var mw *walk.MainWindow

	MainWindow{
		AssignTo: &mw,
		Title:    "Image Processor",
		MinSize:  Size{Width: 600, Height: 400},
		Layout:   VBox{},
		Children: []Widget{
			PushButton{
				Text: "选择图片",
				OnClicked: func() {
					dlg := new(walk.FileDialog)
					dlg.Filter = "Image Files (*.jpg;*.jpeg;*.png)|*.jpg;*.jpeg;*.png"
					if ok, err := dlg.ShowOpen(mw); err != nil {
						log.Println("Error selecting file:", err)
						return
					} else if !ok {
						return
					}
					imagePath = dlg.FilePath

					img, err := walk.NewImageFromFile(imagePath)
					if err != nil {
						log.Println("Error loading image:", err)
						return
					}
					imageView.SetImage(img)
				},
			},
			RadioButtonGroupBox{
				Title: "选择尺寸类型",
				Buttons: []RadioButton{
					{
						Text: "证件照",
						OnClicked: func() {
							sizeDropdown.SetVisible(true)
							unitDropdown.SetVisible(false)
							widthLineEdit.SetVisible(false)
							heightLineEdit.SetVisible(false)
						},
					},
					{
						Text: "自定义尺寸",
						OnClicked: func() {
							sizeDropdown.SetVisible(false)
							unitDropdown.SetVisible(true)
							widthLineEdit.SetVisible(true)
							heightLineEdit.SetVisible(true)
						},
					},
				},
			},
			ComboBox{
				AssignTo: &sizeDropdown,
				Model:    []string{"一寸照 (25x35mm)", "二寸照 (35x49mm)"},
				Visible:  false,
			},
			ComboBox{
				AssignTo: &unitDropdown,
				Model:    []string{"像素", "厘米"},
				Visible:  false,
			},
			Composite{
				Layout: HBox{},
				Children: []Widget{
					Label{
						Text:    "宽度:",
						Visible: true,
					},
					LineEdit{
						AssignTo: &widthLineEdit,
						Visible:  true,
					},
					Label{
						Text:    "高度:",
						Visible: true,
					},
					LineEdit{
						AssignTo: &heightLineEdit,
						Visible:  true,
					},
				},
			},
			PushButton{
				Text: "选择背景颜色",
				OnClicked: func() {
					if colorDialog.Show(mw) == walk.DlgCmdOK {
						selectedColor := colorDialog.Color()
						backgroundColor = color.RGBA{R: uint8(selectedColor.R), G: uint8(selectedColor.G), B: uint8(selectedColor.B), A: 255}
					}
				},
			},
			PushButton{
				Text: "处理图片",
				OnClicked: func() {
					processImage()
				},
			},
			ImageView{
				AssignTo: &imageView,
			},
		},
	}.Run()
}

func processImage() {
	if imagePath == "" {
		log.Println("No image selected")
		return
	}

	img := gocv.IMRead(imagePath, gocv.IMReadColor)
	if img.Empty() {
		log.Fatalf("failed to read image from: %s", imagePath)
		return
	}
	defer img.Close()

	if sizeDropdown.Visible() {
		size := sizeDropdown.Text()
		if size == "一寸照 (25x35mm)" {
			targetWidth = 295
			targetHeight = 413
		} else if size == "二寸照 (35x49mm)" {
			targetWidth = 413
			targetHeight = 579
		}
	} else if unitDropdown.Visible() {
		unit := unitDropdown.Text()
		width, err := strconv.Atoi(widthLineEdit.Text())
		if err != nil {
			log.Println("Invalid width")
			return
		}
		height, err := strconv.Atoi(heightLineEdit.Text())
		if err != nil {
			log.Println("Invalid height")
			return
		}

		if unit == "厘米" {
			width = int(float64(width) * 37.7953) // 1cm = 37.7953 pixels
			height = int(float64(height) * 37.7953)
		}

		targetWidth = width
		targetHeight = height
	}

	newSize := gocv.NewMat()
	gocv.Resize(img, &newSize, image.Point{X: targetWidth, Y: targetHeight}, 0, 0, gocv.InterpolationLinear)
	defer newSize.Close()

	newImg := gocv.NewMatWithSize(targetHeight, targetWidth, gocv.MatTypeCV8UC3)
	defer newImg.Close()
	bgColor := gocv.NewMatWithSizeFromScalar(gocv.NewScalar(float64(backgroundColor.B), float64(backgroundColor.G), float64(backgroundColor.R), 0), targetHeight, targetWidth, gocv.MatTypeCV8UC3)
	defer bgColor.Close()

	mask := gocv.NewMatWithSize(targetHeight, targetWidth, gocv.MatTypeCV8UC1)
	defer mask.Close()
	gocv.InRangeWithScalar(newSize, gocv.NewScalar(0, 0, 0, 0), gocv.NewScalar(255, 255, 255, 0), &mask)

	gocv.CopyToWithMask(newSize, &newImg, mask)

	gocv.AddWeighted(newImg, 1, bgColor, 1, 0, &newImg)

	if ok := gocv.IMWrite("output.jpg", newImg); !ok {
		log.Fatalf("failed to save image to: %s", "output.jpg")
	}
	fmt.Println("Image processed and saved as output.jpg")
}
