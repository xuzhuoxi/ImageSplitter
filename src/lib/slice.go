package lib

import (
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"image"
	"image/jpeg"
	"math"
	"strconv"
	"strings"
)

type ImageSlice struct {
	FullPath string
	SrcPoint image.Point
	Format   formatx.ImageFormat
	Options  interface{}
}

func ParseSlice(flagConfig *FlagConfig, imageSize image.Point, imageFormat string) (slices []*ImageSlice, sliceSize SplitSize, countSize SplitSize) {
	sliceSize, countSize = parseSize(flagConfig.Mode, flagConfig.Size, imageSize)
	slices = make([]*ImageSlice, countSize.Width*countSize.Height)

	format := getFormat(flagConfig.Format, imageFormat)
	options := getOptions(format, flagConfig.FormatRatio)

	for yIndex := 0; yIndex < countSize.Height; yIndex += 1 {
		for xIndex := 0; xIndex < countSize.Width; xIndex += 1 {
			slice := &ImageSlice{Format: format, Options: options}
			slice.SrcPoint = getSrcPoint(flagConfig.Order, xIndex, yIndex, sliceSize, imageSize)
			slice.FullPath = getFullPath(flagConfig.Out, xIndex, yIndex, countSize.Width)

			index := yIndex*countSize.Width + xIndex
			slices[index] = slice
		}
	}
	return
}

func getFormat(flagFormat string, srcFormat string) formatx.ImageFormat {
	if flagFormat == string(formatx.Auto) {
		return formatx.ImageFormat(srcFormat)
	}
	return formatx.ImageFormat(flagFormat)
}

func getOptions(format formatx.ImageFormat, formatRatio int) interface{} {
	if formatx.PNG == format {
		return nil
	}
	return &jpeg.Options{Quality: formatRatio}
}

func getFullPath(outPath string, xIndex int, yIndex int, xWidth int) string {
	out := outPath

	index0 := xWidth*yIndex + xIndex
	index1 := index0 + 1
	str0 := strconv.Itoa(index0)
	str1 := strconv.Itoa(index1)
	out = strings.ReplaceAll(out, N0Low, str0)
	out = strings.ReplaceAll(out, N0Up, str0)
	out = strings.ReplaceAll(out, N1Low, str1)
	out = strings.ReplaceAll(out, N1Up, str1)

	x1 := xIndex + 1
	y1 := yIndex + 1
	strX0 := strconv.Itoa(xIndex)
	strX1 := strconv.Itoa(x1)
	strY0 := strconv.Itoa(yIndex)
	strY1 := strconv.Itoa(y1)

	out = strings.ReplaceAll(out, X0Low, strX0)
	out = strings.ReplaceAll(out, X0Up, strX0)
	out = strings.ReplaceAll(out, X1Low, strX1)
	out = strings.ReplaceAll(out, X1Up, strX1)

	out = strings.ReplaceAll(out, Y0Low, strY0)
	out = strings.ReplaceAll(out, Y0Up, strY0)
	out = strings.ReplaceAll(out, Y1Low, strY1)
	out = strings.ReplaceAll(out, Y1Up, strY1)

	return out
}

func checkStringAnd(str string, check []string) bool {
	for _, c := range check {
		if !checkSubString(str, c) {
			return false
		}
	}
	return true
}

func checkStringOr(str string, check []string) bool {
	for _, c := range check {
		if checkSubString(str, c) {
			return true
		}
	}
	return false
}

func checkSubString(str, check string) bool {
	return strings.Index(str, check) >= 0
}

func getSrcPoint(order SplitOrder, xIndex int, yIndex int, sliceSize SplitSize, imageSize image.Point) (srcPoint image.Point) {
	x, y := 0, 0
	switch order {
	case LeftUp:
		x = xIndex * sliceSize.Width
		y = yIndex * sliceSize.Height
	case LeftDown:
		x = xIndex * sliceSize.Width
		y = imageSize.Y - (yIndex+1)*sliceSize.Height
	}
	srcPoint = image.Point{X: x, Y: y}
	return
}

func parseSize(mode SplitMode, flagSize SplitSize, imageSize image.Point) (sliceSize SplitSize, countSize SplitSize) {
	switch mode {
	case SizeMode:
		sliceSize, countSize = parseSizeBySizeMode(flagSize, imageSize)
	case AvgMode:
		sliceSize, countSize = parseSizeByAvgMode(flagSize, imageSize)
	}
	return
}

func parseSizeBySizeMode(flagSize SplitSize, imageSize image.Point) (sliceSize SplitSize, countSize SplitSize) {
	sliceSize = flagSize
	w := math.Ceil(float64(imageSize.X) / float64(flagSize.Width))
	h := math.Ceil(float64(imageSize.Y) / float64(flagSize.Height))
	countSize = SplitSize{Width: int(w), Height: int(h)}
	return
}

func parseSizeByAvgMode(flagSize SplitSize, imageSize image.Point) (sliceSize SplitSize, countSize SplitSize) {
	countSize = flagSize
	w := math.Ceil(float64(imageSize.X) / float64(flagSize.Width))
	h := math.Ceil(float64(imageSize.Y) / float64(flagSize.Height))
	sliceSize = SplitSize{Width: int(w), Height: int(h)}
	return
}
