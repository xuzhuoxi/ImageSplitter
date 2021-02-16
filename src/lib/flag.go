package lib

import (
	"errors"
	"flag"
	"fmt"
	"github.com/xuzhuoxi/infra-go/filex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/osxu"
	"strconv"
	"strings"
)

const (
	N0Low = "{n0}"
	N0Up  = "{N0}"

	N1Low = "{n1}"
	N1Up  = "{N1}"

	X0Low = "{x0}"
	X0Up  = "{X0}"
	X1Low = "{x1}"
	X1Up  = "{X1}"

	Y0Low = "{y0}"
	Y0Up  = "{Y0}"
	Y1Low = "{y1}"
	Y1Up  = "{Y1}"
)

type FlagConfig struct {
	// 分割模式
	Mode SplitMode

	// 分割顺序
	// 支持：LeftUp(0)、LeftDown(1)
	Order SplitOrder

	// 分割大小
	// 固定分割模式下：每张小图片的大小
	// 平均分割模式下：分割后总数量(Width为水平方向数量，Height为垂直方向数量)
	Size SplitSize

	// 来源图片
	// 要求为图片模式的文件
	In string

	// 输出图片信息，要求使用通配符目录
	// 支持通配格式：0代表名称尾号从0开始、1代表名称尾号从1开始
	// 1. {n0}，{n1}：一维命名
	// 2. {m0}，{m1}，{n0}，{n1}：二维命名
	Out string

	// 输出文件格式
	Format string

	// 输出文件质量
	FormatRatio int
}

func (c *FlagConfig) IsAutoFormat() bool {
	return c.Format == string(formatx.Auto)
}

// -mode 		可选	分割模式(默认1)				1：小图使用固定尺寸；	2：小图使用平均尺寸
// -order 		可选	分割顺序(默认1)				1：左上角为起始点；	2：左下角为起始点
// -size 		必选	尺寸设置						格式：mxn。当mode为1时，m、n代表小图尺寸；当mode为2时，m、n代表分割数量
// -in 			必选	来源地址						字符串路径，文件夹或文件,"./"开头视为相对路径
// -out 		必选	输出地址						字符串路径，文件夹,"./"开头视为相对路径，支持通配符（{n0},{N0},{n1},{N1},{x0},{X0},{x1},{X1},{y0},{Y0},{y1},{Y1}）
// -format 		可选	输出文件格式(默认为in的格式)	图像格式[png,jpeg,jpg,jps]
// -ratio 		可选	压缩比(默认85)				整数(0,100]格式为jpeg、jpg,jps时有效
func ParseFlag() (cfg *FlagConfig, err error) {
	mode := flag.Int("mode", int(SizeMode), "SplitMode!")
	order := flag.Int("order", int(LeftUp), "SplitOrder!")
	size := flag.String("size", "", "Size FlagConfig!")

	in := flag.String("in", "", "Input Path! ")
	out := flag.String("out", "", "Output Path! ")

	format := flag.String("format", string(formatx.Auto), "Format FlagConfig!")
	ratio := flag.Int("ratio", 85, "Ratio FlagConfig!")

	flag.Parse()

	Mode := SplitMode(*mode)
	Order := SplitOrder(*order)

	sizes := strings.Split(strings.ToLower(*size), "x")
	if nil == sizes || len(sizes) != 2 {
		return nil, errors.New("Size Define Error! ")
	}
	w, _ := strconv.Atoi(sizes[0])
	h, _ := strconv.Atoi(sizes[1])
	Size := SplitSize{Width: w, Height: h}

	BasePath := osxu.GetRunningDir()
	In := filex.FormatPath(*in)
	if "" == In {
		return nil, errors.New("FlagConfig:in is empty! ")
	}
	if strings.Index(In, "./") == 0 {
		In = filex.Combine(BasePath, In)
	}
	if !filex.IsExist(In) {
		return nil, errors.New(fmt.Sprintf("FlagConfig:in(%s) is not exist! ", In))
	}
	if filex.IsFolder(In) {
		return nil, errors.New(fmt.Sprintf("FlagConfig:in(%s) is folder! ", In))
	}

	Out := filex.FormatPath(*out)
	if "" == Out {
		return nil, errors.New("FlagConfig:out is empty! ")
	}
	if strings.Index(Out, "./") == 0 {
		Out = filex.Combine(BasePath, Out)
	}
	if filex.IsExist(Out) {
		return nil, errors.New("FlagConfig:out is exist! ")
	}
	Format := *format
	if "" != Format && !formatx.CheckFormatRegistered(Format) {
		return nil, errors.New("Format Define Error: " + Format)
	}
	Ratio := *ratio

	return &FlagConfig{
		Mode: Mode, Order: Order, Size: Size,
		In: In, Out: Out, Format: Format, FormatRatio: Ratio}, nil
}
