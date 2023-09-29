package main

func getPositionFromMazePoint(col int, row int) (float64, float64) {
    // here we are casting to float as we standard position usage in ebiten is float
    return float64(blockSize*col), float64(blockSize*row)
}