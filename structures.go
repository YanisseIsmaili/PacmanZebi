package main

import (
	"fmt"
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

fmt.Println("Hello World")