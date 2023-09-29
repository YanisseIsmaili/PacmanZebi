package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	screenWidth  = 480
	screenHeight = 480
	blockSize    = 40
)

type Game struct {
	// Define the game map. Use 'W' for wall, ' ' for empty space, 'P' for Pac-Man, and 'F' for food.
	gameMap [][]rune

	// Define Pac-Man's initial position and direction.
	pacmanStartX    int
	pacmanStartY    int
	pacmanDirection string

	// Define the sprite images.
	pacmanImage  *ebiten.Image
	foodImage    *ebiten.Image
	ghostImage   *ebiten.Image
	wallImage    *ebiten.Image
	gameOverImage *ebiten.Image

	// Define game variables.
	pacmanX, pacmanY float64
	pacmanSpeed      float64
	pacmanSize       float64
	pacmanAngle      float64
	pacmanScore      int

	foodX, foodY float64
	foodSize     float64

	ghostX, ghostY float64
	ghostSpeed     float64
	ghostSize      float64
	ghostDirection string

	gameOver        bool
	gameOverMessage string
}

func NewGame() *Game {
	game := &Game{
		gameMap: [][]rune{
			{'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
			{'W', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W'},
			{'W', ' ', 'W', 'W', 'W', ' ', 'W', 'W', 'W', ' ', 'W', 'W', 'W', ' ', 'W', 'W', 'W', ' ', 'W', 'W'},
			{'W', ' ', 'W', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W', 'W'},
			{'W', ' ', 'W', 'W', 'W', 'W', 'W', 'W', ' ', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', ' ', 'W', 'W'},
			{'W', ' ', ' ', ' ', ' ', ' ', ' ', 'W', ' ', 'W', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W'},
			{'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', ' ', 'W', 'W', 'W', ' ', 'W', 'W', 'W', ' ', ' ', ' ', 'W'},
			{'W', ' ', ' ', ' ', ' ', ' ', ' ', 'W', ' ', 'W', ' ', 'W', ' ', 'W', ' ', ' ', ' ', ' ', 'W', 'W'},
			{'W', ' ', 'W', 'W', 'W', ' ', 'W', 'W', ' ', 'W', ' ', 'W', ' ', 'W', 'W', 'W', ' ', ' ', 'W', 'W'},
			{'W', ' ', 'W', ' ', ' ', ' ', 'W', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', 'W', ' ', 'W', ' ', 'W'},
			{'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W', 'W'},
		},

		pacmanStartX:    1,
		pacmanStartY:    1,
		pacmanDirection: right,

		pacmanSpeed:  4,
		pacmanSize:   blockSize,
		pacmanAngle:  0,
		pacmanScore:  0,
		foodSize:     blockSize,
		ghostSpeed:   2,
		ghostSize:    blockSize,
		ghostDirection: up,

		gameOver:        false,
		gameOverMessage: "Game Over! Press Space to Restart.",
	}

	return game
}

func (g *Game) Update(screen *ebiten.Image) error {
	if ebiten.IsDrawingSkipped() {
		return nil
	}

	// Handle user input.
	if !g.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyUp) {
			g.pacmanDirection = up
		}
		if ebiten.IsKeyPressed(ebiten.KeyDown) {
			g.pacmanDirection = down
		}
		if ebiten.IsKeyPressed(ebiten.KeyLeft) {
			g.pacmanDirection = left
		}
		if ebiten.IsKeyPressed(ebiten.KeyRight) {
			g.pacmanDirection = right
		}
	}

	// Update Pac-Man's position.
	if !g.gameOver {
		switch g.pacmanDirection {
		case up:
			g.pacmanY -= g.pacmanSpeed
			g.pacmanAngle = 3 * math.Pi / 2
		case down:
			g.pacmanY += g.pacmanSpeed
			g.pacmanAngle = math.Pi / 2
		case left:
			g.pacmanX -= g.pacmanSpeed
			g.pacmanAngle = math.Pi
		case right:
			g.pacmanX += g.pacmanSpeed
			g.pacmanAngle = 0
		}
	}

	// Check for collisions with walls.
	g.checkWallCollision()

	// Check for collisions with food.
	if g.isFoodCollision() {
		g.eatFood()
	}

	// Update ghost's position.
	if !g.gameOver {
		g.moveGhost()
	}

	// Check for collisions with the ghost.
	if g.isGhostCollision() {
		g.gameOver = true
	}

	// Draw the game elements.
	g.drawMap(screen)
	g.drawPacman(screen)
	g.drawFood(screen)
	g.drawGhost(screen)

	// Draw the game over message if the game is over.
	if g.gameOver {
		gameOverOpts := &ebiten.DrawImageOptions{}
		gameOverOpts.GeoM.Translate(120, screenHeight/2)
		screen.DrawImage(g.gameOverImage, gameOverOpts)

		ebitenutil.DebugPrintAt(screen, g.gameOverMessage, 140, screenHeight/2+100)
	}

	return nil
}

func (g *Game) checkWallCollision() {
	// Calculate the cell position that Pac-Man occupies.
	pacmanCellX := int(g.pacmanX / g.pacmanSize)
	pacmanCellY := int(g.pacmanY / g.pacmanSize)

	// Check if Pac-Man hits a wall.
	if pacmanCellX >= 0 && pacmanCellX < len(g.gameMap[0]) && pacmanCellY >= 0 && pacmanCellY < len(g.gameMap) {
		if g.gameMap[pacmanCellY][pacmanCellX] == 'W' {
			g.pacmanX, g.pacmanY = float64(pacmanCellX)*g.pacmanSize, float64(pacmanCellY)*g.pacmanSize
		}
	}
}

func (g *Game) isFoodCollision() bool {
	foodCellX := int(g.foodX / g.foodSize)
	foodCellY := int(g.foodY / g.foodSize)

	pacmanCellX := int(g.pacmanX / g.pacmanSize)
	pacmanCellY := int(g.pacmanY / g.pacmanSize)

	return foodCellX == pacmanCellX && foodCellY == pacmanCellY
}

func (g *Game) eatFood() {
	g.foodX, g.foodY = g.generateRandomPosition()
	g.pacmanScore++
}

func (g *Game) moveGhost() {
	// Calculate the cell position that the ghost occupies.
	ghostCellX := int(g.ghostX / g.ghostSize)
	ghostCellY := int(g.ghostY / g.ghostSize)

	// Calculate the target cell for the ghost.
	targetCellX, targetCellY := int(g.pacmanX/g.pacmanSize), int(g.pacmanY/g.pacmanSize)

	// Calculate the possible directions for the ghost.
	directions := []string{}
	if ghostCellX > 0 && g.gameMap[ghostCellY][ghostCellX-1] != 'W' {
		directions = append(directions, left)
	}
	if ghostCellX < len(g.gameMap[0])-1 && g.gameMap[ghostCellY][ghostCellX+1] != 'W' {
		directions = append(directions, right)
	}
	if ghostCellY > 0 && g.gameMap[ghostCellY-1][ghostCellX] != 'W' {
		directions = append(directions, up)
	}
	if ghostCellY < len(g.gameMap)-1 && g.gameMap[ghostCellY+1][ghostCellX] != 'W' {
		directions = append(directions, down)
	}

	// Calculate the best direction for the ghost based on the target cell.
	bestDirection := ""
	minDistance := math.MaxFloat64
	for _, dir := range directions {
		dx, dy := 0, 0
		switch dir {
		case left:
			dx = -1
		case right:
			dx = 1
		case up:
			dy = -1
		case down:
			dy = 1
		}

		distance := math.Sqrt(math.Pow(float64(targetCellX-ghostCellX+dx), 2) + math.Pow(float64(targetCellY-ghostCellY+dy), 2))
		if distance < minDistance {
			minDistance = distance
			bestDirection = dir
		}
	}

	// Move the ghost.
	switch bestDirection {
	case left:
		g.ghostX -= g.ghostSpeed
	case right:
		g.ghostX += g.ghostSpeed
	case up:
		g.ghostY -= g.ghostSpeed
	case down:
		g.ghostY += g.ghostSpeed
	}
}

func (g *Game) isGhostCollision() bool {
	ghostCellX := int(g.ghostX / g.ghostSize)
	ghostCellY := int(g.ghostY / g.ghostSize)

	pacmanCellX := int(g.pacmanX / g.pacmanSize)
	pacmanCellY := int(g.pacmanY / g.pacmanSize)

	return ghostCellX == pacmanCellX && ghostCellY == pacmanCellY
}

func (g *Game) drawMap(screen *ebiten.Image) {
	for y, row := range g.gameMap {
		for x, cell := range row {
			switch cell {
			case 'W':
				g.drawImage(screen, g.wallImage, float64(x)*g.pacmanSize, float64(y)*g.pacmanSize)
			case ' ':
				// Empty space
			}
		}
	}
}

func (g *Game) drawPacman(screen *ebiten.Image) {
	pacmanOpts := &ebiten.DrawImageOptions{}
	pacmanOpts.GeoM.Translate(g.pacmanX, g.pacmanY)
	pacmanOpts.GeoM.Rotate(g.pacmanAngle)
	pacmanOpts.GeoM.Scale(1.5, 1.5) // Scale Pac-Man for better visibility
	screen.DrawImage(g.pacmanImage, pacmanOpts)
}

func (g *Game) drawFood(screen *ebiten.Image) {
	g.drawImage(screen, g.foodImage, g.foodX, g.foodY)
}

func (g *Game) drawGhost(screen *ebiten.Image) {
	g.drawImage(screen, g.ghostImage, g.ghostX, g.ghostY)
}

func (g *Game) drawImage(screen *ebiten.Image, img *ebiten.Image, x, y float64) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(x, y)
	screen.DrawImage(img, opts)
}

func (g *Game) generateRandomPosition() (x, y float64) {
	rand.Seed(time.Now().UnixNano())
	x = float64(rand.Intn(screenWidth / int(g.foodSize)))
	y = float64(rand.Intn(screenHeight / int(g.foodSize)))
	return x * g.foodSize, y * g.foodSize
}
func main() {
	// Initialize Ebiten and the game.
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pac-Man")

	game := NewGame()

	// Run the game loop.
	if err := ebiten.Run(game.Update, screenWidth, screenHeight, 2, "Pac-Man"); err != nil {
		log.Fatal(err)
	}
}