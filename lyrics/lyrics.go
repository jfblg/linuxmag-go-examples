package main

import (
	"sort"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func main() {
	songdir := "data"
	lyrics, err := songFinder(songdir)
	if err != nil {
		panic(err)
	}

	if err := ui.Init(); err != nil {
		panic(err)
	}
	defer ui.Close()

	// Listbox displaying songs
	lb := widgets.NewList()
	items := []string{}
	for k := range lyrics {
		items = append(items, k)
	}
	sort.Strings(items)
	lb.Title = "Pick a song"
	w, h := ui.TerminalDimensions()
	lb.SetRect(0, 0, w, h)
	lb.Rows = items
	lb.SelectedRow = 0
	lb.SelectedRowStyle = ui.NewStyle(ui.ColorGreen)

	// Listbox displaying lyrics
	ltext := widgets.NewList()
	ltextLines := []string{}
	ltext.Rows = ltextLines
	ltext.SetRect(0, 0, w, h)
	ltext.Title = "Text"
	ltext.TextStyle.Fg = ui.ColorGreen
	ltext.SelectedRowStyle = ui.NewStyle(ui.ColorRed)

	handleUI(lb, ltext, lyrics)
}
