package assets

import (
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/xyy0411/bleachVSnaruto/global"
)

var (
	NilImage     = ebiten.NewImage(1, 1)
	StdImagePool = NewImagePool()
)

// ImagePool 图池数据结构
type ImagePool struct {
	mu          sync.RWMutex
	DefaultTTL  time.Duration //生命周期
	CleanupTick time.Duration // 清理间隔
	images      map[string]*Image
}

type Image struct {
	Meta     *ebiten.Image
	LastUsed time.Time
	LongTime bool //长期有效
}

// NewImagePool 创建一个新的图池实例
func NewImagePool() *ImagePool {
	pool := &ImagePool{
		DefaultTTL:  time.Minute * 5,
		CleanupTick: time.Minute * 6,
		images:      make(map[string]*Image),
	}
	go func(p *ImagePool) {
		for {
			time.Sleep(p.CleanupTick)
			p.CleanUp()
		}
	}(pool)
	return pool
}

// LoadImage 将本地图片存入池
func (p *ImagePool) LoadImage(path string, longTime ...bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	img, _, err := ebitenutil.NewImageFromFile(path)
	if err != nil {
		global.Logger.Fatal("load image failed:", path, err)
	}
	p.images[path] = &Image{Meta: img, LastUsed: time.Now(), LongTime: len(longTime) == 1 && longTime[0]}
}

func (p *ImagePool) LoadImageArray(arg ...string) {
	for _, path := range arg {
		p.LoadImage(path)
	}
}

func (p *ImagePool) LoadLongTimeImageArray(arg ...string) {
	for _, path := range arg {
		p.LoadImage(path, true)
	}
}

// PostImage 存入图片
func (p *ImagePool) PostImage(key string, img *ebiten.Image, longTime ...bool) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.images[key] = &Image{Meta: img, LastUsed: time.Now(), LongTime: len(longTime) == 1 && longTime[0]}
}

// GetImage 从图片池返回图片,如果不存在尝试本地获取,如果依旧不存在则返回NilImage
func (p *ImagePool) GetImage(key string, longTime ...bool) *ebiten.Image {
	p.mu.Lock()
	defer p.mu.Unlock()
	if image, ok := p.images[key]; ok {
		image.LastUsed = time.Now()
		return image.Meta
	}
	img, _, err := ebitenutil.NewImageFromFile(key)
	if err != nil {
		global.Logger.Fatal("load image failed:", key, err)
		return NilImage
	}
	p.images[key] = &Image{Meta: img, LastUsed: time.Now(), LongTime: len(longTime) == 1 && longTime[0]}
	return img
}

// DeleteImage 从图池中移除指定键的图像
func (p *ImagePool) DeleteImage(key string) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if image, ok := p.images[key]; ok {
		image.Meta.Deallocate()
	 	delete(p.images, key)
	}
}

// Clear 清空整个图池
func (p *ImagePool) Clear() {
	p.images = make(map[string]*Image)
}

// Size 返回图池中图像的数量
func (p *ImagePool) Size() int {
	return len(p.images)
}

// CleanUp 清理过期图片
func (p *ImagePool) CleanUp() {
	expirytime := time.Now().Add(-p.CleanupTick)
	for key, elem := range p.images {
		if elem.LongTime {
			continue
		}
		if expirytime.After(elem.LastUsed) {
			p.DeleteImage(key)
		}
	}
}
