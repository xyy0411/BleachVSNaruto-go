package app

import (
	"fmt"
	"github.com/xyy0411/ebiten_paractice/internal/character"
	"github.com/xyy0411/ebiten_paractice/internal/render"
	"github.com/xyy0411/ebiten_paractice/internal/world"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game 游戏主结构
type Game struct {
	player  *character.Character // 玩家角色实例
	terrain *world.Terrain       // 地形实例
}

// New 创建并初始化游戏实例
func New() *Game {
	render.LoadAssets()
	terrain := world.NewTerrain()
	player := character.NewCharacter(200, terrain.GroundY(0), terrain)
	return &Game{player: player, terrain: terrain}
}

// Update 每帧更新逻辑
func (g *Game) Update() error {
	character.GlobalFrame++
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
