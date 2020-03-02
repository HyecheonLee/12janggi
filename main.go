package main

import (
	"12janggi/scenemanager"
	"12janggi/scenes"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"log"
)

// Type aliasing
type GimulType int

const (
	GimulNone GimulType = -1 + iota
	GimulGreenWang
	GimulGreenJa
	GimulGreenJang
	GimulGreenSang
	GimulRedWang
	GimulRedJa
	GimulRedJang
	GimulRedSang
	GimulMax
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

type TeamType int

const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

var (
	board       [BoardWidth][BoardHeight]GimulType
	bgimg       *ebiten.Image
	selectedImg *ebiten.Image
	gimulImgs   [GimulMax]*ebiten.Image
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType = TeamGreen
	gameOver    bool     = false
)

func main() {

	var err error
	err = loadImages(err)
	boardInit()

	scenemanager.SetScene(&scenes.StartScene{})

	err = ebiten.Run(scenemanager.Update, ScreenWidth, ScreenHeight, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
func boardInit() {
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			board[i][j] = GimulNone
		}
	}
	board[0][0] = GimulGreenSang
	board[0][1] = GimulGreenWang
	board[0][2] = GimulGreenJang
	board[1][1] = GimulGreenJa

	board[3][0] = GimulRedSang
	board[3][1] = GimulRedWang
	board[3][2] = GimulRedJang
	board[2][1] = GimulRedJa
}

func loadImages(err error) error {
	bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	gimulImgs[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}

	return err
}
func GetTeamType(gimulType GimulType) TeamType {
	if gimulType == GimulGreenJa ||
		gimulType == GimulGreenJang ||
		gimulType == GimulGreenSang ||
		gimulType == GimulGreenWang {
		return TeamGreen
	}
	if gimulType == GimulRedJa ||
		gimulType == GimulRedJang ||
		gimulType == GimulRedSang ||
		gimulType == GimulRedWang {
		return TeamRed
	}
	return TeamNone
}
func update(screen *ebiten.Image) error {
	screen.DrawImage(bgimg, nil)
	if gameOver {
		return nil
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/GridWidth, y/GridHeight
		if i >= 0 && i < GridWidth && j >= 0 && j < GridHeight {
			if selected {
				if i == selectedX && j == selectedY {
					selected = false
				} else {
					moveGimul(selectedX, selectedY, i, j)
				}
			} else {
				if board[i][j] != GimulNone && currentTeam == GetTeamType(board[i][j]) {
					selected = true
					selectedX, selectedY = i, j
				}
			}
		}
	}
	for i := 0; i < BoardWidth; i++ {
		for j := 0; j < BoardHeight; j++ {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(GimulStartX+GridWidth*i), float64(GimulStartY+j*GridHeight))
			switch board[i][j] {
			case GimulGreenWang:
				screen.DrawImage(gimulImgs[GimulGreenWang], opts)
				//Draw GreenWang
			case GimulGreenJa:
				screen.DrawImage(gimulImgs[GimulGreenJa], opts)
				//Draw GreenJa
			case GimulGreenJang:
				screen.DrawImage(gimulImgs[GimulGreenJang], opts)
				//Draw GreenJang
			case GimulGreenSang:
				screen.DrawImage(gimulImgs[GimulGreenSang], opts)
				//Draw GreenSang
			case GimulRedWang:
				screen.DrawImage(gimulImgs[GimulRedWang], opts)
				//Draw RedWang
			case GimulRedJa:
				screen.DrawImage(gimulImgs[GimulRedJa], opts)
				//Draw RedJa
			case GimulRedJang:
				screen.DrawImage(gimulImgs[GimulRedJang], opts)
				//Draw RedJang
			case GimulRedSang:
				screen.DrawImage(gimulImgs[GimulRedSang], opts)
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

func moveGimul(prevX int, prevY int, tarX int, tarY int) {
	if isMovable(prevX, prevY, tarX, tarY) {
		OnDie(board[tarX][tarY])
		board[prevX][prevY], board[tarX][tarY] = GimulNone, board[prevX][prevY]
		selected = false
		if currentTeam == TeamGreen {
			currentTeam = TeamRed
		} else {
			currentTeam = TeamGreen
		}
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func OnDie(gimulType GimulType) {
	if gimulType == GimulGreenWang || gimulType == GimulRedWang {
		gameOver = true
	}
}
func isMovable(prevX int, prevY int, tarX int, tarY int) bool {
	if tarX < 0 || tarY < 0 {
		return false
	}
	if tarX >= BoardWidth || tarY >= BoardHeight {
		return false
	}
	if GetTeamType(board[prevX][prevY]) == GetTeamType(board[tarX][tarY]) {
		return false
	}
	switch board[prevX][prevY] {
	case GimulGreenJa:
		return prevX+1 == tarX && prevY == tarY
	case GimulRedJa:
		return prevX-1 == tarX && prevY == tarY
	case GimulGreenSang, GimulRedSang:
		return (abs(prevX-tarX) == 1) && (abs(prevY-tarY) == 1)
	case GimulGreenJang, GimulRedJang:
		return abs(prevX-tarX)+abs(prevY-tarY) == 1
	case GimulGreenWang, GimulRedWang:
		return abs(prevX-tarX) <= 1 || abs(prevY-tarY) <= 1
	}
	return false
}
