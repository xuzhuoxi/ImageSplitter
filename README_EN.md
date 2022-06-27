# ImageResizer

ImageSplitter is mainly used to split images into multiple small images.

[中文](/README.md) | English

**Note：** This document is provided with translations based on [Google Translate](https://translate.google.cn/).

## <span id="p1">Compatibility
go1.16

## <span id="p2">How To Get Started

You can choose [Download Release Version](#p2.1) or [Construction](#p2.2) to get the executable file.

### <span id="p2.1">Download Release Version

- Download from the following address: [https://github.com/xuzhuoxi/ImageSplitter/releases](https://github.com/xuzhuoxi/ImageSplitter/releases).

### <span id="p2.2">Build

- Download repository

	```sh
	go get -u github.com/xuzhuoxi/ImageSplitter
	```

- Build

  + The construction depends on the third-party library [goxc](https://github.com/laher/goxc).

  + If necessary, you can modify the relevant construction scripts.

  + It is recommended to turn off gomod first: `go env -w GO111MODULE=off`, because goxc is old.

  + Execute the build script [goxc_build.sh](/build/goxc_build.sh) or [goxc_build.bat](/build/goxc_build.bat), the executable file will be generated in the "build/release" directory .

## <span id="p3">Run

The tool supports command line execution only.

### <span id="p3.1">Command Line Parameter Description

  - -env
    + [**Optional**] Runtime environment path, supports absolute path and relative path relative to the current execution directory, empty means to use the directory where the execution file is located
    + Example:
      `-env=D:/workspaces`
  - -mode
    + [**Required**] Split mode, support: fixed(1), avg(2)
    + Example:
      `-mode=fixed` and `-mode=1` both indicate the use of fixed mode, fixed division, and fill in the gaps
      `-mode=avg` , `-mode=2` both indicate the use of avg mode, average division, and average horizontal and vertical division according to the total size of the picture
  - -order
    + [**Required**] Split order, support: LeftUp(1), LeftDown(2)
    + Example:
      `-order=LeftUp` , `-order=lu` , `-mode=1` all indicate that the starting point is **upper left**.
      `-order=LeftDown` or `-mode=ld` or `-mode=2` all means start from **lower left**.
  - -size
    + [**Required**] Split parameter, format "mxn","mXn","m*n"
    + Example:
      `-modex=fixed -size=512*512` means to divide the image according to the small size of 512 in length and 512 in width
      `-modex=avg -size=10*10` means to divide the image into 10 times 10 small images
  - -trim
    + [**Required**] Tail cropping, support: on, off
    + Example:
      `-trim=on` , `-trim=true` , `-trim=1` means enable tail trimming
      `-trim=off` , `-trim=false` , `-trim=0` means to turn off tail clipping
  - -format
    + 【**Not necessary**】Forcibly specify the image file format, if not specified, the format of the source image will be used
    + Example:
      `-format=png` indicates that the output thumbnail file uses png
  - -ratio
    + [**non-essential**] Force to specify the image file quality (if necessary), if not specified, use the default value of 85
    + Example:
      `-ratio=60` means that the output small image quality is 85.
      **Note**: Image formats that do not require quality parameters will be ignored here, such as png format
  - -in
    + [**Required**] Source image path, which requires a file in image format
    + Absolute paths can be used.
    + You can use a relative path, which will be used with the -env parameter value or the current execution file directory.
    + Example:
      `-env=D:/workspaces -in=res/In.png` means to use the image `D:/workspaces/res/In.png`.
  - -out
    + [**Required**] To output image information, it is required to use the **wildcard** directory,
    + Absolute paths can be used.
    + You can use a relative path, which will be used with the -env parameter value or the current execution file directory.
    + wildcard description:
      - "{n0}", "{N0}": Indicates that it is replaced with a division sequence number starting from 0.
      - "{n1}", "{N1}": Indicates to use 1-based division sequence number to replace.
      - "{x1}", "{X1}": Indicates that the **horizontal** division sequence number starting from 0 is used to replace.
      - "{n1}", "{N1}": Indicates that the **horizontal** division sequence number starting from 1 is used to replace.
      - "{y0}", "{Y0}": Indicates that the **vertical direction** division sequence number starting from 0 is used to replace.
      - "{y1}", "{Y1}": Indicates that the **vertical direction** division sequence number starting from 1 is used to replace.
      - "{ext}": Indicates the extension corresponding to the format of the auto-fill generated image.
    + Example:
      `-env=D:/workspaces -out=dir/Slice{n1}.png`
      `-env=D:/workspaces -out=dir/Slice{y1}_{x1}.{ext}`

### <span id="p3.3">Example

- The example directory is located at [demo](/demo).

- The Win64 platform can execute [DemoRun.bat.bat](/demo/DemoRun.bat.bat) for testing.

- The Mac platform can execute [DemoRun.bat.sh](/demo/DemoRun.bat.sh) for testing.

- Modify the execution file path in the Mac test script for testing on the Linux platform.

  [Command line parameter description](#p3.1)

## <span id="p4">Dependency Library

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## <span id="p5">Contact

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## <span id="p6">License

ImageSplitter source code is available under the MIT [License](/LICENSE).