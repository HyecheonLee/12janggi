package scenes

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type StartScene struct {
}

func (s *StartScene) Update(screen *ebiten.Image) error {
	ebitenutil.DebugPrint(screen, "Start Scene")
	return nil
}
