package gui

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// Starting with a quick example from the fyne repository.
// Still don't know if I'm going to stick with fyne, but I guess I'll find out.
func StartGui() {
	appl := app.New()
	win := appl.NewWindow("Test")

	hello := widget.NewLabel("Oi, teste")
	win.SetContent(container.NewVBox(
		hello,
		widget.NewButton("Hi!", func() {
			hello.SetText("Bem venido!")
		}),
	))

	win.ShowAndRun()
}
