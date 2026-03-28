package assets

import (
	"fmt"
	"os"
	"sort"
)

func LoadCharacterActionFrames(roleId string, types string) (frames []string) {
	uri := fmt.Sprintf("assets/characters/%s/animation/%s/", roleId, types)
	entries, _ := os.ReadDir(uri)
	for i := 0; i < len(entries); i++ {
		frames = append(frames, uri+entries[i].Name())
	}
	sort.Strings(frames)
	return
}
