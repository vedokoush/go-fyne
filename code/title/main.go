package main

import (
	"fmt"
	"fyne.io/fyne/v2/app"
)

func main() {
	fmt.Println("Test Fyne...")

	a := app.New()

	w := a.NewWindow("Hello World!")

	w.ShowAndRun()
}