package assets

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/xyy0411/bleachVSnaruto/common/state"
)

func TestLoadCharacterAnimationConfig(t *testing.T) {
	withRepoRoot(t)

	cfg, err := LoadCharacterAnimationConfig("narutoS")
	if err != nil {
		t.Fatalf("LoadCharacterAnimationConfig returned error: %v", err)
	}

	if cfg.Meta.Texture != filepath.Join("assets", "characters", "narutoS", "texture_atlas.png") {
		t.Fatalf("unexpected texture path: %s", cfg.Meta.Texture)
	}

	if _, ok := cfg.Animations["idle"]; !ok {
		t.Fatalf("idle animation was not loaded")
	}
}

func TestBuildAnimationSetFromAtlas(t *testing.T) {
	withRepoRoot(t)

	cfg, err := LoadCharacterAnimationConfig("narutoS")
	if err != nil {
		t.Fatalf("LoadCharacterAnimationConfig returned error: %v", err)
	}

	set, err := BuildAnimationSetFromAtlas(
		"narutoS",
		cfg,
		map[string]int{"idle": 6, "run": 8, "jump": 10, "dash": 12, "just_landed": 10},
		map[string]bool{"idle": true, "run": true, "jump": true},
	)
	if err != nil {
		t.Fatalf("BuildAnimationSetFromAtlas returned error: %v", err)
	}

	idleAnim := set.ByState[state.Idle]
	if idleAnim == nil || len(idleAnim.FramesKeys) == 0 {
		t.Fatalf("idle animation was not built")
	}
	if len(idleAnim.Frames) != len(idleAnim.FramesKeys) {
		t.Fatalf("frame metadata count mismatch: got %d keys and %d frames", len(idleAnim.FramesKeys), len(idleAnim.Frames))
	}
	if idleAnim.Frames[0].Origin.X == 0 || idleAnim.Frames[0].Origin.Y == 0 {
		t.Fatalf("frame origin metadata was not preserved")
	}

	key := idleAnim.FramesKeys[0]
	if img := StdImagePool.GetImage(key); img == nil {
		t.Fatalf("image pool missing generated frame for key %s", key)
	}
}

func withRepoRoot(t *testing.T) {
	t.Helper()

	cwd, err := os.Getwd()
	if err != nil {
		t.Fatalf("Getwd failed: %v", err)
	}
	repoRoot := filepath.Dir(cwd)
	if err := os.Chdir(repoRoot); err != nil {
		t.Fatalf("Chdir(%q) failed: %v", repoRoot, err)
	}
	t.Cleanup(func() {
		_ = os.Chdir(cwd)
	})
}
