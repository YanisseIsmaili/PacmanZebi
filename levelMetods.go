package main

func (l *LevelInfo) Init(pacmanSpeed int, enemySpeed int, enemyNumber int, mazeFile string){
	l.pacmanSpeed = pacmanSpeed
	l.enemySpeed = enemySpeed
	l.enemyNumber = enemyNumber
	l.mazeMap = readMazeFile(mazeFile)
}
