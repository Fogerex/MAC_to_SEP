package main

import (
	"regexp"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

func main() {

	a := app.New()
	win := a.NewWindow("MAC to SEP")
	wMAC := widget.NewEntry()
	wSEP := widget.NewEntry()
	//eCompanyNum.SetText("5120")
	//wPending := widget.NewLabel(lPending)
	//wActive := widget.NewLabel(lActive)
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

	win.SetContent(widget.NewVBox(
		wMAC,
		wSEP,
		btnConvert,
		btnCopy,
	))
	win.Resize(fyne.NewSize(250, 150))
	win.ShowAndRun()

}
