package action

import "github.com/xyy0411/bleachVSnaruto/core/input"

// Context 动作上下文
type Context struct {
	Action string
	Stage  int
	Frame  int
}

// CancelWindow 可取消窗口
type CancelWindow struct {
	MinFrame int
	MaxFrame int
}

// Transition 动作转移
type Transition struct {
	FromAction   string
	FromStage    int
	Input        string
	ToAction     string
	ToStage      int
	CancelWindow CancelWindow
	BufferBefore int
}

// ComboEntry 连招条目
type ComboEntry struct {
	Command      string
	Sequence     []string
	MaxGapFrames int
	Repeatable   bool
}

// ComboTable 连招表
type ComboTable struct {
	Entries     []ComboEntry
	Transitions []Transition
}

// NewComboTable 创建新的连招表
func NewComboTable() *ComboTable {
	return &ComboTable{
		Entries:     make([]ComboEntry, 0),
		Transitions: make([]Transition, 0),
	}
}

// AddEntry 添加连招条目
func (t *ComboTable) AddEntry(entry ComboEntry) {
	t.Entries = append(t.Entries, entry)
}

// AddTransition 添加转移
func (t *ComboTable) AddTransition(trans Transition) {
	t.Transitions = append(t.Transitions, trans)
}

// Resolve 解析连招
func (t *ComboTable) Resolve(ctx Context, buffer *input.Buffer, currentFrame int64) *Transition {
	// 检查转移
	for _, trans := range t.Transitions {
		if trans.FromAction == ctx.Action && trans.FromStage == ctx.Stage {
			if trans.Input != "" && buffer.HasRecentPress(trans.Input, currentFrame-int64(trans.BufferBefore)) {
				return &trans
			}
		}
	}

	return nil
}