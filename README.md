# ImageSplitter

ImageSplitter can be used to split image into multiple small images. You can specify the size or the number of divisions.

## Compatibility

go 1.15.8

## Getting Started

### Download Release

- Download the release [here](https://github.com/xuzhuoxi/ImageSplitter/releases).

- Download the repository:

	```sh
	go get -u github.com/xuzhuoxi/ImageSplitter
	```
	
	This will retrieve the library.

### Build

Execution the construction file([build.sh](/build/build.sh)) to get the releases if you have already downloaded the repository.

You can modify the construction file([build.sh](/build/build.sh)) to achieve what you want if necessary. The command line description is [here](https://github.com/laher/goxc).

## Run

### Demo

[Here](/demo/mac) is a running demo for MacOS platform.

The running command is consistent of all platforms.

Goto <a href="#command-line">Command Line Description</a>.

### Command Line

Supportted command line parameters as follow:

| -       | -            | -                                                            |
| :------ | :----------- | ------------------------------------------------------------ |
| -mode   | optional | The mode of the divisions.  1：小图使用固定尺寸；	2：小图使用平均尺寸|
| -order  | optional | The order of the divisions. 1：左上角为起始点；	2：左下角为起始点|
| -size   | **required**     | The size info of divisions. 格式：mxn。当mode为1时，m、n代表小图尺寸；当mode为2时，m、n代表分割数量|
| -in     | **required**     | Custom source file. |
| -out    | **required**     | Custom output files. 支持通配符（{n0},{N0},{n1},{N1},{x0},{X0},{x1},{X1},{y0},{Y0},{y1},{Y1}）|
| -format | optional     | The format of the generated image. Supported as follows: png, jpg, jpeg, jps |
| -ratio  | optional     | The quality of the generated image. Supported for jpg,jpeg,jps. |

E.g.:

-mode=1

-mode=2

-order=1

-order=2

-size=256x256

-in=./source/image.png

-in=/Users/aaa/image.jpg

-out=/Users/aaa/image_{n0}_{y1}_{x1}.png

-out=./out/image_{n0}_{y1}_{x1}.png

-format=jpeg

-format=jpg

-format=png

-ratio=85

## Related Library

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)

- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) 

## Contact

xuzhuoxi 

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## License

ImageSplitter source code is available under the MIT [License](/LICENSE).