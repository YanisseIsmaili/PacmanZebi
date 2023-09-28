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

type sprite struct {
	x float64
	y float64
	speed float64
	direction byte
}