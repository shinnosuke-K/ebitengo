package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

type Board [][]Piece

func newBoard() Board {
	return Board{
		{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceWhite, PieceBlack, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceBlack, PieceWhite, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
		{PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone, PieceNone},
	}
}

// DrawPlate は盤面を描画する（色は茶色）
func (b Board) DrawPlate(screen *ebiten.Image) {
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(float64(230), float64(70))
	plate := ebiten.NewImage(340, 340)
	plate.Fill(color.RGBA{R: 116, G: 80, B: 48})
	screen.DrawImage(plate, opts)
}

func (b Board) DrawBoard(screen *ebiten.Image) {
	var (
		offsetX = 240
		offsetY = 80
	)
	for y, row := range b {
		for x, cell := range row {
			opts := &ebiten.DrawImageOptions{}
			opts.GeoM.Translate(float64(offsetX+x*40), float64(offsetY+y*40))
			screen.DrawImage(cell.Draw(), opts)
		}
	}
}

func (b Board) Set(x, y int, piece Piece) { b[y][x] = piece }
