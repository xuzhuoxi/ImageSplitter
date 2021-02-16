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
	slices, sliceSize, _ := lib.ParseSlice(flagConfig, imageSize, imgFormat)
	//fmt.Print("看看：", len(slices), sliceSize, countSize, imgFormat)
	bound := image.Rectangle{Min: image.Point{X: 0, Y: 0}, Max: image.Point{X: sliceSize.Width, Y: sliceSize.Height}}
	for _, slice := range slices {
		newImg := image.NewRGBA(bound)
		draw.Draw(newImg, bound, img, slice.SrcPoint, draw.Src)
		imagex.SaveImage(newImg, slice.FullPath, slice.Format, slice.Options)
	}
}
