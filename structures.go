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
	speedpacman int
	speedenemy int
	nuberenemy int
	mazefile string
}

/*type sprite struct {
	img *ebiten.Image
	faces map[byte]*ebiten.Image
	x float64
	y float64
	speed float64
	direction byte
}
