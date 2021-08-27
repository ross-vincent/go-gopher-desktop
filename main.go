package main

import (
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

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Gopher")

	// Main menu
	fileMenu := fyne.NewMenu("File",
		fyne.NewMenuItem("Quit", func() { myApp.Quit() }),
	)

	helpMenu := fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			dialog.ShowCustom("About", "Close", container.NewVBox(
				widget.NewLabel("Welcome to Gopher, a simple Desktop app created in Go with Fyne."),
				widget.NewLabel("Version: v0.1"),
				widget.NewLabel("Author: Ross Vincent"),
			), myWindow)
		}))
	mainMenu := fyne.NewMainMenu(
		fileMenu,
		helpMenu,
	)
	myWindow.SetMainMenu(mainMenu)

	// Define a welcome text centered
	text := canvas.NewText("Display a Gopher!", color.Black)
	text.Alignment = fyne.TextAlignCenter

	// Define a Gopher image
	var resource, currentImageUrlIndex = loadNextGopherImage(-1)
	gopherImg := canvas.NewImageFromResource(resource)
	gopherImg.SetMinSize(fyne.Size{Width: 500, Height: 500}) // by default size is 0, 0

	switchImageBtn := widget.NewButton("Next", func() {
		resource, currentImageUrlIndex = loadNextGopherImage(currentImageUrlIndex)
		if resource != nil {
			gopherImg.Resource = resource
		}

		gopherImg.Refresh()
	})
	switchImageBtn.Importance = widget.HighImportance

	// Display a vertical box containing text, image and button
	box := container.NewVBox(
		text,
		gopherImg,
		switchImageBtn,
	)

	// Display our content
	myWindow.SetContent(box)

	// Close the App when Escape key is pressed
	myWindow.Canvas().SetOnTypedKey(func(keyEvent *fyne.KeyEvent) {

		if keyEvent.Name == fyne.KeyEscape {
			myApp.Quit()
		}
	})

	// Show window and run app
	myWindow.ShowAndRun()
}