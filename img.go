package main

import (
	"fmt"
	"log"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func createSprite(imgPath string, speed float32, visibility bool, x float32, y float32) *Sprite {
    sprite := &Sprite{}
    sprite.Init(imgPath, speed, visibility, x, y)
    return sprite
}


func drawSprite(sprite *Sprite) {
    if sprite.visibility {
        rl.DrawTexture(sprite.texture, int32(sprite.x), int32(sprite.y), rl.RayWhite)
    }
}

func loadImage(imgPath string) rl.Texture2D {
    if img := rl.LoadTexture(imgPath); img.Width > 0 && img.Height > 0 {
        fmt.Println("nike baba")
        return img
    } else {
        log.Fatalf("Failed to load image: %s", imgPath)
        return rl.Texture2D{}
    }
}
