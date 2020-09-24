package main

import (
	"bufio"
	"strings"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func handleUI(lb *widgets.List, ltext *widgets.List, lyrics map[string]Lyrics) {
	ui.Render(lb)
	inFocus := lb

	uiEvents := ui.PollEvents()
	var scanner *bufio.Scanner

	for {
		select {
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "j", "<Down>":
				if inFocus == lb {
					lb.ScrollDown()
					ui.Render(lb)
				}
			case "k", "<Up>":
				if inFocus == lb {
					lb.ScrollUp()
					ui.Render(lb)
				}
			case "<Enter>":
				if inFocus == lb {
					sel := lb.Rows[lb.SelectedRow]
					ltext.Title = sel
					inFocus = ltext
					text := lyrics[sel].Text
					scanner = bufio.NewScanner(strings.NewReader(text))
					ui.Render(ltext)
				}
				if inFocus == ltext {
					morelines := false
					for scanner.Scan() {
						line := scanner.Text()
						if line == "" {
							continue
						}
						ltext.Rows = append(ltext.Rows, line)
						morelines = true
						ltext.ScrollDown()
						ui.Render(ltext)
						break
					}
					if !morelines {
						inFocus = lb
						ltext.Rows = ltext.Rows[:0]
						ui.Render(lb)
					}
				}
			case "<Escape>":
				inFocus = lb
				ltext.Rows = ltext.Rows[:0]
				ui.Render(lb)
			}

		}
	}
}
