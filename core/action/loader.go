package action

import (
	"encoding/json"
	"os"
)

// ComboConfigFile JSON配置文件结构
type ComboConfigFile struct {
	Character        string               `json:"character"`
	ActionConfigPath string               `json:"actionConfigPath"`
	Entries          []ComboEntryConfig   `json:"entries"`
	Transitions      []TransitionConfig   `json:"transitions"`
	ComboCommands    []ComboCommandConfig `json:"comboCommands"`
}

// ComboEntryConfig 连招条目配置
type ComboEntryConfig struct {
	Command      string   `json:"command"`
	Sequence     []string `json:"sequence"`
	MaxGapFrames int      `json:"maxGapFrames"`
	Repeatable   bool     `json:"repeatable"`
}

// TransitionConfig 转移配置
type TransitionConfig struct {
	FromAction   string `json:"fromAction"`
	FromStage    int    `json:"fromStage"`
	Input        string `json:"input"`
	ToAction     string `json:"toAction"`
	ToStage      int    `json:"toStage"`
	BufferBefore int    `json:"bufferBefore"`
}

// ComboCommandConfig 组合命令配置
type ComboCommandConfig struct {
	Command      string   `json:"command"`
	Sequence     []string `json:"sequence"`
	MaxGapFrames int      `json:"maxGapFrames"`
	ToAction     string   `json:"toAction"`
	ToStage      int      `json:"toStage"`
}

// LoadComboTableFromJSON 从JSON文件加载连招表
func LoadComboTableFromJSON(path string) (*ComboTable, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var config ComboConfigFile
	if err := json.Unmarshal(data, &config); err != nil {
		return nil, err
	}

	table := NewComboTable()

	// 加载连招条目
	for _, entry := range config.Entries {
		table.AddEntry(ComboEntry{
			Command:      entry.Command,
			Sequence:     entry.Sequence,
			MaxGapFrames: entry.MaxGapFrames,
			Repeatable:   entry.Repeatable,
		})
	}

	// 尝试加载动作配置（用于获取可取消窗口）
	var actionConfig *ConfigFile
	if config.ActionConfigPath != "" {
		actionConfig, _ = LoadActionConfig(config.ActionConfigPath)
	}

	// 加载状态转移
	for _, trans := range config.Transitions {
		// 获取可取消窗口（优先从动作配置获取）
		cancelWindow := CancelWindow{MinFrame: 0, MaxFrame: 0}

		if actionConfig != nil {
			if action, exists := actionConfig.GetAction(trans.FromAction); exists {
				for _, cw := range action.CancelWindows {
					if cw.Input == trans.Input {
						cancelWindow = CancelWindow{
							MinFrame: cw.MinFrame,
							MaxFrame: cw.MaxFrame,
						}
						break
					}
				}
			}
		}

		table.AddTransition(Transition{
			FromAction:   trans.FromAction,
			FromStage:    trans.FromStage,
			Input:        trans.Input,
			ToAction:     trans.ToAction,
			ToStage:      trans.ToStage,
			CancelWindow: cancelWindow,
			BufferBefore: trans.BufferBefore,
		})
	}

	// 加载组合命令（生成对应的转移）
	for _, cmd := range config.ComboCommands {
		table.AddTransition(Transition{
			FromAction:   "idle",
			FromStage:    0,
			Input:        cmd.Command,
			ToAction:     cmd.ToAction,
			ToStage:      cmd.ToStage,
			CancelWindow: CancelWindow{MinFrame: 0, MaxFrame: 0},
			BufferBefore: 6,
		})
	}

	return table, nil
}
