package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"fmt"
)

func main() {
	var _map mapgame
	var level []LevelInfo
	var level1,level2 LevelInfo 
	level1.Init(1,2,3,"maps/mazeLVL1.txt")
	level2.Init(4,5,6,"maps/mazeLVL2.txt")
	level = append(level,level1,level2)
	_map.Init(level)
	fmt.Println(_map.level[0].mazeMap)

	rl.InitWindow(800, 600, "rl Sprite Demo")
    rl.SetTargetFPS(60)

    sprite := createSprite("assets/wall.png", 1.0, true, 100, 100)

    for !rl.WindowShouldClose() {
        rl.BeginDrawing()

        rl.ClearBackground(rl.RayWhite)
        drawSprite(sprite)

        rl.EndDrawing()
    }

}

