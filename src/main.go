package main

import (
	"fmt"
	"github.com/xuzhuoxi/ImageSplitter/src/core"
	"github.com/xuzhuoxi/ImageSplitter/src/env"
	"github.com/xuzhuoxi/infra-go/filex"
	"image"
	"image/draw"
	"os"
)

func main() {
	core.InitLogger()
	ctx := env.ParseFlags()
	if err := ctx.InitContext(); nil != err {
		core.Logger.Warnln(err)
		return
	}

	img, format, err := core.LoadImage(ctx.InImagePath)
	if nil != err {
		core.Logger.Warnln(fmt.Sprintf("LoadImage Error At [%s]", err))
		return
	}

	ctx.SetImageSize(img.Bounds().Size())
	ctx.SetDefaultFormat(format)

	core.Logger.Infoln(fmt.Sprintf("开始分割图片：\"%s\"", ctx.InImagePath))
	core.Logger.Infoln(fmt.Sprintf("设置如下："))
	core.Logger.Infoln(fmt.Sprintf("环境路径=%v", ctx.EnvPath))
	core.Logger.Infoln(fmt.Sprintf("分割模式=%v", ctx.GetMode().Desc))
	core.Logger.Infoln(fmt.Sprintf("分割顺序=%v", ctx.GetOrder().Desc))
	core.Logger.Infoln(fmt.Sprintf("尾部裁剪=%v", ctx.GetTrim().Desc))
	core.Logger.Infoln(fmt.Sprintf("分割Size参数=%v", ctx.Size))

	slices := core.ParseSlice(ctx)
	sliceSize := ctx.SliceSize
	bound := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: sliceSize.Width, Y: sliceSize.Height}}
	slicesLen := len(slices)

	for index, slice := range slices {
		srcPoint, newBound := getDrawInfo(ctx, slice.SrcPoint, bound)
		newImg := image.NewRGBA(newBound)
		draw.Draw(newImg, newBound, img, srcPoint, draw.Src)
		if err := tryMakeDir(slice.FullPath); nil != err {
			core.Logger.Warnln(fmt.Sprintf("Gen image (%d/%d) fail at [%s].", index+1, slicesLen, slice.FullPath))
			return
		}
		if err := core.SaveImage(newImg, slice.FullPath, slice.Format, slice.Options); nil != err {
			core.Logger.Warnln(fmt.Sprintf("Gen image (%d/%d) fail at: \"%s\"", index+1, slicesLen, slice.FullPath))
			return
		}
		core.Logger.Infoln(fmt.Sprintf("Gen image[%v] (%d/%d) to [%s].",
			toSize(newBound.Size()), index+1, slicesLen, slice.FullPath))
	}

	countSize := ctx.CountSize
	core.Logger.Infoln(fmt.Sprintf("共生成%d张，水平方向%d张，垂直方向%d张.",
		slicesLen, countSize.Width, countSize.Height))
}

func toSize(size image.Point) string {
	return fmt.Sprintf("%dx%d", size.X, size.Y)
}

func getDrawInfo(ctx *env.SplitContext, sp image.Point, defaultBound image.Rectangle) (srcPoint image.Point, bound image.Rectangle) {
	if !ctx.TrimOn() {
		return sp, defaultBound
	}
	imageSize := ctx.ImageSize
	sliceSize := ctx.SliceSize

	startX := sp.X
	startY := sp.Y
	endX := sp.X + sliceSize.Width
	endY := sp.Y + sliceSize.Height
	if endX <= imageSize.X && endY <= imageSize.Y && startX >= 0 && startY >= 0 {
		return sp, defaultBound
	}

	if startY < 0 {
		fmt.Println("111:", startX, startY)
	}

	width := sliceSize.Width
	height := sliceSize.Height
	if endX > imageSize.X {
		width = width - endX + imageSize.X
	}
	if startY < 0 {
		height = height + startY
		startY = 0
	}
	if endY > imageSize.Y {
		height = height - endY + imageSize.Y
	}

	//return sp, defaultBound
	return image.Point{X: startX, Y: startY},
		image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: width, Y: height}}
}

func tryMakeDir(imgPath string) error {
	upDir, err := filex.GetUpDir(imgPath)
	if nil != err {
		return err
	}
	if filex.IsDir(upDir) {
		return nil
	}
	if err := os.MkdirAll(upDir, os.ModePerm); nil != err {
		return err
	}
	return nil
}
