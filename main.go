package main

import (
	"regexp"

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
	btn := widget.NewButton("Convert", func() {

		var re = regexp.MustCompile(`[[:punct:]]`)
		str := wMAC.Text
		str45 := re.ReplaceAllString(str, "")
		wSEP.SetText("SEP" + str45)

	})

	win.SetContent(widget.NewVBox(
		wMAC,
		wSEP,
		btn,
	))
	win.ShowAndRun()

}
