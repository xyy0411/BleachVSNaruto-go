package world

// Camera 表示战斗场景中的横向摄像机状态与缩放参数。
type Camera struct {
	ViewportWidth  float64
	ViewportHeight float64
	X              float64
	Zoom           float64
	MinZoom        float64
	MaxZoom        float64
	FocusPadding   float64
}

// ClampX 将摄像机的 X 位置限制在地图边界内，避免视口越界。
func (c *Camera) ClampX(bound Bound) {
	if c == nil {
		return
	}

	visibleWidth := c.VisibleWidth()
	maxX := bound.Right - visibleWidth
	if maxX < bound.Left {
		maxX = bound.Left
	}

	if c.X < bound.Left {
		c.X = bound.Left
	}
	if c.X > maxX {
		c.X = maxX
	}
}

// VisibleWidth 返回当前缩放下摄像机在世界坐标中可见的宽度。
func (c *Camera) VisibleWidth() float64 {
	if c == nil {
		return 0
	}
	return c.ViewportWidth / c.Scale()
}

// ScreenOffsetY 返回缩放后在屏幕 Y 方向上的居中偏移量。
func (c *Camera) ScreenOffsetY() float64 {
	if c == nil {
		return 0
	}
	return (c.ViewportHeight - c.ViewportHeight*c.Scale()) / 2
}

// ApplyZoomLimit 根据配置与地图边界限制摄像机缩放值。
func (c *Camera) ApplyZoomLimit(bound Bound) {
	if c == nil {
		return
	}

	minZoom := c.MinZoom
	if minZoom <= 0 && bound.Right > bound.Left {
		minZoom = c.ViewportWidth / (bound.Right - bound.Left)
	}
	if minZoom <= 0 {
		minZoom = 1
	}

	maxZoom := c.MaxZoom
	if maxZoom <= 0 {
		maxZoom = 1
	}

	if c.Zoom < minZoom {
		c.Zoom = minZoom
	}
	if c.Zoom > maxZoom {
		c.Zoom = maxZoom
	}
}

// FollowTargets 根据多个目标点自动调整摄像机位置与缩放。
func (c *Camera) FollowTargets(bound Bound, targets ...float64) {
	if c == nil || len(targets) == 0 {
		return
	}

	minX, maxX := targets[0], targets[0]
	var sum float64
	for _, target := range targets {
		sum += target
		if target < minX {
			minX = target
		}
		if target > maxX {
			maxX = target
		}
	}

	padding := c.FocusPadding
	if padding <= 0 {
		padding = 160
	}

	// 两个角色拉开时扩大可视范围
	requiredWidth := (maxX - minX) + padding*2
	if requiredWidth <= 0 {
		requiredWidth = c.ViewportWidth
	}

	c.Zoom = c.ViewportWidth / requiredWidth
	c.ApplyZoomLimit(bound)

	targetX := sum / float64(len(targets))
	c.X = targetX - c.VisibleWidth()/2
	c.ClampX(bound)
}

// Scale 返回当前摄像机缩放倍率，异常值时回退为 1。
func (c *Camera) Scale() float64 {
	if c == nil || c.Zoom <= 0 {
		return 1
	}
	return c.Zoom
}

// WorldToScreen 将世界坐标转换为屏幕坐标。
func (c *Camera) WorldToScreen(x, y float64) (float64, float64) {
	scale := c.Scale()
	if c == nil {
		return x, y
	}
	return (x - c.X) * scale, y*scale + c.ScreenOffsetY()
}
