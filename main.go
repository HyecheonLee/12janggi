package main

import (
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	"log"
)

var (
	bgimg *ebiten.Image
)

func main() {
	var err error
	bgimg, _, err = ebitenutil.NewImageFromFile("./12janggi.png", ebiten.FilterDefault)
	if err != nil {
		log.Fatalf("read file error: %v", err)
	}
	err = ebiten.Run(func(screen *ebiten.Image) error {
		screen.DrawImage(bgimg, nil)
		return nil
	}, 500, 400, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}

}
