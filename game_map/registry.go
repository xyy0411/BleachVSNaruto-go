package game_map

import "github.com/hajimehoshi/ebiten/v2"

var StdRegistry = NewRegistry()

type MapInter interface {
	Draw(*ebiten.Image, float64)
	Init()
	GetBaseInfo() BaseInfo
}

type BaseInfo struct {
	BirdViewKey string
	ID          string
}

type MapRegistry map[string]MapInter

//  新建地图注册中心实例
func NewRegistry() MapRegistry {
	return make(map[string]MapInter)

}

func (r MapRegistry) RegisterMap(id string, mapImpl MapInter) {
	r[id] = mapImpl
}
