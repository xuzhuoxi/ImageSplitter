// Create on 2022/6/26
// @author xuzhuoxi
package core

import (
	"github.com/xuzhuoxi/infra-go/logx"
	"github.com/xuzhuoxi/infra-go/mathx"
	"github.com/xuzhuoxi/infra-go/osxu"
)

var (
	Logger      logx.ILogger
	LogFileName = "ImageSplitter"
)

func InitLogger() {
	Logger = logx.NewLogger()
	Logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	Logger.SetConfig(logx.LogConfig{Type: logx.TypeRollingFile, Level: logx.LevelAll,
		FileDir: osxu.GetRunningDir(), FileName: LogFileName, FileExtName: ".log", MaxSize: 10 * mathx.MB})
}
