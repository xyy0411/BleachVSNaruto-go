package input

import "github.com/xyy0411/bleachVSnaruto/models"

// Source 定义了输入源的接口
// 用于从不同来源读取输入帧数据
type Source interface {
	// Read 读取一个输入帧
	// 返回输入帧数据的快照
	Read() models.InputFrame
}
