package main

import (
	"12janggi/global"
	"12janggi/scenemanager"
	"12janggi/scenes"
	"github.com/hajimehoshi/ebiten"
	"log"
)

func main() {
	scenemanager.SetScene(&scenes.StartScene{})

	err := ebiten.Run(scenemanager.Update, global.ScreenWidth, global.ScreenHeight, 1.0, "12 Janggi")
	if err != nil {
		log.Fatalf("Ebiten run error: %v", err)
	}
}
