package counters

//alterCounter 未公开的标识符
type alterCounter int

//工厂模式 Go中通常命名为New
//New创建一个公开的 alterCounter类型的值
func New(value int) alterCounter {
	return alterCounter(value)
}
