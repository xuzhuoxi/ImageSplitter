package lib

// 分割模式
type SplitMode int

const (
	// 固定分割，不足的补空
	SizeMode SplitMode = iota + 1

	// 平均分割，根据图片总大小进行水平与垂直的平均分割
	AvgMode
)
