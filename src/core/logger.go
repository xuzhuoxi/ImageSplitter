// Create on 2022/6/26
// @author xuzhuoxi
package core

import "github.com/xuzhuoxi/infra-go/logx"

var (
	Logger logx.ILogger
)

func InitLogger() {
	Logger = logx.NewLogger()
	Logger.SetConfig(logx.LogConfig{Type: logx.TypeConsole, Level: logx.LevelAll})
	Logger.SetConfig(logx.LogConfig{Type: logx.TypeRollingFile, Level: logx.LevelAll})
}
