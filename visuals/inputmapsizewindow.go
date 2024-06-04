package visuals

import (
	"errors"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var errInputIsEmpty = errors.New("Input is empty")
var errInputIsNotPosInt = errors.New("Input is not a positive integer")

func CreateMapSizeInputWindow(app fyne.App) {
	w := app.NewWindow("Input Map Size")

	// Map Size Form
	labelCol := widget.NewLabel("Column")
	entryCol := widget.NewEntry()
	entryCol.SetPlaceHolder("Input column size...")

	labelRow := widget.NewLabel("Row")
	entryRow := widget.NewEntry()
	entryRow.SetPlaceHolder("Input row size...")

	form := container.New(layout.NewFormLayout(), labelCol, entryCol, labelRow, entryRow)

	// Button
	button := widget.NewButton("Create Map", func() {
		colSize, colErr := validateMapSizeInput(entryCol.Text)
		if colErr != nil {
			dialog.NewInformation("Error", colErr.Error(), w).Show()
			return
		}

		rowSize, rowErr := validateMapSizeInput(entryRow.Text)
		if rowErr != nil {
			dialog.NewInformation("Error", rowErr.Error(), w).Show()
			return
		}

		CreateVisualizerWindow(app, colSize, rowSize)

		w.Close()
	})

	content := container.NewVBox(form, button)

	w.SetContent(content)
	w.Resize(fyne.NewSize(640, 480))
	w.Show()
}

func validateMapSizeInput(input string) (int, error) {
	if input == "" {
		return -1, errInputIsEmpty
	}

	i, err := strconv.Atoi(input)
	if err != nil {
		return -1, err
	}

	if i <= 0 {
		return -1, errInputIsNotPosInt
	}

	return i, nil
}
