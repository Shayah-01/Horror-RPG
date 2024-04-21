package main

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() { // Directly firing the game from func main is not done, Run func is the conventional func main here
	pixelgl.Run(run)
}

func run() {

	cfg := pixelgl.WindowConfig{ // configuring the details of how our window would start
		Title:     "HRPG by Shayah-01",
		Bounds:    pixel.R(0, 0, 250, 370),
		Resizable: true,
		VSync:     true,
	}
	win, err := pixelgl.NewWindow(cfg) // Opening the windows with prev. conf. - error if failed

	if err != nil { //
		panic(err) // Will print error & help me know what went wrong if window fails to open
	}

	// !win.Closed() =  !win.JustPressed(pixelgl.KeyEscape) = closing with key esc or mouse close button

	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) { // Keep windows updating for keypresses/events until as long as its NOT closed!
		win.Clear(colornames.Rosybrown) // Window will display CLEAR darkgrey space
		win.Update()
	}
}
