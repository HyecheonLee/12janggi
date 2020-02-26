package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"log"
)

// Type aliasing
type GimulType int

const (
	GimulTypeNone GimulType = -1 + iota
	GimulTypeGreenWang
	GimulTypeGreenJa
	GimulTypeGreenJang
	GimulTypeGreenSang
	GimulTypeRedWang
	GimulTypeRedJa
	GimulTypeRedJang
	GimulTypeRedSang
	GimulTypeMax
)
const (
	GimulStartX  = 20
	GridWidth    = 116
	GimulStartY  = 23
	GridHeight   = 116
	BoardWidth   = 4
	BoardHeight  = 3
	ScreenWidth  = 480
	ScreenHeight = 362
)

var (
	board       [BoardWidth][BoardHeight]GimulType
	bgimg       *ebiten.Image
	selectedImg *ebiten.Image
	gimulImgs   [GimulTypeMax]*ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
)

func main() {
	var err error
	err = loadImages(err)
	boardInit()
	err = ebiten.Run(update, ScreenWidth, ScreenHeight, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
func boardInit() {
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulTypeNone
		}
	}
	board[0][0] = GimulTypeGreenSang
	board[0][1] = GimulTypeGreenWang
	board[0][2] = GimulTypeGreenJang
	board[1][1] = GimulTypeGreenJa

	board[3][0] = GimulTypeRedSang
	board[3][1] = GimulTypeRedWang
	board[3][2] = GimulTypeRedJang
	board[2][1] = GimulTypeRedJa
}

func loadImages(err error) error {
	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulTypeRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	return err
}

func update(screen *ebiten.Image) error {
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridWidth, y/GridHeight
		if selected {
			if i == selectedX && j == selectedY {
				selected = false
			} else {
				move(selectedX, selectedY, i, j)
			}
		} else {
			if board[i][j] != GimulTypeNone {
				selected = true
				selectedX, selectedY = i, j
			}
		}
	}
	screen.DrawImage(bgimg, nil)
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+j*GridHeight))
			switch board[i][j] {
			case GimulTypeGreenWang:
				screen.DrawImage(gimulImgs[GimulTypeGreenWang], opts)
				//Draw GreenWang
			case GimulTypeGreenJa:
				screen.DrawImage(gimulImgs[GimulTypeGreenJa], opts)
				//Draw GreenJa
			case GimulTypeGreenJang:
				screen.DrawImage(gimulImgs[GimulTypeGreenJang], opts)
				//Draw GreenJang
			case GimulTypeGreenSang:
				screen.DrawImage(gimulImgs[GimulTypeGreenSang], opts)
				//Draw GreenSang
			case GimulTypeRedWang:
				screen.DrawImage(gimulImgs[GimulTypeRedWang], opts)
				//Draw RedWang
			case GimulTypeRedJa:
				screen.DrawImage(gimulImgs[GimulTypeRedJa], opts)
				//Draw RedJa
			case GimulTypeRedJang:
				screen.DrawImage(gimulImgs[GimulTypeRedJang], opts)
				//Draw RedJang
			case GimulTypeRedSang:
				screen.DrawImage(gimulImgs[GimulTypeRedSang], opts)
				//Draw RedJa
			}
		}
	}
	if selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(GimulStartX+GridWidth*selectedX), float64(GimulStartY+GridHeight*selectedY))
		screen.DrawImage(selectedImg, opts)
	}
	return nil
}

func move(prevX int, prevY int, tarX int, tarY int) {
	if isMovable(prevX, prevY, tarX, tarY) {
		board[prevX][prevY], board[tarX][tarY] = GimulTypeNone, board[prevX][prevY]
		selected = false
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func isMovable(prevX int, prevY int, tarX int, tarY int) bool {
	if tarX < 0 || tarY < 0 {
		return false
	}
	if tarX >= BoardWidth || tarY >= BoardHeight {
		return false
	}
	switch board[prevX][prevY] {
	case GimulTypeGreenJa:
		return prevX+1 == tarX && prevY == tarY
	case GimulTypeRedJa:
		return prevX-1 == tarX && prevY == tarY
	case GimulTypeGreenSang, GimulTypeRedSang:
		return (abs(prevX-tarX) == 1) && (abs(prevY-tarY) == 1)
	case GimulTypeGreenJang, GimulTypeRedJang:
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulTypeGreenWang, GimulTypeRedWang:
		return (abs(prevX-tarX) == 1) || (abs(prevY-tarY) == 1)
	}
	return false
}
