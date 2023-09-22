package main

import (
	"regexp"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {

	a := app.New()
	win := a.NewWindow("MAC to SEP")
	wMAC := widget.NewEntry()
	wSEP := widget.NewEntry()
	btnConvert := widget.NewButton("Convert", func() {

		var re = regexp.MustCompile(`[[:punct:]]`)
		str := wMAC.Text
		str45 := re.ReplaceAllString(str, "")
		wSEP.SetText("SEP" + str45)

	})
	btnCopy := widget.NewButton("Копировать", func() {
		// Записываем содержимое в буфер обмена
		win.Clipboard().SetContent(wSEP.Text)
	})


	vbox := container.NewVBox(
		wMAC,
		wSEP,
		btnConvert,
		btnCopy,
	)
	
	win.SetContent(
		vbox,
	)
	win.Resize(fyne.NewSize(250, 150))
	win.ShowAndRun()

}
