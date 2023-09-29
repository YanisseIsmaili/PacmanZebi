func getMazePointFromPosition(x float64, y float64) (int, int) {

    col := int(math.Round(x/float64(blockSize)))
    row := int(math.Round(y/float64(blockSize)))
    return col, row
}