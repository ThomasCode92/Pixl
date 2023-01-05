package ui

import (
	"image/color"
	"pixl/swatch"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
)

func BuildSwatches(app *AppInit) *fyne.Container {
	canvasSwatches := make([]fyne.CanvasObject, 0, 64)
	for i := 0; i < cap(app.Swatches); i++ {
		initialColor := color.NRGBA{255, 255, 255, 255}
		swatch := swatch.NewSwatch(app.State, initialColor, i, func(s *swatch.Swatch) {
			for j := 0; j < len(app.Swatches); j++ {
				app.Swatches[j].Selected = false
				canvasSwatches[j].Refresh()
			}
			app.State.SwatchSelected = s.SwatchIndex
			app.State.BrushColor = s.Color
		})
		if i == 0 {
			swatch.Selected = true
			app.State.SwatchSelected = 0
			swatch.Refresh()
		}
		app.Swatches = append(app.Swatches, swatch)
		canvasSwatches = append(canvasSwatches, swatch)
	}

	return container.NewGridWrap(fyne.NewSize(20, 20), canvasSwatches...)
}
