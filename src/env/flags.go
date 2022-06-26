// Create on 2022/6/26
// @author xuzhuoxi
package env

import (
	"flag"
	"github.com/xuzhuoxi/infra-go/imagex/formatx"
	"strings"
)

func ParseFlags() *SplitContext {
	env := flag.String("env", "", "Evn Path!")
	mode := flag.String("mode", "", "Split Mode!")
	order := flag.String("order", OrderLeftUp.DefaultValue(), "Split Order!")
	flagSize := flag.String("size", "", "Flag Size!")
	trim := flag.String("trim", EndTrimOff.DefaultValue(), "Flag Size!")

	in := flag.String("in", "", "Input Path! ")
	out := flag.String("out", "", "Output Path! ")

	format := flag.String("format", string(formatx.Auto), "Formats FlagConfig!")
	ratio := flag.Int("ratio", DefaultRatio, "Ratio FlagConfig!")

	flag.Parse()

	return &SplitContext{
		EnvPath: *env, Mode: strings.ToLower(*mode), Order: strings.ToLower(*order),
		Size: strings.ToLower(*flagSize), Trim: strings.ToLower(*trim),
		InImagePath: *in, OutImagePath: *out,
		Format: *format, FormatRatio: *ratio}
}
