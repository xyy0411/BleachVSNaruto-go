package time

type Time struct {
	GlobalFrame int64   //计数器
	Delta       float64 // 秒
	TPS         float64 // 逻辑帧
}

func (t *Time) Tick() {
	t.GlobalFrame++
}

func (t *Time) UpdataTPS(TPS float64) *Time {
	t.TPS = TPS
	t.Delta = 1.0 / TPS
	return t
}
