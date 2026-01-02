package time

type Time struct {
	GlobalFrame int64
	Delta       float64 // 秒
	TPS         float64 // 逻辑帧
}

func (t *Time) Tick() {
	t.GlobalFrame++
	t.Delta = 1.0 / t.TPS
}
