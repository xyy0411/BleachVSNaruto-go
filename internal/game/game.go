package game

import (
	"fmt"
	"image/color"

	"ebiten_paractice/internal/assets"
	"ebiten_paractice/internal/character"
	"ebiten_paractice/internal/world"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game 游戏主结构
type Game struct {
	player *character.Character // 玩家角色实例
}

// New 创建并初始化游戏实例
func New() *Game {
	assets.LoadAssets()
	player := character.NewCharacter(200, characterGround())
	return &Game{player: player}
}

// Update 每帧更新逻辑
func (g *Game) Update() error {
	g.player.Update()
	return nil
}

// Draw 每帧绘制
func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 30, G: 30, B: 40, A: 255})
	g.player.Draw(screen)

	ebitenutil.DebugPrint(screen, "← → 移动角色，蓝色=Idle，绿色=Run")
	ebitenutil.DebugPrint(screen, fmt.Sprintf("\nFPS: %.2f\nTPS: %.2f", ebiten.ActualFPS(), ebiten.ActualTPS()))
}

// Layout 设置逻辑画布大小
func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 800, 600
}

// characterGround 提供角色初始落点
func characterGround() float64 {
	return world.GroundY
}
