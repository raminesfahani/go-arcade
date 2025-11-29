package game

import (
	"bytes"
	"image/color"
	"log"
	"math"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/user/go-arcade/audio"
	"github.com/user/go-arcade/internal"
	"golang.org/x/image/font/basicfont"
)

const (
	screenWidth  = 800
	screenHeight = 600
)

// NewGame constructs and initializes the game state.
func NewGame() (*Game, error) {
	rand.Seed(time.Now().UnixNano())
	g := &Game{
		playerX: screenWidth / 2,
		playerY: screenHeight - 64,
		playerW: 28,
		playerH: 18,
		bgHue:   200,
	}
	// Pre-create simple images for sprites
	g.playerImage = ebiten.NewImage(32, 16)
	g.playerImage.Fill(color.RGBA{0x6f, 0xd1, 0xff, 0xff})
	g.enemyImage = ebiten.NewImage(28, 16)
	g.enemyImage.Fill(color.RGBA{0xff, 0x6f, 0x6f, 0xff})
	g.bulletImage = ebiten.NewImage(6, 10)
	g.bulletImage.Fill(color.RGBA{0xff, 0xff, 0x6f, 0xff})

	// minimal audio context
	if err := audio.Init(); err != nil {
		log.Printf("audio init warning: %v", err)
	}
	return g, nil
}

// Game is the main Ebiten game state.
type Game struct {
	playerX, playerY float64
	playerW, playerH float64
	playerImage      *ebiten.Image
	enemyImage       *ebiten.Image
	bulletImage      *ebiten.Image
	bullets          []Bullet
	enemies          []Enemy
	particles        []Particle
	score            int
	frames           int
	bgHue            float64
}

// Bullet is a player projectile.
type Bullet struct {
	x, y float64
	vy   float64
}

// Enemy is a simple falling enemy.
type Enemy struct {
	x, y   float64
	vx, vy float64
}

// Particle used for simple visual polish.
type Particle struct {
	x, y   float64
	vx, vy float64
	life   int
	col    color.RGBA
}

// Update advances the game state.
func (g *Game) Update() error {
	g.frames++
	// Input
	if ebiten.IsKeyPressed(ebiten.KeyArrowLeft) || ebiten.IsKeyPressed(ebiten.KeyA) {
		g.playerX -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowRight) || ebiten.IsKeyPressed(ebiten.KeyD) {
		g.playerX += 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowUp) || ebiten.IsKeyPressed(ebiten.KeyW) {
		g.playerY -= 4
	}
	if ebiten.IsKeyPressed(ebiten.KeyArrowDown) || ebiten.IsKeyPressed(ebiten.KeyS) {
		g.playerY += 4
	}
	if ebiten.IsKeyPressed(ebiten.KeySpace) {
		// fire
		if g.frames%8 == 0 {
			g.bullets = append(g.bullets, Bullet{g.playerX, g.playerY - 12, -8})
			audio.PlayShot()
		}
	}
	// Keep player on screen
	if g.playerX < 16 {
		g.playerX = 16
	}
	if g.playerX > screenWidth-16 {
		g.playerX = screenWidth - 16
	}
	if g.playerY < 16 {
		g.playerY = 16
	}
	if g.playerY > screenHeight-16 {
		g.playerY = screenHeight - 16
	}

	// Spawn enemies occasionally
	if g.frames%45 == 0 {
		x := float64(20 + rand.Intn(screenWidth-40))
		vx := (rand.Float64() - 0.5) * 1.5
		vy := 1 + rand.Float64()*1.5
		g.enemies = append(g.enemies, Enemy{x, -20, vx, vy})
	}

	// Update bullets
	for i := 0; i < len(g.bullets); i++ {
		b := &g.bullets[i]
		b.y += b.vy
	}
	// Remove off-screen bullets
	nb := g.bullets[:0]
	for _, b := range g.bullets {
		if b.y > -20 && b.y < screenHeight+20 {
			nb = append(nb, b)
		}
	}
	g.bullets = nb

	// Update enemies
	for i := 0; i < len(g.enemies); i++ {
		e := &g.enemies[i]
		e.x += e.vx
		e.y += e.vy
	}
	// Collision: bullets vs enemies
	ne := g.enemies[:0]
	for _, e := range g.enemies {
		hit := false
		for bi := range g.bullets {
			b := g.bullets[bi]
			if dist(e.x, e.y, b.x, b.y) < 20 {
				hit = true
				// spawn particles
				for i := 0; i < 12; i++ {
					g.particles = append(g.particles, Particle{e.x, e.y, (rand.Float64() - 0.5) * 4, (rand.Float64() - 0.5) * 4, 20 + rand.Intn(20), color.RGBA{0xff, 0xc0, 0x40, 0xff}})
				}
				g.score += 10
				audio.PlayPop()
				// remove one bullet (the first that hit)
				// mark bullet out of screen
				b.y = -9999
				if bi < len(g.bullets) {
					g.bullets[bi] = b
				}
				break
			}
		}
		if !hit && e.y < screenHeight+40 {
			ne = append(ne, e)
		}
	}
	g.enemies = ne

	// Update particles
	np := g.particles[:0]
	for _, p := range g.particles {
		p.x += p.vx
		p.y += p.vy
		p.vy += 0.05
		p.life--
		if p.life > 0 {
			np = append(np, p)
		}
	}
	g.particles = np

	// subtle background shift
	g.bgHue += 0.02
	if g.bgHue > 360 {
		g.bgHue = 0
	}

	return nil
}

// Draw the game screen.
func (g *Game) Draw(screen *ebiten.Image) {
	// Background
	bg := color.RGBA{0x10, 0x18, 0x30, 0xff}
	screen.Fill(bg)

	// Stars (procedural subtle)
	for i := 0; i < 30; i++ {
		x := (float64((g.frames*7 + i*13) % screenWidth))
		y := float64((i * 37) % screenHeight)
		col := color.RGBA{0xff, 0xff, 0xff, uint8(50 + (i%3)*40)}
		small := ebiten.NewImage(2, 2)
		small.Fill(col)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(small, op)
	}

	// Draw player
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.playerX-16, g.playerY-8)
	screen.DrawImage(g.playerImage, op)

	// Draw enemies
	for _, e := range g.enemies {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(e.x-14, e.y-8)
		screen.DrawImage(g.enemyImage, op)
	}

	// Draw bullets
	for _, b := range g.bullets {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(b.x-3, b.y-6)
		screen.DrawImage(g.bulletImage, op)
	}

	// Draw particles
	for _, p := range g.particles {
		s := ebiten.NewImage(3, 3)
		s.Fill(p.col)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(p.x, p.y)
		screen.DrawImage(s, op)
	}

	// HUD
	text.Draw(screen, "Score: "+itoa(g.score), basicfont.Face7x13, 8, 16, color.White)
	internal.DrawFPS(screen, g.frames)

	// small instruction
	ebitenutil.DebugPrintAt(screen, "Arrows/WASD to move — Space to shoot", 8, screenHeight-18)
}

// Layout reports the screen dimensions.
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func dist(x1, y1, x2, y2 float64) float64 {
	dx := x1 - x2
	dy := y1 - y2
	return math.Sqrt(dx*dx + dy*dy)
}

func itoa(v int) string { return fmtInt(v) }

// fmtInt is a tiny fast int->string helper (avoids fmt import noise).
func fmtInt(v int) string {
	buf := bytes.NewBuffer(make([]byte, 0, 16))
	if v == 0 {
		return "0"
	}
	neg := false
	if v < 0 {
		neg = true
		v = -v
	}
	digits := []byte{}
	for v > 0 {
		d := byte(v%10) + '0'
		digits = append(digits, d)
		v /= 10
	}
	if neg {
		digits = append(digits, '-')
	}
	// reverse
	for i := len(digits) - 1; i >= 0; i-- {
		buf.WriteByte(digits[i])
	}
	return buf.String()
}
