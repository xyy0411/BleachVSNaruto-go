package render

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
		loadImage("assets/images/8.png"),
		loadImage("assets/images/10.png"),
		loadImage("assets/images/12.png"),
		loadImage("assets/images/14.png"),
	}
	Assets.RunFrames = []*ebiten.Image{
		loadImage("assets/images/17.png"),
		loadImage("assets/images/19.png"),
		loadImage("assets/images/21.png"),
		loadImage("assets/images/23.png"),
		loadImage("assets/images/25.png"),
		loadImage("assets/images/27.png"),
		loadImage("assets/images/29.png"),
	}
	Assets.JumpFrames = []*ebiten.Image{
		loadImage("assets/images/39.png"),
		loadImage("assets/images/41.png"),
		loadImage("assets/images/43.png"),
	}
	Assets.LandFrames = []*ebiten.Image{
		loadImage("assets/images/45.png"),
		loadImage("assets/images/47.png"),
		loadImage("assets/images/49.png"),
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

// DrawSprite 根据朝向绘制一帧
func DrawSprite(screen *ebiten.Image, img *ebiten.Image, x, y float64, facing int) {
	op := &ebiten.DrawImageOptions{}

	if facing == -1 {
		w := img.Bounds().Dx()
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(w), 0)
	}

	op.GeoM.Translate(x, y-float64(img.Bounds().Dy()))
	screen.DrawImage(img, op)
}
