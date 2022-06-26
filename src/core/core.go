// Create on 2022/6/26
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/ImageSplitter/src/env"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	"image"
)

type ImageSlice struct {
	FullPath string
	SrcPoint image.Point
	Format   formatx.ImageFormat
	Options  interface{}
}

func ParseSlice(ctx *env.SplitContext) (slices []*ImageSlice) {
	countSize := ctx.CountSize
	slices = make([]*ImageSlice, countSize.Width*countSize.Height)

	format := formatx.ImageFormat(ctx.Format)
	options := GetOptions(format, ctx.FormatRatio)

	for yIndex := 0; yIndex < countSize.Height; yIndex += 1 {
		for xIndex := 0; xIndex < countSize.Width; xIndex += 1 {
			slice := &ImageSlice{Format: format, Options: options}
			slice.SrcPoint = ctx.GetSrcPoint(xIndex, yIndex)
			slice.FullPath = GetFullPath(ctx.OutImagePath, xIndex, yIndex, countSize.Width)

			index := yIndex*countSize.Width + xIndex
			slices[index] = slice
		}
	}
	return
}

func GetOptions(format formatx.ImageFormat, ratio int) interface{} {
	if formatx.PNG == format {
		return nil
	}
	if formatx.Auto == format {
		return nil
	}
	return jpegx.NewJpegOptions(ratio)
}
