package main

import (
	"context"
	"fmt"
	"github.com/mum4k/termdash"
	"github.com/mum4k/termdash/cell"
	"github.com/mum4k/termdash/container"
	"github.com/mum4k/termdash/linestyle"
	"github.com/mum4k/termdash/terminal/termbox"
	"github.com/mum4k/termdash/terminal/terminalapi"
	"github.com/mum4k/termdash/widgets/text"
)

func main() {
	commands := [][]string{
		{"/usr/bin/uptime"},
		{"/bin/bash", "-c", "dmesg | tail -10"}, 
		{"/usr/bin/vmstat", "1", "1"}, 
		{"/usr/bin/pidstat", "1", "1"},
	}

	t, err := termbox.New()
	if err != nil {
		panic(err)
	}
	defer t.Close()

	ctx, cancel := context.WithCancel(context.Background())

	widgets := []container.Option{
		container.ID("top"),
		container.Border(linestyle.Light),
		container.BorderTitle(" Greggalizer "),
	}

	panes := []*text.Text{}

	for _, command := range commands {
		pane, err := text.New(
			text.RollContent(),
			text.WrapAtWords(),
		)
		if err != nil {
			panic(err)
		}

		red := text.WriteCellOpts(
			cell.FgColor(cell.ColorRed),
		)
		pane.Write(
			fmt.Sprintf("%v\n", command), red)
		)
		pane.Write(panes, pane)

		rows := paneSplit(panes)

		widgets = append(widgets, rows)

		c, err := 

		
	
	}


}