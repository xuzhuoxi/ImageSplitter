package lib

type SplitOrder int

const (
	// 左上角起始
	LeftUp SplitOrder = iota + 1

	// 左下角起始
	LeftDown
)
