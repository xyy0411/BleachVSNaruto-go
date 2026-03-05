package characters

import "github.com/xyy0411/bleachVSnaruto/core/charactor"

var CharList = map[string]func() charactor.Character{}

func SelectChar(id string) func() charactor.Character {
	return CharList[id]
}

func AddChar(id string, newFunc func() charactor.Character) {
	CharList[id] = newFunc
}

func RangeChar(f func(id string, newFunc func() charactor.Character)) {
	for k, v := range CharList {
		f(k, v)
	}
}
