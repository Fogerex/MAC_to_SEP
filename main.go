package main

import (
	"regexp"
	"flag"
	"fmt"


	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	version     = "0.1.0" // версия приложения
	showVersion bool      // флаг, указывающий на необходимость вывода версии
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "вывести версию приложения и выйти")
	flag.Parse()
}

func main() {
	// версия приложения
	if showVersion {
		fmt.Println(version)
		return
	}

	a := app.New()
	win := a.NewWindow(fmt.Sprintf("MAC to SEP %s",version))
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
	win.Resize(fyne.NewSize(280, 150))
	win.ShowAndRun()

}
