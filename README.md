# ImageSplitter

ImageSplitter 主要用于把一张图像分割成多张小图。

中文 | [English](/README_EN.md)

## <span id="p1">兼容性

Go 1.24

## <span id="p2">如何开始

你可以选择[下载发行版本](#p2.1)或者[构造](#p2.2)获得执行文件。

### <span id="p2.1">下载发行版本

- 到以下地址下载: [https://github.com/xuzhuoxi/ImageSplitter/releases](https://github.com/xuzhuoxi/ImageSplitter/releases).

### <span id="p2.2">构造

- 克隆仓库

	```sh
	git clone https://github.com/xuzhuoxi/ImageSplitter.git
	```

- 推荐在仓库根目录执行 [`build/build.sh`](/build/build.sh)（Windows 可用 [`build/build.bat`](/build/build.bat)）：会跑单元测试、交叉编译可执行文件并打包源码，产物在 `build/release`。

- 也可以直接编译：

	```sh
	go build -o ImageSplitter ./src
	```

- 可选：使用 [`build/goxc_build.sh`](/build/goxc_build.sh) / [`build/goxc_build.bat`](/build/goxc_build.bat) 通过 [goxc](https://github.com/laher/goxc) 构建。

## <span id="p3">运行

工具仅支持命令行执行。运行时会在可执行文件所在目录写入 `ImageSplitter.log`。

当前支持两种分割模式：

- **fixed**：按固定小图像素尺寸切割；边缘不足的部分可补空或裁剪
- **avg**：按指定的水平 / 垂直张数平均切割

相对路径的解析规则：

- `-env` 为空时，使用**可执行文件所在目录**。路径本身已是存在的目录则直接使用（绝对路径，或相对当前工作目录）；否则拼到可执行文件目录下。
- `-in`：路径本身已存在则直接使用；否则拼到可执行文件目录下（**不**与 `-env` 拼接）。
- `-out`：其父目录已存在则直接使用；否则拼到 `-env` 下。产出目录不存在时会自动创建。

### <span id="p3.1">命令行参数说明

- -env
  + 【**可选**】运行时环境路径，主要用于解析 `-out`。空表示使用可执行文件所在目录
  + 例子:
    - `-env=D:/workspaces`
- -mode
  + 【**必要**】分割模式，大小写不敏感。`fixed` / `1`：固定尺寸；`avg` / `2`：固定数量
  + 例子:
    - `-mode=fixed`、`-mode=1`：按小图像素尺寸切割
    - `-mode=avg`、`-mode=2`：按水平 × 垂直张数平均切割
- -order
  + 【**可选**】分割顺序，默认 `lu`（左上）。支持：`lu` / `leftup` / `1`（左上），`ld` / `leftdown` / `2`（左下）
  + 例子:
    - `-order=LeftUp`、`-order=lu`、`-order=1`：以**左上**为起点
    - `-order=LeftDown`、`-order=ld`、`-order=2`：以**左下**为起点
- -size
  + 【**必要**】分割参数，格式 `mxn`、`mXn`、`m*n`（大小写不敏感）。`fixed` 表示小图宽 × 高（像素）；`avg` 表示水平张数 × 垂直张数
  + 例子:
    - `-mode=fixed -size=512*512`：按 512×512 像素切割
    - `-mode=avg -size=10*10`：切成 10×10 张小图
- -trim
  + 【**可选**】边缘裁剪，默认 `off`。支持：`on` / `true` / `enable` / `1`，`off` / `false` / `disable` / `0`
  + `on`：超出原图范围的边缘小图按实际剩余像素裁剪
  + `off`：边缘小图仍按完整尺寸输出，不足部分补空
  + 例子:
    - `-trim=on`、`-trim=true`、`-trim=1`：启用裁剪
    - `-trim=off`、`-trim=false`、`-trim=0`：关闭裁剪
- -format
  + 【**可选**】强制指定输出格式，目前为 `png` / `jpg`（`jpeg`）。未指定或为 `auto` 时使用源图像格式
  + 例子:
    - `-format=png`：输出 png
- -ratio
  + 【**可选**】JPEG 品质，整数，默认 **80**。PNG 忽略该参数
  + 例子:
    - `-ratio=80`：JPEG 品质为 80
- -in
  + 【**必要**】来源图片路径，须为图像文件
  + 可以使用绝对路径，或相对当前工作目录 / 可执行文件目录的相对路径
  + 例子:
    - `-in=D:/workspaces/res/In.png`
    - `-in=res/In.png`
- -out
  + 【**必要**】输出路径，须包含**通配符**以便为每张小图生成文件名
  + 可以使用绝对路径；相对路径将配合 `-env` 使用
  + 通配符说明：
    - `{n0}` / `{N0}`：从 0 开始的分割序号
    - `{n1}` / `{N1}`：从 1 开始的分割序号
    - `{x0}` / `{X0}`：从 0 开始的**水平**序号
    - `{x1}` / `{X1}`：从 1 开始的**水平**序号
    - `{y0}` / `{Y0}`：从 0 开始的**垂直**序号
    - `{y1}` / `{Y1}`：从 1 开始的**垂直**序号
    - `{ext}`：按最终输出格式填充扩展名
  + 例子:
    - `-env=D:/workspaces -out=dir/Slice{n1}.png`
    - `-env=D:/workspaces -out=dir/Slice{y1}_{x1}.{ext}`

### <span id="p3.2">应用场景举例

**注意**：以下相对路径均相对于运行时的工作目录或 `-env`（见上方解析规则）。

- 按 1024×1024 切割，启用边缘裁剪，文件名带行列号：

	```sh
	ImageSplitter -mode=fixed -order=LeftUp -size=1024x1024 -trim=on -in=src/In.jpeg -out=tar/slice_{y1}_{x1}.{ext}
	```

- 平均切成 5×4 张，输出 png，按从 1 开始的序号命名：

	```sh
	ImageSplitter -mode=avg -order=2 -size=5x4 -in=src/In.jpeg -out=tar/slice_{n1}.{ext} -format=png
	```

### <span id="p3.3">例子

- 例子目录位于 [demo](/demo)。
- Windows 可执行 [DemoRunWin.bat](/demo/DemoRunWin.bat)。
- macOS 可参考 [DemoRunMac.sh](/demo/DemoRunMac.sh)；Linux 可同样参考该脚本。请按本机可执行文件名修改脚本中的路径后再运行。

  [命令行参数说明](#p3.1)

## <span id="p4">依赖库

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)
- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc)（可选，旧构建脚本）

## <span id="p5">联系作者

xuzhuoxi

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## <span id="p6">License

ImageSplitter source code is available under the MIT [License](/LICENSE).
