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
	countSize := ctx.CountSize
	bound := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: sliceSize.Width, Y: sliceSize.Height}}
	slicesLen := len(slices)

	for index, slice := range slices {
		newImg := image.NewRGBA(bound)
		draw.Draw(newImg, bound, img, slice.SrcPoint, draw.Src)
		if err := tryMakeDir(slice.FullPath); nil != err {
			core.Logger.Warnln(fmt.Sprintf("Gen image (%d/%d) fail at: \"%s\"", index+1, slicesLen, slice.FullPath))
			return
		}
		if err := core.SaveImage(newImg, slice.FullPath, slice.Format, slice.Options); nil != err {
			core.Logger.Warnln(fmt.Sprintf("Gen image (%d/%d) fail at: \"%s\"", index+1, slicesLen, slice.FullPath))
			return
		}
		core.Logger.Infoln(fmt.Sprintf("Gen image (%d/%d) at: \"%s\"", index+1, slicesLen, slice.FullPath))
	}
	core.Logger.Infoln(fmt.Sprintf("共生成%d张，水平方向%d张，垂直方向%d张，每张尺寸为%dx%d",
		slicesLen, countSize.Width, countSize.Height, sliceSize.Width, sliceSize.Height))
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
