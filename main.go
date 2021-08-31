package main

import (
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var GopherImageUrls = []string{
	"https://t4.ftcdn.net/jpg/02/12/91/13/360_F_212911335_0jEapN9npAyc3hkJITBd5UWal14qnnod.jpg",
	"https://upload.wikimedia.org/wikipedia/commons/c/cb/Pocket-Gopher_Ano-Nuevo-SP.jpg",
}

func loadNextGopherImage(currentImageUrlIndex int) (fyne.Resource, int) {
	newImageUrlIndex := currentImageUrlIndex + 1
	if newImageUrlIndex > len(GopherImageUrls) - 1 {
		newImageUrlIndex = 0
	}

	resource, err := fyne.LoadResourceFromURLString(GopherImageUrls[newImageUrlIndex])
	if err != nil {
		log.Print("failed to load resource")
		return nil, currentImageUrlIndex
	}

	return resource, newImageUrlIndex
}

func setUpHomePageContainer() *fyne.Container {
	// Define a welcome text centered
	text := canvas.NewText("Display a Gopher!", color.Black)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, currentImageUrlIndex = loadNextGopherImage(-1)
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0
	gopherImg.FillMode = canvas.ImageFillContain

	// Define a button to show next image
	switchImageBtn := widget.NewButton("Next", func() {
		resource, currentImageUrlIndex = loadNextGopherImage(currentImageUrlIndex)
		if resource != nil {
			gopherImg.Resource = resource
		}

		gopherImg.Refresh()
	})
	switchImageBtn.Importance = widget.HighImportance

	// Display a vertical box containing text, image and button
	homePageContainer := container.NewVBox(
		text,
		gopherImg,
		switchImageBtn,
	)

	return homePageContainer
}

func setUpShapesContainer() *fyne.Container {
	rectangle := canvas.NewRectangle(color.RGBA{R: 55, G: 74, B: 88, A:255})
	rectangle.SetMinSize(fyne.NewSize(300, 300))

	gridContainer := container.NewGridWithColumns(
		4,
	)

	circle := canvas.NewCircle(color.RGBA{R:231, G:51, B:137, A:255})
	circle.StrokeColor = color.White
	circle.StrokeWidth = 5

	line := canvas.NewLine(color.White)
	line.StrokeWidth = 5

	shapesContainer := container.New(
		layout.NewMaxLayout(),
		rectangle,
		circle,
		line,
		gridContainer,
	)

	col1 := canvas.NewText("Grid 1", color.Black)
	col1.Alignment = fyne.TextAlignCenter
	col2 := widget.NewButton("Max layout", func() {
		shapesContainer.Layout = layout.NewMaxLayout()
		shapesContainer.Refresh()
	})
	col3 := widget.NewButton("VBox layout", func() {
		shapesContainer.Layout = layout.NewVBoxLayout()
		shapesContainer.Refresh()
	})
	col4 := canvas.NewText("Grid 4", color.Black)
	col4.Alignment = fyne.TextAlignCenter
	col5 := canvas.NewText("Grid 5", color.Black)
	col5.Alignment = fyne.TextAlignCenter
	col6 := widget.NewButton("Center layout", func() {
		shapesContainer.Layout = layout.NewCenterLayout()
		shapesContainer.Refresh()
	})
	col7 := widget.NewButton("Padded layout", func() {
		shapesContainer.Layout = layout.NewPaddedLayout()
		shapesContainer.Refresh()
	})
	col8 := canvas.NewText("Grid 8", color.Black)
	col8.Alignment = fyne.TextAlignCenter
	cols := []fyne.CanvasObject{col1, col2, col3, col4, col5, col6, col7, col8}

	for _, col := range cols {
		gridContainer.Add(col)
	}

	return shapesContainer
}

func setUpMainMenu(app fyne.App, window fyne.Window) {
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { app.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Ross Vincent"),
			), window)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	window.SetMainMenu(mainMenu)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gopher")

	setUpMainMenu(myApp, myWindow)

	homePageContainer := setUpHomePageContainer()

	shapesContainer := setUpShapesContainer()

	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), homePageContainer),
		container.NewTabItem("Shapes", shapesContainer),
	)
	tabs.SetTabLocation(container.TabLocationTop)

	myWindow.SetContent(tabs)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	myWindow.ShowAndRun()
}