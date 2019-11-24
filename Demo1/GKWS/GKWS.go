package main

import (
	"fmt"
	"github.com/visualfc/goqt/ui"
)

func main() {
	ui.Run(func() {
		widget := ui.NewWidget()
		widget.Show()
		btn := ui.NewPushButton()
		hbox := ui.NewHBoxLayout()
		hbox.AddWidget(btn)
		widget.SetLayout(hbox)
		btn.OnClicked(
			func() {
				fmt.Println("AAAA")
			})
	})
}