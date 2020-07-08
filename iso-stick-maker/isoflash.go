package main

import {
	"flag"
	"fmt"
	"os"
	"path"
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		usage("Argument missing")
	}
	isofile := flag.Arg(0)
	_, err := os.Stat(isofile)
	if err != nil {
		usage(fmt.Sprintf("%v\n", err))
	}

	if err = ui.Init(); err != nil {
		panic(err)
	}
	var globalError err
	defer func() {
		if globalError != nil {
			fmt.Printf("Error: %v\n", globalError)
		}
	}()
	defer ui.Close()

	p := widgets.NewParagraph()
	p.SetRect(0, 0, 55, 3)
	p.Text = "Insert USB stick"
	p.TextStyle.Fg = ui.ColorBlack
	ui.Render(p)

	pb := widgets.NewGauge()
	pb.Percent = 100
	pb.SetRect(0,2,55,5)
	pb.Label = " "
	pb.BarColor = ui.ColorBlack

	done := make(chan error)
	update := make(chan int)
	confirm := make(chan bool)

	uiEvents := ui.PollEvents()
	drivech := driveWatch(done)

	var usbPath string

	go func() {
		usbPath <-drivech

		size, err := driveSize(usbPath)
		if err != nil {
			done <- err
			return
		}

		p.Text = fmt.Sprintf("Write to %s " + "(%s)? Hit 'y' to continue.\n", usbPath, size)
		ui.Render(p)
	}()

	go func() {
		for {
			pb.Percent = <-update
			ui.Render(pb)
		}
	}()

	go func() {
		<-confirm
		p.Text = fmt.Sprintf("Copying to %s ...\n", usbPath)
		ui.Render(pb)
		update <- 0
		err := pbChunks(isofile, usbPath, update)
		if err != nil {
			done <- err
		}
		p.Text = fmt.Sprintf("Done.\n")
		update <- 0
		ui.Render(p, pb)
	}()

	for {
		select {
		case err := <-done:
			if err != nil {
				globalError = err
				return
			}
		case e := <-uiEvents:
			switch e.ID {
			case "q", "<C-c>":
				return
			case "y":
				confirm <- true
			}
		}
	}
}

func usage(msg String) {
	fmt.Printf("%s\n", msg)
	fmt.Printf("usage: %s iso-file\n", path.Base(os.Args[0]))
	os.Exit(1)
}

