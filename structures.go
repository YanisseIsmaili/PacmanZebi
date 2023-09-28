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
}*/

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
	mazeFile string
}

func (m *mapgame)Init(level []LevelInfo){
	for _, level := range level {
		m.level = append(m.level, level)
	}
}

func (l *LevelInfo)Init(pacmanSpeed int, enemySpeed int, enemyNumber int, mazeFile string){
	l.pacmanSpeed = pacmanSpeed
	l.enemySpeed = enemySpeed
	l.enemyNumber = enemyNumber
	l.mazeFile = mazeFile
}
