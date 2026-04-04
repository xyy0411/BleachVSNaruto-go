package animatable

// Animatable ...
type Animatable interface {
	AnimationPlayer() AnimationPlayer
}

// AnimationPlayer 定义动画播放器的最小更新接口
type AnimationPlayer interface {
	Update(delta float64)
}

// ActionAnimation 动作动画，包含一个动作的完整动画信息
type ActionAnimation struct {
	// Frames 动画帧序列，存储该动作的所有帧图像 Key
	FramesKeys []string
	// Frames 保存每一帧的元数据，索引与 FramesKeys 一一对应
	Frames []AnimationFrame
	// FPS 表示动画每秒播放帧数
	FPS int
	// Loop 是否循环播放
	Loop bool
}

// Rect 用来记录一个动画图的图片的详细信息
// 比如坐标在(1,1)处为起始点往下延伸H，往右延伸W
type Rect struct {
	X int `json:"x"`
	Y int `json:"y"`
	W int `json:"w"`
	H int `json:"h"`
}

// Point 表示二维坐标点
type Point struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

// AnimationFrame 描述单帧动画在图集中的位置和锚点
type AnimationFrame struct {
	AtlasIndex int   `json:"atlasIndex"`
	Rect       Rect  `json:"rect"`
	Origin     Point `json:"origin"`
}

// AnimationConfig 描述单个动作的动画配置
type AnimationConfig struct {
	Start  int              `json:"start"`
	End    int              `json:"end"`
	Loop   bool             `json:"loop,omitempty"`
	Frames []AnimationFrame `json:"frames"`
}

// AnimationsMap 保存动作名到动画配置的映射
type AnimationsMap map[string]AnimationConfig

// Meta 保存动画配置文件的元数据
type Meta struct {
	Name    string `json:"name"`
	Texture string `json:"texture"`
}

// FullAnimationConfig 表示完整的动画配置文件结构
type FullAnimationConfig struct {
	Meta       Meta          `json:"meta"`
	Animations AnimationsMap `json:"animations"`
}
