package main

import (
	"image/color"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

type object struct {
	x, y   float64
	dx, dy float64
	color  color.Color
	radius float64
}

const (
	minSpeed       = 2 // 最小の速度値
	maxSpeed       = 5 // 最大の速度値
	objectRadius   = 5 // オブジェクトの半径
	speedVariation = 2
)

func newObject() *object {
	return &object{
		x:      rand.Float64() * screenWidth,
		y:      rand.Float64() * screenHeight,
		dx:     (minSpeed + rand.Float64()*(maxSpeed-minSpeed)) * math.Pow(-1, float64(rand.Intn(2))),
		dy:     (minSpeed + rand.Float64()*(maxSpeed-minSpeed)) * math.Pow(-1, float64(rand.Intn(2))),
		radius: objectRadius,
		color:  color.RGBA{R: uint8(rand.Intn(256)), G: uint8(rand.Intn(256)), B: uint8(rand.Intn(256)), A: 255},
	}
}

func (o *object) Move() {
	o.x += o.dx
	o.y += o.dy

	// X方向の跳ね返り
	if o.x-o.radius <= 0 {
		o.x = o.radius
		o.dx = func() float64 {
			dx := -o.dx * (1 + float64(getRandomSpeedVariation()))
			if dx > 0 {
				return -dx
			}
			return minSpeed + rand.Float64()*(maxSpeed-minSpeed)
		}()
	} else if o.x+o.radius >= screenWidth {
		o.x = screenWidth - o.radius
		o.dx = func() float64 {
			dx := -o.dx * (1 + float64(getRandomSpeedVariation()))
			if dx > 0 {
				return -dx
			}
			return minSpeed + rand.Float64()*(maxSpeed-minSpeed)
		}()
	}

	// Y方向の跳ね返り
	if o.y-o.radius <= 0 {
		o.y = o.radius
		o.dy = func() float64 {
			dy := -o.dy * (1 + float64(getRandomSpeedVariation()))
			if dy > 0 {
				return -dy
			}
			return minSpeed + rand.Float64()*(maxSpeed-minSpeed)
		}()
	} else if o.y+o.radius >= screenHeight {
		o.y = screenHeight - o.radius
		o.dy = func() float64 {
			dy := -o.dy * (1 + float64(getRandomSpeedVariation()))
			if dy > 0 {
				return -dy
			}
			return minSpeed + rand.Float64()*(maxSpeed-minSpeed)
		}()
	}
}

func getRandomSpeedVariation() float64 {
	var variation int
	for {
		variation = rand.Intn(2*speedVariation) - speedVariation
		if variation != 0 {
			break
		}
	}
	return float64(variation)
}

func (o *object) Draw(screen *ebiten.Image) {
	// 新しい透明なイメージを作成
	diameter := int(2 * o.radius)
	img := ebiten.NewImage(diameter, diameter)

	// 中心からの距離に基づいてピクセルを設定
	for y := -o.radius; y < o.radius; y++ {
		for x := -o.radius; x < o.radius; x++ {
			if x*x+y*y < o.radius*o.radius {
				img.Set(int(o.radius+x), int(o.radius+y), o.color)
			}
		}
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(o.x-o.radius, o.y-o.radius) // 中心を基準に移動
	screen.DrawImage(img, op)
}
