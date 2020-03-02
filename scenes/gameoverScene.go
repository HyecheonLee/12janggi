package scenes

import (
	"12janggi/scenemanager"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"github.com/hajimehoshi/ebiten/inpututil"
	"log"
)

type GameOverScene struct {
	gameOverImg *ebiten.Image
}

func (g *GameOverScene) Startup() {
	var err error
	g.gameOverImg, _, err = ebitenutil.NewImageFromFile("./images/gameover.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
}

func (g *GameOverScene) Update(screen *ebiten.Image) error {
	screen.DrawImage(g.gameOverImg, nil)
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		// Set GameScene
		scenemanager.SetScene(&StartScene{})
	}
	return nil
}
