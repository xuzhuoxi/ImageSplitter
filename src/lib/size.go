package lib

import "fmt"

type SplitSize struct {
	//宽
	Width int
	//高
	Height int
}

func (s SplitSize) String() string {
	return fmt.Sprintf("{Width=%d, Height=%d}", s.Width, s.Height)
}
