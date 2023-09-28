package main

import (
	"fmt"
)

func main() {
	var _map mapgame
	var level []LevelInfo
	var level1,level2 LevelInfo 
	level1.Init(1,2,3,"maze1.txt")
	level2.Init(4,5,6,"maze2.txt")
	level = append(level,level1,level2)
	_map.Init(level)
	fmt.Println(_map.level)
}