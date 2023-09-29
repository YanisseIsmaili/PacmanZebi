package main

import (
    "math"
    rl "github.com/gen2brain/raylib-go/raylib"
)

const blockSize = 15



func getMazePointFromPosition(x float32, y float32) (int, int) {
    col := int(math.Round(float64(x / blockSize)))
    row := int(math.Round(float64(y / blockSize)))
    return col, row
}

func getPositionFromMazePoint(col int, row int) (float32, float32) {
    return blockSize * float32(col), blockSize * float32(row)
}

func (sprite *Sprite) Init(imgPath string, size float32, visibility bool, x float32, y float32) {
    sprite.texture = loadImage(imgPath)
    sprite.x = x
    sprite.y = y
   // sprite.size = size
    sprite.visibility = visibility
}

func locateGameObjects(gameInfo GameInfo) GameObjects {
    var objects GameObjects
    var wall, PACMAN Sprite
    var pacman Pacman

    for row, line := range gameInfo.maze {
        objects.mazeSprite = append(objects.mazeSprite, make([]Sprite, len(line)))

        for col, char := range line {
            x, y := getPositionFromMazePoint(col, row)

            switch char {
            case '0':
                wall.Init("assets/wall.png", blockSize, true, x, y)
                objects.mazeWall = append(objects.mazeWall, wall)
                objects.mazeSprite[row][col] = wall

            case 'P':
                PACMAN.Init("assets/pacman.png", blockSize, true, x, y)
                UP_SPRITE := createSprite("assets/pacmanU.png", blockSize, true, x, y)
                RIGHT_SPRITE := createSprite("assets/pacmanR.png", blockSize, true, x, y)
                DOWN_SPRITE := createSprite("assets/pacmanD.png", blockSize, true, x, y)
                LEFT_SPRITE := createSprite("assets/pacmanL.png", blockSize, true, x, y)
                IDLE_SPRITE := createSprite("assets/pacmanI.png", blockSize, true, x, y)

                pacman.faces = map[byte]rl.Texture2D{
                    'U': UP_SPRITE.texture,
                    'R': RIGHT_SPRITE.texture,
                    'D': DOWN_SPRITE.texture,
                    'L': LEFT_SPRITE.texture,
                    'I': IDLE_SPRITE.texture,
                }
                objects.pacman = pacman

            }
        }
    }
    return objects
}
