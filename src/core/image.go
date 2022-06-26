// Create on 2022/6/26
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/infra-go/imagex"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/jpegx"
	_ "github.com/xuzhuoxi/infra-go/imagex/formatx/pngx"
	"image"
)

func LoadImage(imagePath string) (img image.Image, imgFormat string, err error) {
	return imagex.LoadImage(imagePath, "")
}

func SaveImage(img image.Image, fullPath string, format formatx.ImageFormat, options interface{}) error {
	return imagex.SaveImage(img, fullPath, format, options)
}
