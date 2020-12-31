package main

import (
	"log"

	_ "image/png"

	"github.com/egnimos/pacman/pacman"
	"github.com/hajimehoshi/ebiten"
)

func main() {
	g := pacman.NewGame()

	if err := ebiten.Run(g.Update, g.ScreenWidth(), g.ScreenHeight(), 1, "Pacman"); err != nil {
		log.Fatal(err)
	}
}
