package assets

import (
	"bytes"
	"encoding/json"
	"fmt"
	"image"
	"os"
	"path/filepath"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/xyy0411/bleachVSnaruto/common/state"
	"github.com/xyy0411/bleachVSnaruto/core/animatable"
)

const defaultAnimationFPS = 10

// LoadCharacterAnimationConfig 读取角色动画配置文件并归一化图集路径。
func LoadCharacterAnimationConfig(roleID string) (*animatable.FullAnimationConfig, error) {
	configPath := filepath.Join("assets", "characters", roleID, "config", "animation_config.full.json")
	data, err := os.ReadFile(configPath)
	if err != nil {
		return nil, fmt.Errorf("read character animation config %q: %w", configPath, err)
	}
	data = bytes.TrimPrefix(data, []byte{0xEF, 0xBB, 0xBF})

	var cfg animatable.FullAnimationConfig
	if err := json.Unmarshal(data, &cfg); err != nil {
		return nil, fmt.Errorf("decode character animation config %q: %w", configPath, err)
	}

	texturePath, err := normalizeCharacterTexturePath(roleID, cfg.Meta.Texture)
	if err != nil {
		return nil, err
	}
	cfg.Meta.Texture = texturePath

	return &cfg, nil
}

// BuildAnimationSetFromAtlas 将角色图集配置预裁切为当前播放链可直接使用的动画集合
func BuildAnimationSetFromAtlas(roleID string, cfg *animatable.FullAnimationConfig, fpsByAction map[string]int, loopByAction map[string]bool) (animatable.Set, error) {
	if cfg == nil {
		return animatable.Set{}, fmt.Errorf("animation config is nil")
	}

	atlas, _, err := ebitenutil.NewImageFromFile(cfg.Meta.Texture)
	if err != nil {
		return animatable.Set{}, fmt.Errorf("load atlas texture %q: %w", cfg.Meta.Texture, err)
	}

	byState := make(map[state.State]*animatable.ActionAnimation)
	for st := range state.NumOfState {
		actionName := st.String()
		actionCfg, ok := cfg.Animations[actionName]
		if !ok {
			continue
		}

		framesKeys, err := buildAtlasFrames(roleID, actionName, atlas, actionCfg)
		if err != nil {
			return animatable.Set{}, err
		}

		byState[st] = &animatable.ActionAnimation{
			FramesKeys: framesKeys,
			Frames:     actionCfg.Frames,
			FPS:        resolveAnimationFPS(actionName, fpsByAction),
			Loop:       resolveAnimationLoop(actionName, actionCfg.Loop, loopByAction),
		}
	}

	return animatable.Set{ByState: byState}, nil
}

func buildAtlasFrames(roleID, actionName string, atlas *ebiten.Image, actionCfg animatable.AnimationConfig) ([]string, error) {
	if len(actionCfg.Frames) == 0 {
		return nil, fmt.Errorf("action %q has no frames", actionName)
	}

	framesKeys := make([]string, 0, len(actionCfg.Frames))
	atlasBounds := atlas.Bounds()
	for i, frameCfg := range actionCfg.Frames {
		frameRect := image.Rect(
			frameCfg.Rect.X,
			frameCfg.Rect.Y,
			frameCfg.Rect.X+frameCfg.Rect.W,
			frameCfg.Rect.Y+frameCfg.Rect.H,
		)
		if !frameRect.In(atlasBounds) && frameRect != atlasBounds {
			if frameRect.Min.X < atlasBounds.Min.X || frameRect.Min.Y < atlasBounds.Min.Y || frameRect.Max.X > atlasBounds.Max.X || frameRect.Max.Y > atlasBounds.Max.Y {
				return nil, fmt.Errorf("frame %d for action %q exceeds atlas bounds", i, actionName)
			}
		}

		subImage := atlas.SubImage(frameRect)
		ebitenImage, ok := subImage.(*ebiten.Image)
		if !ok {
			return nil, fmt.Errorf("frame %d for action %q is not an ebiten image", i, actionName)
		}

		key := atlasFrameKey(roleID, actionName, i)
		StdImagePool.PostImage(key, ebitenImage, true)
		framesKeys = append(framesKeys, key)
	}

	return framesKeys, nil
}

func atlasFrameKey(roleID, actionName string, frameIndex int) string {
	return fmt.Sprintf("char:%s:anim:%s:frame:%d", roleID, actionName, frameIndex)
}

func resolveAnimationFPS(actionName string, fpsByAction map[string]int) int {
	if fps, ok := fpsByAction[actionName]; ok && fps > 0 {
		return fps
	}

	return defaultAnimationFPS
}

func resolveAnimationLoop(actionName string, defaultLoop bool, loopByAction map[string]bool) bool {
	if loop, ok := loopByAction[actionName]; ok {
		return loop
	}

	return defaultLoop
}

func normalizeCharacterTexturePath(roleID, configuredPath string) (string, error) {
	preferredPath := filepath.Join("assets", "characters", roleID, "texture_atlas.png")
	if _, err := os.Stat(preferredPath); err == nil {
		return preferredPath, nil
	}

	if configuredPath != "" {
		if _, err := os.Stat(configuredPath); err == nil {
			return configuredPath, nil
		}
	}

	return "", fmt.Errorf("character %q texture not found, preferred=%q configured=%q", roleID, preferredPath, configuredPath)
}
