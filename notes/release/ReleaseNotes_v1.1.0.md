## Release Notes

+ 本版本将构建要求提升到 Go 1.24.0，适配 Go module 交叉编译，并补齐本地构建脚本与 GitHub CI / 发版工作流。

### Known Issues

+ 中英文 README 仍写兼容性 go1.16，构建说明仍以 goxc 为主，尚未改为 Go 1.24 与 `build/build.sh` / `build/build.bat`。

### Improvements

+ 新增 `build/build.sh` / `build/build.bat`：单元测试、多平台交叉编译可执行文件、打包源码。
+ 新增 GitHub Workflows：`CI.yml`（构建与测试）、`Release.yml`（打 tag 发版）、`ReleaseNote.yml`（更新已有 Release 正文）。
+ 添加 Cursor Skill `generate-note`，用于根据提交范围生成 Release 说明。
+ `.gitignore` 增加 `go.work`、`go.work.sum`。

### Changes

+ 构建所需 Go 版本提升到 1.24.0。
+ 构建改为 Go module 模式，新增 `go.mod` / `go.sum`。

## Library Changes

+ `github.com/xuzhuoxi/infra-go`：新增 v1.4.1
