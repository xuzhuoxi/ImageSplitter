// Create on 2022/6/26
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/ImageSplitter/src/env"
	"strconv"
	"strings"
)

// 替换路径中的通配符
func GetFullPath(outPath string, xIndex int, yIndex int, xWidth int) string {
	out := outPath
	index0 := xWidth*yIndex + xIndex
	index1 := index0 + 1
	str0 := strconv.Itoa(index0)
	str1 := strconv.Itoa(index1)
	out = replaceWildcards(out, env.WildcardN0, str0)
	out = replaceWildcards(out, env.WildcardN1, str1)

	x1 := xIndex + 1
	y1 := yIndex + 1
	strX0 := strconv.Itoa(xIndex)
	strX1 := strconv.Itoa(x1)
	strY0 := strconv.Itoa(yIndex)
	strY1 := strconv.Itoa(y1)

	out = replaceWildcards(out, env.WildcardX0, strX0)
	out = replaceWildcards(out, env.WildcardX1, strX1)

	out = replaceWildcards(out, env.WildcardY0, strY0)
	out = replaceWildcards(out, env.WildcardY1, strY1)
	return out
}

func replaceWildcards(path string, wildcards []string, rep string) string {
	for _, wildcard := range wildcards {
		path = strings.ReplaceAll(path, wildcard, rep)
	}
	return path
}
