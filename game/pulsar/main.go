package main

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	reversi := newPulsar()
	ebiten.SetWindowSize(reversi.WindowSize())
	ebiten.SetWindowTitle(reversi.Title())
	if err := ebiten.RunGame(reversi); err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}
}
