package main

import (
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	cursorX float64
	cursorY float64
	radius  float64
	color   color.Color
}

func newPlayer() *Player {
	return &Player{
		cursorX: 240,
		cursorY: 240,
		radius:  10,
		color:   color.White,
	}
}

func (p *Player) MoveCursor() {
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) {
		if p.cursorY-1 >= 0 {
			p.cursorY -= maxSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) {
		if p.cursorY+1 <= screenHeight {
			p.cursorY += maxSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) {
		if p.cursorX+1 <= screenWidth {
			p.cursorX += maxSpeed
		}
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) {
		if p.cursorX-1 >= 0 {
			p.cursorX -= maxSpeed
		}
	}
}

// Intersects　は指定されたObjectとの当たり判定を行う
func (p *Player) Intersects(obj *object) bool {
	distance := math.Pow(p.cursorX-obj.x, 2) + math.Pow(p.cursorY-obj.y, 2)
	return distance <= p.radius*p.radius+obj.radius*obj.radius
}

// Draw はプレイヤーのカーソル（円形）を描画する
func (p *Player) Draw(screen *ebiten.Image) {
	// 新しい透明なイメージを作成
	diameter := int(2 * p.radius)
	img := ebiten.NewImage(diameter, diameter)

	// 中心からの距離に基づいてピクセルを設定
	for y := -p.radius; y < p.radius; y++ {
		for x := -p.radius; x < p.radius; x++ {
			if x*x+y*y < p.radius*p.radius {
				img.Set(int(p.radius+x), int(p.radius+y), p.color)
			}
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.cursorX-p.radius, p.cursorY-p.radius) // 中心を基準に移動
	screen.DrawImage(img, op)
}
