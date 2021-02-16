package main

import (
	"fmt"
	"github.com/xuzhuoxi/ImageSplitter/src/lib"
	"github.com/xuzhuoxi/infra-go/imagex"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"github.com/xuzhuoxi/infra-go/logx"
	"image"
	"image/draw"
)

var (
	globalLogger logx.ILogger
)

func main() {
	globalLogger = logx.NewLogger()
	globalLogger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})

	flagConfig, err := lib.ParseFlag()
	if err != nil {
		globalLogger.Error(err)
		return
	}

	img, imgFormat, err := imagex.LoadImage(flagConfig.In, "")
	if nil != err {
		fmt.Println(err)
		return
	}

	imageSize := img.Bounds().Size()
	slices, sliceSize, countSize := lib.ParseSlice(flagConfig, imageSize, imgFormat)
	bound := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: sliceSize.Width, Y: sliceSize.Height}}
	slicesLen := len(slices)

	globalLogger.Infoln(fmt.Sprintf("开始分割图片：\"%s\"", flagConfig.In))
	globalLogger.Infoln(fmt.Sprintf("设置如下："))
	globalLogger.Infoln(fmt.Sprintf("分割模式=%v", flagConfig.Mode))
	globalLogger.Infoln(fmt.Sprintf("分割顺序=%v", flagConfig.Order))
	globalLogger.Infoln(fmt.Sprintf("分割Size=%v", flagConfig.Size))
	globalLogger.Infoln(fmt.Sprintf("共生成%d张，水平方向%d张，垂直方向%d张，每张尺寸为%dx%d", slicesLen, countSize.Width, countSize.Height, sliceSize.Width, sliceSize.Height))

	for index, slice := range slices {
		newImg := image.NewRGBA(bound)
		draw.Draw(newImg, bound, img, slice.SrcPoint, draw.Src)
		imagex.SaveImage(newImg, slice.FullPath, slice.Format, slice.Options)
		globalLogger.Infoln(fmt.Sprintf("Gen image (%d/%d) at: \"%s\"", index+1, slicesLen, slice.FullPath))
	}
}
