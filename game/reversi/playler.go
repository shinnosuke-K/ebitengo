package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const (
	defaultCursorX = 3
	defaultCursorY = 3
)

type ManualPlayer struct {
	CursorX int
	CursorY int

	piece        Piece
	name         string
	cursorColor  color.Color
	defaultColor color.Color
	count        int
}

func newManualPlayer(name string, piece Piece, color color.Color) *ManualPlayer {
	return &ManualPlayer{
		CursorX: defaultCursorX,
		CursorY: defaultCursorY,

		piece:        piece,
		name:         name,
		cursorColor:  color,
		defaultColor: color,
		count:        2,
	}
}

func (p *ManualPlayer) MoveCursor() {
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowUp) {
		if p.CursorY-1 >= 0 {
			p.CursorY--
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowDown) {
		if p.CursorY+1 <= 7 {
			p.CursorY++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowRight) {
		if p.CursorX+1 <= 7 {
			p.CursorX++
		}
	}
	if inpututil.IsKeyJustPressed(ebiten.KeyArrowLeft) {
		if p.CursorX-1 >= 0 {
			p.CursorX--
		}
	}
}

// SetCount はプレイヤーの石の数を設定する
func (p *ManualPlayer) SetCount(count int) { p.count = count }

func (p *ManualPlayer) SetColor(color color.Color) { p.cursorColor = color }

func (p *ManualPlayer) ResetColor() { p.cursorColor = p.defaultColor }

func (p *ManualPlayer) ResetCursor() {
	p.CursorX = defaultCursorX
	p.CursorY = defaultCursorY
}

func (p *ManualPlayer) DrawCursor(screen *ebiten.Image) {
	vector.StrokeLine(screen, float32(240+p.CursorX*40), float32(80+p.CursorY*40), float32(280+p.CursorX*40), float32(80+p.CursorY*40), 4, p.cursorColor, true)
	vector.StrokeLine(screen, float32(240+p.CursorX*40), float32(80+p.CursorY*40), float32(240+p.CursorX*40), float32(120+p.CursorY*40), 4, p.cursorColor, true)
	vector.StrokeLine(screen, float32(280+p.CursorX*40), float32(80+p.CursorY*40), float32(280+p.CursorX*40), float32(120+p.CursorY*40), 4, p.cursorColor, true)
	vector.StrokeLine(screen, float32(240+p.CursorX*40), float32(120+p.CursorY*40), float32(280+p.CursorX*40), float32(120+p.CursorY*40), 4, p.cursorColor, true)
}

// DrawDisplayName はプレイヤーの名前を描画する
func (p *ManualPlayer) DrawDisplayName(screen *ebiten.Image) {
	var x, y int
	if p.piece == PieceBlack {
		x = 100
		y = 200
	} else {
		x = 650
		y = 200
	}
	text.Draw(screen, p.name, basicfont.Face7x13, x, y, color.White)
}

// DrawPieceCount はプレイヤーの石の数を描画する
func (p *ManualPlayer) DrawPieceCount(screen *ebiten.Image) {
	var x, y int
	if p.piece == PieceBlack {
		x = 100
		y = 220
	} else {
		x = 650
		y = 220
	}
	text.Draw(screen, fmt.Sprintf("Count: %d", p.count), basicfont.Face7x13, x, y, color.White)
}

// DrawWinner は勝者を描画する
func (p *ManualPlayer) DrawWinner(screen *ebiten.Image) {
	var x, y int
	if p.piece == PieceBlack {
		x = 100
		y = 100
	} else {
		x = 650
		y = 100
	}
	text.Draw(screen, "Win!", basicfont.Face7x13, x, y, colornames.Red)
}

// DrawLooser は敗者を描画する
func (p *ManualPlayer) DrawLooser(screen *ebiten.Image) {
	var x, y int
	if p.piece == PieceBlack {
		x = 100
		y = 100
	} else {
		x = 650
		y = 100
	}
	text.Draw(screen, "Lose...", basicfont.Face7x13, x, y, colornames.Blue)
}
