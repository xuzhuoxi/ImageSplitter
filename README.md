# ImageSplitter

ImageSplitter 主要用于分割图像成为多张小图像。

中文 | [English](/README_EN.md)

## <span id="p1">兼容性
go1.16

## <span id="p2">如何开始

你可以选择[下载发行版本](#p2.1)或者[构造](#p2.2)获得执行文件。

### <span id="p2.1">下载发行版本

- 到以下地址下载: [https://github.com/xuzhuoxi/- 

### <span id="p2.2">构造

- 下载仓库

	```sh
	go get -u github.com/xuzhuoxi/ImageResizer
	```

- 构造

  + 构造依赖到第三方库[goxc](https://github.com/laher/goxc)。

  + 如有必要，你可以修改相关构造脚本。

  + 建议先关闭gomod：`go env -w GO111MODULE=off`，由于goxc已经比较旧。

  + 执行构造脚本[goxc_build.sh](/build/goxc_build.sh)或[goxc_build.bat](/build/goxc_build.bat),执行文件将生成在"build/release"目录中。

## <span id="p3">运行

工具仅支持命令行执行。

### <span id="p3.1">命令行参数说明

- -env 
  + 【**可选**】运行时环境路径，支持绝对路径与相对于当前执行目录的相对路径，空表示使用执行文件所在目录
  + 例子: 
    `-env=D:/workspaces`
- -mode
  + 【**必要**】分割模式，支持：fixed(1)、avg(2)
  + 例子: 
    `-mode=fixed` 、 `-mode=1` 都表示使用fixed模式, 固定分割，不足的补空
    `-mode=avg` 、 `-mode=2` 都表示使用avg模式, 平均分割，根据图片总大小进行水平与垂直的平均分割
- -order
  + 【**必要**】分割顺序，支持：LeftUp(1)、LeftDown(2)
  + 例子: 
    `-order=LeftUp` 、 `-order=lu` 、 `-mode=1` 都表示以**左上**为起点。
    `-order=LeftDown` 或 `-mode=ld` 或 `-mode=2` 都表示以**左下**为起点。
- -size
  + 【**必要**】分割参数，格式 "mxn","mXn","m*n"
  + 例子: 
    `-modex=fixed -size=512*512` 表示按长为512宽为512的小图尺寸分割图像
    `-modex=avg -size=10*10` 表示把图像分割为10乘以10张小图
- -trim
  + 【**必要**】尾部裁剪，支持：on、off
  + 例子: 
    `-trim=on` 、 `-trim=true`  、`-trim=1` 表示启用尾部裁剪
    `-trim=off` 、 `-trim=false`  、`-trim=0` 表示关闭尾部裁剪
- -format
  + 【**非必要**】强制指定图像文件格式, 如果不指定，将使用源图像的格式
  + 例子: 
    `-format=png` 表示输出小图文件使用png
- -ratio
  + 【**非必要**】强制指定图像文件质量(如有必要)，如果不指定，使用默认值85
  + 例子: 
    `-ratio=60` 表示输出小图文质量为85。
    **注意**： 无须使用质量参数的图像格式将会忽略此处，如png格式 
- -in
  + 【**必要**】来源图片路径,要求为图片格式的文件
  + 可以使用绝对路径。
  + 可以使用相对路径，将配合-env参数值或当前执行文件目录使用。
  + 例子: 
    `-env=D:/workspaces -in=res/In.png` 表示使用`D:/workspaces/res/In.png`这个图像。
- -out
  + 【**必要**】输出图片信息，要求使用**通配符**目录,
  + 可以使用绝对路径。
  + 可以使用相对路径，将配合-env参数值或当前执行文件目录使用。
  + 通配符说明：
   - "{n0}", "{N0}": 表示使用 从0开始的分割顺序数 替换。
   - "{n1}", "{N1}": 表示使用 从1开始的分割顺序数 替换。
   - "{x1}", "{X1}": 表示使用 从0开始的**水平方向**分割顺序数 替换。
   - "{n1}", "{N1}": 表示使用 从1开始的**水平方向**分割顺序数 替换。
   - "{y0}", "{Y0}": 表示使用 从0开始的**垂直方向**分割顺序数 替换。
   - "{y1}", "{Y1}": 表示使用 从1开始的**垂直方向**分割顺序数 替换。
   - "{ext}": 表示自动填充生成图像的格式对应的扩展名。
  + 例子: 
    `-env=D:/workspaces -out=dir/Slice{n1}.png` 
    `-env=D:/workspaces -out=dir/Slice{y1}_{x1}.{ext}` 

### <span id="p3.3">例子

- 例子目录位于[demo](/demo).

- Win64平台可执行[DemoRun.bat.bat](/demo/DemoRun.bat.bat)进行测试。

- Mac平台可执行[DemoRun.bat.sh](/demo/DemoRun.bat.sh)进行测试。

- Linux平台修改Mac测试脚本中的执行文件路径进行测试。

  [命令行参数说明](#p3.1)

## <span id="p4">依赖库

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## <span id="p5">联系作者

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## <span id="p6">License

ImageSplitter source code is available under the MIT [License](/LICENSE).


