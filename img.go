package main

import (
	"log"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"

)

func createSprite(imgFile string, width int, height int, x float64, y float64) Sprite {
    img, _ := ebiten.NewImage(width, height, ebiten.FilterDefault)
    imgFromFile, _, err := ebitenutil.NewImageFromFile(imgFile, ebiten.FilterDefault)
    originalWidth, originalHeight := imgFromFile.Size()
    scaleX := float64(width)/float64(originalWidth)
    scaleY := float64(height)/float64(originalHeight)
    opts := &ebiten.DrawImageOptions{}
    opts.GeoM.Scale(scaleX, scaleY)

    img.DrawImage(imgFromFile, opts)

	if err != nil {
		log.Fatal(err)
	}


	return Sprite{
	    img: img,
	    visibility: true,
	    x: x,
	    y: y,
	    speed: 1,
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	// Write your game's rendering.
}

func (g *ma) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}