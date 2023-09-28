package main

import (
    "os"
    "log"
    "bufio"
)

func (m *mapgame)Init(level []LevelInfo){
	for _, level := range level {
		m.level = append(m.level, level)
	}
}


func readMazeFile(fileName string) []string {

    maze := []string{}
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()
    	maze = append(maze, line)
    }

    return maze
}
