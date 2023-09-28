package main

import (
	"bufio"
	"os"
)

func readMazeFile(fileName string) []string {

	maze := []string{}

	file, err := os.Open(fileName)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		maze = append(maze, line)
	}
	return maze
}
