package assets

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// AssetSet 保存游戏所需的贴图资源
type AssetSet struct {
	IdleFrames   []*ebiten.Image // 站立动画帧
	RunFrames    []*ebiten.Image // 跑步动画帧
	AttackFrames []*ebiten.Image // 攻击动画帧
	JumpFrames   []*ebiten.Image // 跳跃动画帧
	LandFrames   []*ebiten.Image // 落地动画帧
}

// Assets 全局唯一的资源容器
var Assets AssetSet

// LoadAssets 从本地文件加载图片资源
func LoadAssets() {
	Assets.IdleFrames = []*ebiten.Image{
		loadImage("character/shapes/8.png"),
		loadImage("character/shapes/10.png"),
		loadImage("character/shapes/12.png"),
		loadImage("character/shapes/14.png"),
	}
	Assets.RunFrames = []*ebiten.Image{
		loadImage("character/shapes/17.png"),
		loadImage("character/shapes/19.png"),
		loadImage("character/shapes/21.png"),
		loadImage("character/shapes/23.png"),
		loadImage("character/shapes/25.png"),
		loadImage("character/shapes/27.png"),
		loadImage("character/shapes/29.png"),
	}
	Assets.JumpFrames = []*ebiten.Image{
		loadImage("character/shapes/39.png"),
		loadImage("character/shapes/41.png"),
		loadImage("character/shapes/43.png"),
	}
	Assets.LandFrames = []*ebiten.Image{
		loadImage("character/shapes/45.png"),
		loadImage("character/shapes/47.png"),
		loadImage("character/shapes/49.png"),
	}

	// 预留攻击动画资源，后续可按需补充
	Assets.AttackFrames = []*ebiten.Image{}
}

// loadImage 读取单张图片，失败时直接终止
func loadImage(path string) *ebiten.Image {
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		log.Fatalf("加载图片失败: %s: %v", path, err)
	}
	return img
}
