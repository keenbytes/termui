package main

import (
	"time"

	"github.com/keenbytes/termui"
)

type widgetTime struct {}

func (w *widgetTime) iterate(pane *termui.Pane) {
	now := time.Now()
	pane.Write(0, 0, now.Format("15:04:05"))
}
