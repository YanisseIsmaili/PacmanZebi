package main

import (
	"github.com/hajimehoshi/ebiten"
)

type infogame struct {
	score int
	startgame bool
	gameover bool
	levelcomplete bool
	level int
	maze []string
	scoremax int
}

type caracter struct {
	speedpacman int
	speedenemy int
	nuberenemy int
	mazefile string
}

type Sprite struct {
	img *ebiten.Image
	faces map[byte]*ebiten.Image
	x float64
	y float64
	speed float64
	direction byte
	visibility bool
}

type screen struct {
	ScreensizeX int
	ScreensizeY int
}

type spriteMap struct {
	blocksize int
}

type mapgame struct {
	level []LevelInfo
}

type LevelInfo struct {
	pacmanSpeed int
	enemySpeed int
	enemyNumber int
	mazeMap []string
}

