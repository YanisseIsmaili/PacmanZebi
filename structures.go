package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"	
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
texture    rl.Texture2D
x          float32
y          float32
speed      float32
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

type GameObjects struct{
	mazeWall    []Sprite
    mazeSprite  [][]Sprite
    pacman      Pacman
}

type Ghost struct{
	sprite []Sprite
	aLive bool
}

type Pacman struct {
	sprite []Sprite
	aLive bool
    faces map[byte]rl.Texture2D
}

type GameInfo struct {
	maze [][]byte
	score int
	isStarted bool
	isOver bool
	isLevelComplete bool
	level int
	maxScore int
	blockSize int
}

