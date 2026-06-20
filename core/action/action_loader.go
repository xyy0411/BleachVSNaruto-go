package action

import (
	"encoding/json"
	"os"
)

// ConfigFile 动作配置文件结构
type ConfigFile struct {
	Character string            `json:"character"`
	Actions   map[string]Action `json:"actions"`
}

// Action 动作配置
type Action struct {
	DurationFrames int            `json:"durationFrames"` // 总帧数
	FPS            int            `json:"fps"`            // 帧率
	Loop           bool           `json:"loop"`           // 是否循环
	CancelWindows  []CancelWindow `json:"cancelWindows"`  // 可取消窗口列表
}

// CancelWindow 可取消窗口配置（简化版）
type CancelWindow struct {
	Input    string `json:"input"`
	MinFrame int    `json:"minFrame"`
	MaxFrame int    `json:"maxFrame"`
}

// DurationSeconds 返回动作持续时间（秒）
func (a *Action) DurationSeconds() float64 {
	if a.FPS <= 0 {
		return 0
	}
	return float64(a.DurationFrames) / float64(a.FPS)
}

// LoadActionConfig 从JSON文件加载动作配置
func LoadActionConfig(path string) (*ConfigFile, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ConfigFile
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	return &config, nil
}

// GetAction 获取指定动作配置
func (c *ConfigFile) GetAction(name string) (Action, bool) {
	action, ok := c.Actions[name]
	return action, ok
}
