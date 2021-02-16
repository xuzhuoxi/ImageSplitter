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

// -mode 		必选	自定义基目录				字符串路径，文件夹或文件,"./"开头视为相对路径
// -order 		必选	自定义基目录				字符串路径，文件夹或文件,"./"开头视为相对路径
// -size 		必选	输出大小					[整数/宽x高],...
// -in 			可选	来源地址					字符串路径，文件夹或文件,"./"开头视为相对路径
// -out 		可选	输出地址					字符串路径，文件夹,"./"开头视为相对路径
// -format 		可选	输出文件格式				图像格式[pngx,jpeg,gifx,jpg]
// -ratio 		可选	压缩比					整数(0,100]
func ParseFlag() (cfg *FlagConfig, err error) {
	mode := flag.Int("mode", int(SizeMode), "SplitMode!")
	order := flag.Int("order", int(LeftUp), "SplitOrder!")
	size := flag.String("size", "", "Size FlagConfig!")

	in := flag.String("in", "", "Input Path! ")
	out := flag.String("out", "", "Output Path! ")

	format := flag.String("format", string(formatx.Auto), "Format FlagConfig!")
	ratio := flag.Int("ratio", 75, "Ratio FlagConfig!")

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
