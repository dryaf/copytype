package main

import (
	"fmt"
	"strconv"
	"time"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	"github.com/dryaf/copytype/typer"
)

func main() {
	st, err := typer.NewStringTyper(typer.KeysOSXus)
	if err != nil {
		panic(err)
	}

	app := app.New()
	var waitAndTypeBtn *widget.Button

	w := app.NewWindow("Copy Type")
	timerLbl := widget.NewLabel("typing in # seconds")
	textToType := widget.NewMultiLineEntry()

	secondsTxt := widget.NewEntry()
	secondsTxt.SetText("5")

	waitAndTypeBtn = widget.NewButton("Wait and type", func() {
		waitAndTypeBtn.Disable()
		secondsTxt.SetReadOnly(true)
		seconds, err := strconv.ParseInt(secondsTxt.Text, 10, 64)
		if err != nil {
			secondsTxt.SetText("this needs to be a number")
			return
		}
		for ; seconds != -1; seconds-- {
			time.Sleep(1 * time.Second)
			secondsTxt.SetText(fmt.Sprint(seconds))
		}
		err = st.TypeString(textToType.Text)
		waitAndTypeBtn.Enable()
		secondsTxt.SetReadOnly(false)
		if err != nil {
			secondsTxt.SetText(err.Error())
		}
	})

	w.SetContent(widget.NewVBox(
		widget.NewLabel("Paste text you want to be typed:"),
		textToType,
		timerLbl,
		secondsTxt,
		waitAndTypeBtn,
		widget.NewButton("Quit", func() {
			app.Quit()
		}),
	))

	w.ShowAndRun()
}
