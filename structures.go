package main

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
	speedpacman float64
	speedenemy float64
	nuberenemy int
	mazefile string
}