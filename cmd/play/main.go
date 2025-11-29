package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/user/go-arcade/game"
)

func main() {
	g, err := game.NewGame()
	if err != nil {
		log.Fatalf("failed to create game: %v", err)
	}
	ebiten.SetWindowSize(800, 600)
	ebiten.SetWindowTitle("Go Arcade — Minimal Ebiten Game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatalf("game error: %v", err)
	}
}
