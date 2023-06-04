package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"

	"github.com/Jimba-lol/ebitengine-counter/constants"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, world")
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return constants.ResX, constants.ResY
}

func main() {
	ebiten.SetWindowSize(constants.ResX*2, constants.ResY*2)
	ebiten.SetWindowTitle("Counter")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
