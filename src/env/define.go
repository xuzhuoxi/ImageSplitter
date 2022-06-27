// Create on 2022/6/26
// @author xuzhuoxi
package env

import (
	"fmt"
	"strings"
)

var (
	ModeSize = CmdFlag{core: []string{"fixed", "1"}, Desc: "固定尺寸[Fixed]"} //固定分割，不足的补空
	ModeAvg  = CmdFlag{core: []string{"avg", "2"}, Desc: "固定数量[Avg]"}     //平均分割，根据图片总大小进行水平与垂直的平均分割
)

var (
	OrderLeftUp    = CmdFlag{core: []string{"lu", "leftup", "1"}, Desc: "左上[LeftUp]"}       // 左上角起始位
	OrderLeftDown  = CmdFlag{core: []string{"ld", "leftdown", "2"}, Desc: "左下[LeftDown]"}   // 左下角起始位
	OrderRightUp   = CmdFlag{core: []string{"ru", "rightup", "3"}, Desc: "右上[RightUp]"}     // 右上角起始位
	OrderRightDown = CmdFlag{core: []string{"rd", "rightdown", "4"}, Desc: "右下[RightDown]"} // 右下角起始位
)

var (
	EndTrimOff = CmdFlag{core: []string{"off", "false", "disable", "0"}, Desc: "关闭裁剪[TrimOff]"} // 关闭裁剪，不足补空
	EndTrimOn  = CmdFlag{core: []string{"on", "true", "enable", "1"}, Desc: "启用裁剪[TrimOn]"}     // 启用裁剪
)

var (
	WildcardN0 = []string{"{n0}", "{N0}"} // 从0开始的分割顺序数。
	WildcardN1 = []string{"{n1}", "{N1}"} // 从1开始的分割顺序数。

	WildcardX0 = []string{"{x0}", "{X0}"} // 从0开始的水平方向分割顺序数。
	WildcardX1 = []string{"{x1}", "{X1}"} // 从1开始的水平方向分割顺序数。

	WildcardY0 = []string{"{y0}", "{Y0}"} // 从0开始的垂直方向分割顺序数。
	WildcardY1 = []string{"{y1}", "{Y1}"} // 从1开始的垂直方向分割顺序数。

	WildcardExt = []string{"{ext}"}
)

const (
	DefaultRatio = 85
)

type Size struct {
	Width  int // 宽 | 水平数量
	Height int // 高 | 垂直数量
}

func (s Size) String() string {
	return fmt.Sprintf("{Width=%d, Height=%d}", s.Width, s.Height)
}

type CmdFlag struct {
	core []string
	Desc string
}

func (o CmdFlag) Match(value string) bool {
	value = strings.ToLower(value)
	return o.MatchCase(value)
}

func (o CmdFlag) MatchCase(value string) bool {
	for index := range o.core {
		if o.core[index] == value {
			return true
		}
	}
	return false
}

func (o CmdFlag) DefaultValue() string {
	return o.core[0]
}
