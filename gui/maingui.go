package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

// Starting with a quick example from the fyne repository.
// Still don't know if I'm going to stick with fyne, but I guess I'll find out.
func StartGui() {
	appl := app.New()
	win := appl.NewWindow("ciphgoer")

	win.Resize(fyne.NewSize(400, 400))
	ab := widget.NewLabel("Selected file path will appear here.")

	filePicker := func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.NewError(err, win).Show()
				return
			}

			if f == nil {
				dialog.NewInformation("Error", "Action canceled by user", win).Show()
				return
			}

			ab.SetText(f.URI().String())
		}, win)
	}

	win.SetContent(container.NewVBox(
		ab,
		widget.NewButton("Select a file", func() {
			filePicker()
		}),
	))

	win.ShowAndRun()
}
