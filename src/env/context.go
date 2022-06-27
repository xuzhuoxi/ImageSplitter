// Create on 2022/6/26
// @author xuzhuoxi
package env

import (
	"errors"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"image"
	"math"
	"strconv"
	"strings"
)

type SplitContext struct {
	EnvPath string // 【**可选**】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
	Mode    string // 【**必要**】分割模式，支持：fixed(1)、avg(2)
	Order   string // 【**必要**】分割顺序，支持：LeftUp(1)、LeftDown(2)
	Size    string // 【**必要**】分割参数，格式 "mxn","mXn","m*n"
	Trim    string // 【**必要**】尾部裁剪，支持：on、off

	InImagePath  string // 【**必要**】来源图片路径,要求为图片格式的文件
	OutImagePath string // 【**必要**】输出图片信息，要求使用通配符目录,

	Format      string // 【**非必要**】强制指定图像文件格式, 如果不指定，将使用源图像的格式
	FormatRatio int    // 【**非必要**】强制指定图像文件质量(如有必要)，如果不指定，使用默认值85

	FlagSize             Size
	ImageSize            image.Point
	SliceSize, CountSize Size
	FormatExt            string // 图像格式对应的扩展名
}

func (c *SplitContext) InitContext() error {
	if err := c.initEnv(); nil != err {
		return err
	}
	if err := c.checkMode(); nil != err {
		return err
	}
	if err := c.checkOrder(); nil != err {
		return err
	}
	if err := c.checkInImage(); nil != err {
		return err
	}
	if err := c.initOutImagePath(); nil != err {
		return err
	}
	if err := c.checkFormat(); nil != err {
		return err
	}
	if err := c.initFlagSize(); nil != err {
		return err
	}
	if err := c.checkEndCrop(); nil != err {
		return err
	}
	return nil
}

func (c *SplitContext) GetMode() CmdFlag {
	if ModeSize.MatchCase(c.Mode) {
		return ModeSize
	}
	return ModeAvg
}

func (c *SplitContext) GetOrder() CmdFlag {
	if OrderLeftUp.MatchCase(c.Order) {
		return OrderLeftUp
	}
	return OrderLeftDown
}

func (c *SplitContext) GetTrim() CmdFlag {
	if EndTrimOff.MatchCase(c.Trim) {
		return EndTrimOff
	}
	return EndTrimOn
}

func (c *SplitContext) SetImageSize(imageSize image.Point) {
	c.ImageSize = imageSize
	if c.checkFlag(ModeSize, c.Mode) {
		c.SliceSize, c.CountSize = c.parseSizeBySizeMode(c.FlagSize, c.ImageSize)
	} else {
		c.SliceSize, c.CountSize = c.parseSizeByAvgMode(c.FlagSize, c.ImageSize)
	}
}

func (c *SplitContext) SetDefaultFormat(format string) {
	if c.Format == "" {
		c.Format = format
	}
	c.FormatExt = formatx.GetExtName(c.Format)
}

func (c *SplitContext) GetSrcPoint(xIndex int, yIndex int) (srcPoint image.Point) {
	if c.checkFlag(OrderLeftUp, c.Order) {
		return c.getLeftUpSrcPoint(xIndex, yIndex)
	}
	if c.checkFlag(OrderLeftDown, c.Order) {
		return c.getLeftDownSrcPoint(xIndex, yIndex)
	}
	return
}

func (c *SplitContext) initEnv() error {
	if c.EnvPath != "" {
		if filex.IsDir(c.EnvPath) {
			return nil
		}
		newEnv := filex.Combine(osxu.GetRunningDir(), c.EnvPath)
		if filex.IsDir(newEnv) {
			c.EnvPath = newEnv
			return nil
		}
		return errors.New(fmt.Sprintf("evn[%s] is not exist!", c.EnvPath))
	}
	c.EnvPath = osxu.GetRunningDir()
	return nil
}

func (c *SplitContext) checkMode() error {
	if c.checkFlag(ModeSize, c.Mode) || c.checkFlag(ModeAvg, c.Mode) {
		return nil
	}
	return errors.New(fmt.Sprintf("mode unknown[%s]. ", c.Mode))
}

func (c *SplitContext) checkOrder() error {
	if c.checkFlag(OrderLeftUp, c.Order) || c.checkFlag(OrderLeftDown, c.Order) {
		return nil
	}
	return errors.New(fmt.Sprintf("order unknown[%s]. ", c.Order))
}

func (c *SplitContext) checkEndCrop() error {
	if c.checkFlag(EndTrimOn, c.Trim) || c.checkFlag(EndTrimOff, c.Trim) {
		return nil
	}
	return errors.New(fmt.Sprintf("crop unknown[%s]. ", c.Trim))
}

func (c *SplitContext) checkInImage() error {
	if filex.IsFile(c.InImagePath) {
		return nil
	}
	newPath := filex.Combine(osxu.GetRunningDir(), c.InImagePath)
	if filex.IsFile(newPath) {
		c.InImagePath = newPath
		return nil
	}
	return errors.New(fmt.Sprintf("in image[%s] is not exist!", c.InImagePath))
}

func (c *SplitContext) initOutImagePath() error {
	upDir, _ := filex.GetUpDir(c.OutImagePath)
	if filex.IsDir(upDir) {
		return nil
	}
	c.OutImagePath = filex.Combine(c.EnvPath, c.OutImagePath)
	return nil
}

func (c *SplitContext) checkFormat() error {
	if "" == c.Format {
		return nil
	}
	if !formatx.CheckFormatRegistered(c.Format) {
		return errors.New(fmt.Sprintf("format unknown[%s]", c.Format))
	}
	return nil
}

func (c *SplitContext) initFlagSize() error {
	size := strings.ToLower(c.Size)
	var arr []string
	if strings.Contains(size, "x") {
		arr = strings.Split(size, "x")
	} else if strings.Contains(size, "*") {
		arr = strings.Split(size, "*")
	} else {
		return errors.New(fmt.Sprintf("size unknown[%s]. ", c.Size))
	}
	x, err := strconv.ParseUint(arr[0], 10, 32)
	if nil != err {
		return errors.New(fmt.Sprintf("size unknown[%s][%s]. ", c.Size, err))
	}
	y, err := strconv.ParseUint(arr[0], 10, 32)
	if nil != err {
		return errors.New(fmt.Sprintf("size unknown[%s][%s]. ", c.Size, err))
	}
	c.FlagSize = Size{Width: int(x), Height: int(y)}
	return nil
}

func (c *SplitContext) checkFlag(flag CmdFlag, value string) bool {
	return flag.MatchCase(value)
}

func (c *SplitContext) parseSizeBySizeMode(flagSize Size, imageSize image.Point) (sliceSize Size, countSize Size) {
	sliceSize = flagSize
	w := math.Ceil(float64(imageSize.X) / float64(flagSize.Width))
	h := math.Ceil(float64(imageSize.Y) / float64(flagSize.Height))
	countSize = Size{Width: int(w), Height: int(h)}
	return
}

func (c *SplitContext) parseSizeByAvgMode(flagSize Size, imageSize image.Point) (sliceSize Size, countSize Size) {
	countSize = flagSize
	w := math.Ceil(float64(imageSize.X) / float64(flagSize.Width))
	h := math.Ceil(float64(imageSize.Y) / float64(flagSize.Height))
	sliceSize = Size{Width: int(w), Height: int(h)}
	return
}

func (c *SplitContext) getLeftUpSrcPoint(xIndex int, yIndex int) image.Point {
	x := xIndex * c.SliceSize.Width
	y := yIndex * c.SliceSize.Height
	return image.Point{X: x, Y: y}
}

func (c *SplitContext) getLeftDownSrcPoint(xIndex int, yIndex int) image.Point {
	x := xIndex * c.SliceSize.Width
	y := c.ImageSize.Y - (yIndex+1)*c.SliceSize.Height
	return image.Point{X: x, Y: y}
}
