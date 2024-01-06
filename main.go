package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 300
	screenHeight = 300
)

type Game struct {
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw() error {
	return nil
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return screenWidth, screenHeight
}

func main() {
	fmt.Println("Hello, World!")
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Hello, World!")
}
