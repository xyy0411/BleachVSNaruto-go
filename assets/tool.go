package assets

import (
	"fmt"
	"os"
)

func LoadCharacterActionFrames(roleId string, types string) (frames []string) {
	uri := fmt.Sprintf("assets/characters/%s/animation/%s/", roleId, types)
	entries, _ := os.ReadDir(uri)
	for _, entry := range entries {
		frames = append(frames, uri+entry.Name())
	}
	return
}
