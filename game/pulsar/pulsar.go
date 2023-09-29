package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2/text"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font/basicfont"
)

type Pulsar struct {
	objects       []*object
	elapsedFrames int
	elapsedTime   int
	player        *Player
	gameOver      bool
}

const (
	screenWidth  = 800
	screenHeight = 480
	numObjects   = 10
	gameTime     = 60 // 5 seconds
)

func newPulsar() *Pulsar {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	return &Pulsar{
		objects: func() []*object {
			objects := make([]*object, 0, numObjects)
			for i := 0; i < numObjects; i++ {
				objects = append(objects, newObject())
			}
			return objects
		}(),
		elapsedFrames: 0,
		elapsedTime:   0,
		player:        newPlayer(),
		gameOver:      false,
	}
}

func (p *Pulsar) Title() string { return "Pulsar" }

func (p *Pulsar) WindowSize() (int, int) { return screenWidth, screenHeight }

func (p *Pulsar) Layout(_, _ int) (int, int) { return screenWidth, screenHeight }

func (p *Pulsar) Update() error {
	if p.gameOver {
		if ebiten.IsKeyPressed(ebiten.KeyEnter) {
			p.Restart()
		}
		return nil
	}

	p.player.MoveCursor()
	for _, obj := range p.objects {
		obj.Move()
	}

	// オブジェクトの生成
	p.elapsedFrames++
	if p.elapsedFrames%60 == 0 {
		p.objects = append(p.objects, newObject(), newObject())
		p.elapsedTime++
	}

	// 衝突判定
	for _, obj := range p.objects {
		if p.player.Intersects(obj) {
			p.gameOver = true
		}
	}
	return nil
}

func (p *Pulsar) Draw(screen *ebiten.Image) {
	if p.gameOver {
		p.DrawText(screen, fmt.Sprintf("Score %d", p.elapsedFrames), 320, 170)
		p.DrawText(screen, "Press Enter to restart", 230, 200)
		return
	}

	p.player.Draw(screen)
	for _, obj := range p.objects {
		obj.Draw(screen)
	}
}

func (p *Pulsar) DrawText(screen *ebiten.Image, str string, x, y int) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Scale(2, 2)
	opt.GeoM.Translate(float64(x), float64(y))
	text.DrawWithOptions(screen, str, basicfont.Face7x13, opt)
}

func (p *Pulsar) Restart() {
	p.gameOver = false
	p.player = newPlayer()
	p.objects = nil
	p.elapsedFrames = 0
	p.elapsedTime = 0
	for i := 0; i < numObjects; i++ {
		p.objects = append(p.objects, newObject())
	}
}
