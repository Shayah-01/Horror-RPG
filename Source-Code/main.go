package main

import (
	"image"
	"os"

	_ "image/png"

	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func main() { // Directly firing the game from func main is not done, Run func is the conventional func main here
	pixelgl.Run(run)
}

func run() {

	// configuring the details of how our window would start
	cfg := pixelgl.WindowConfig{
		Title:     "HRPG by Shayah-01",
		Bounds:    pixel.R(0, 0, 360, 480),
		Resizable: true,
		VSync:     true,
	}
	// Opening the windows with prev. conf. - error if failed
	win, err := pixelgl.NewWindow(cfg)
	win.SetSmooth(true)
	// Will print error & help me know what went wrong if window fails to open
	if err != nil {
		panic(err)
	}

	// Calling the image from the helperFunction
	pic, err := loadPicture("sprites/pxArt.png")
	if err != nil {
		panic(err)
	}
	sprite := pixel.NewSprite(pic, pic.Bounds())
	spritePosition := win.Bounds().Center()

	// Defining what will happen as long as the window stays open
	// Keep windows updating for keypresses/events until as long as its NOT closed!
	for !win.Closed() && !win.JustPressed(pixelgl.KeyEscape) {
		win.Clear(colornames.Rosybrown)

		// Only pixel.IM = "no spacial transformations" = static image
		// .Scaled = resize & .Moved = position
		sprite.Draw(win, pixel.IM.Scaled(pixel.ZV, 0.1).Moved(spritePosition))

		// Configuring movements of characters START
		if win.Pressed(pixelgl.KeyRight) {
			spritePosition.X += 2.0
		}
		if win.Pressed(pixelgl.KeyUp) {
			spritePosition.Y += 2.0
		}
		if win.Pressed(pixelgl.KeyLeft) {
			spritePosition.X -= 2.0
		}
		if win.Pressed(pixelgl.KeyDown) {
			spritePosition.Y -= 2.0
		}
		// Configuring movements of characters END

		win.Update()
	}
}

func loadPicture(path string) (pixel.Picture, error) { // helper func to load image
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		return nil, err
	}
	return pixel.PictureDataFromImage(img), nil
}
