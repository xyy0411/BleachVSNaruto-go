package game_map

import "github.com/hajimehoshi/ebiten/v2"

type MapInter interface {
	Draw(*ebiten.Image, float64)
	Init()
	GetBaseInfo() BaseInfo
}

type BaseInfo struct {
	BirdView *ebiten.Image
}

type MapRegistry struct {
	maps map[string]MapInter
}

var registry *MapRegistry

// GetRegistry 获取地图注册中心实例
func GetRegistry() *MapRegistry {
	if registry == nil {
		registry = &MapRegistry{
			maps: make(map[string]MapInter),
		}
	}
	return registry
}

func (r *MapRegistry) RegisterMap(id string, mapImpl MapInter) {
	r.maps[id] = mapImpl
}
