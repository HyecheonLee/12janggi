package scenes

import (
	"12janggi/global"
	"12janggi/scenemanager"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"log"
)

type GimulType int
type TeamType int

const (
	TeamNone TeamType = iota
	TeamGreen
	TeamRed
)

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

type GameScene struct {
	bgimg       *ebiten.Image
	gimulImgs   [GimulMax]*ebiten.Image
	selectedImg *ebiten.Image
	board       [global.BoardWidth][global.BoardHeight]GimulType
	selected    bool
	selectedX   int
	selectedY   int
	currentTeam TeamType
	gameOver    bool
}

func (g *GameScene) Startup() {
	g.gameOver = false
	g.currentTeam = TeamGreen

	var err error
	g.bgimg, _, err = ebitenutil.NewImageFromFile("./images/bgimg.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulGreenWang], _, err = ebitenutil.NewImageFromFile("./images/green_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulGreenJa], _, err = ebitenutil.NewImageFromFile("./images/green_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulGreenJang], _, err = ebitenutil.NewImageFromFile("./images/green_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulGreenSang], _, err = ebitenutil.NewImageFromFile("./images/green_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulRedWang], _, err = ebitenutil.NewImageFromFile("./images/red_wang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulRedJa], _, err = ebitenutil.NewImageFromFile("./images/red_ja.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulRedJang], _, err = ebitenutil.NewImageFromFile("./images/red_jang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.gimulImgs[GimulRedSang], _, err = ebitenutil.NewImageFromFile("./images/red_sang.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	g.selectedImg, _, err = ebitenutil.NewImageFromFile("./images/selected.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {
			g.board[i][j] = GimulNone
		}
	}
	g.board[0][0] = GimulGreenSang
	g.board[0][1] = GimulGreenWang
	g.board[0][2] = GimulGreenJang
	g.board[1][1] = GimulGreenJa

	g.board[3][0] = GimulRedSang
	g.board[3][1] = GimulRedWang
	g.board[3][2] = GimulRedJang
	g.board[2][1] = GimulRedJa
}

func (g *GameScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(g.bgimg, nil)
	if g.gameOver {
		return nil
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		i, j := x/global.GridWidth, y/global.GridHeight
		if i >= 0 && i < global.GridWidth && j >= 0 && j < global.GridHeight {
			if g.selected {
				if i == g.selectedX && j == g.selectedY {
					g.selected = false
				} else {
					g.moveGimul(g.selectedX, g.selectedY, i, j)
				}
			} else {
				if g.board[i][j] != GimulNone && g.currentTeam == GetTeamType(g.board[i][j]) {
					g.selected = true
					g.selectedX, g.selectedY = i, j
				}
			}
		}
	}
	for i := 0; i < global.BoardWidth; i++ {
		for j := 0; j < global.BoardHeight; j++ {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*i), float64(global.GimulStartY+j*global.GridHeight))
			switch g.board[i][j] {
			case GimulGreenWang:
				screen.DrawImage(g.gimulImgs[GimulGreenWang], opts)
			case GimulGreenJa:
				screen.DrawImage(g.gimulImgs[GimulGreenJa], opts)
			case GimulGreenJang:
				screen.DrawImage(g.gimulImgs[GimulGreenJang], opts)
			case GimulGreenSang:
				screen.DrawImage(g.gimulImgs[GimulGreenSang], opts)
			case GimulRedWang:
				screen.DrawImage(g.gimulImgs[GimulRedWang], opts)
			case GimulRedJa:
				screen.DrawImage(g.gimulImgs[GimulRedJa], opts)
			case GimulRedJang:
				screen.DrawImage(g.gimulImgs[GimulRedJang], opts)
			case GimulRedSang:
				screen.DrawImage(g.gimulImgs[GimulRedSang], opts)
			}
		}
	}
	if g.selected {
		opts := &ebiten.DrawImageOptions{}
		opts.GeoM.Translate(float64(global.GimulStartX+global.GridWidth*g.selectedX), float64(global.GimulStartY+global.GridHeight*g.selectedY))
		screen.DrawImage(g.selectedImg, opts)
	}
	return nil
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
func (g *GameScene) moveGimul(prevX int, prevY int, tarX int, tarY int) {
	if g.isMovable(prevX, prevY, tarX, tarY) {
		g.OnDie(g.board[tarX][tarY])
		g.board[prevX][prevY], g.board[tarX][tarY] = GimulNone, g.board[prevX][prevY]
		g.selected = false
		if g.currentTeam == TeamGreen {
			g.currentTeam = TeamRed
		} else {
			g.currentTeam = TeamGreen
		}
	}
}
func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
func (g *GameScene) OnDie(gimulType GimulType) {
	if gimulType == GimulGreenWang || gimulType == GimulRedWang {
		scenemanager.SetScene(&GameOverScene{})
	}
}
func (g *GameScene) isMovable(prevX int, prevY int, tarX int, tarY int) bool {
	if tarX < 0 || tarY < 0 {
		return false
	}
	if tarX >= global.BoardWidth || tarY >= global.BoardHeight {
		return false
	}
	if GetTeamType(g.board[prevX][prevY]) == GetTeamType(g.board[tarX][tarY]) {
		return false
	}
	switch g.board[prevX][prevY] {
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
