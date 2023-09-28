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
	img *ebiten.Image
	faces map[byte]*ebiten.Image
	x float64
	y float64
	speed float64
	direction byte
<<<<<<< HEAD
}
=======
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
	pacmanSpeed float64
	enemySpeed float64
	enemyNumber int
	mazeFile string
}

func (m *mapgame)Init(level []LevelInfo){
}

func (l *LevelInfo)Init(pacmanSpeed int, enemySpeed int, enemyNumber int, mazeFile string){
	l.pacmanSpeed = pacmanSpeed
	l.enemySpeed = enemySpeed
	l.enemyNumber = enemyNumber
	l.mazeFile = mazeFile
}
>>>>>>> 6d64d59e0a6ec0a57eb9a5b4c045bf3e3d8d3bb0
