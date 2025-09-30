package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("Started Fyne...")

	a := app.New()
	w := a.NewWindow("Resize window")

	w.Resize(fyne.NewSize(600, 400))

	w.ShowAndRun()
}