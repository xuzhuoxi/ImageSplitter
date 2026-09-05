# ImageSplitter

ImageSplitter splits one image into multiple smaller tiles.

[中文](/README.md) | English

## <span id="p1">Compatibility

Go 1.24

## <span id="p2">How To Get Started

You can [download a release](#p2.1) or [build from source](#p2.2).

### <span id="p2.1">Download Release Version

- Download from [https://github.com/xuzhuoxi/ImageSplitter/releases](https://github.com/xuzhuoxi/ImageSplitter/releases).

### <span id="p2.2">Build

- Clone the repository

	```sh
	git clone https://github.com/xuzhuoxi/ImageSplitter.git
	```

- Preferred: run [`build/build.sh`](/build/build.sh) from the repo root (Windows: [`build/build.bat`](/build/build.bat)). It runs tests, cross-compiles binaries, and packs source into `build/release`.

- Or compile directly:

	```sh
	go build -o ImageSplitter ./src
	```

- Optional: [`build/goxc_build.sh`](/build/goxc_build.sh) / [`build/goxc_build.bat`](/build/goxc_build.bat) via [goxc](https://github.com/laher/goxc).

## <span id="p3">Run

Command-line only. A rolling log file `ImageSplitter.log` is written next to the executable.

Modes:

- **fixed**: split by a fixed tile size in pixels; leftover edges can be padded or trimmed
- **avg**: split into a given number of columns × rows

Path resolution:

- Empty `-env` uses the **directory of the executable**. If the path already exists as a directory (absolute, or relative to the current working directory), it is used as-is; otherwise it is joined with the executable directory.
- `-in`: if the path already exists, it is used as-is; otherwise it is joined with the executable directory (**not** with `-env`).
- `-out`: if its parent directory already exists, it is used as-is; otherwise it is joined with `-env`. Missing output directories are created automatically.

### <span id="p3.1">Command Line Parameter Description

- -env
  + [**Optional**] Runtime environment path, used mainly to resolve `-out`. Empty means the directory of the executable
  + Example:
    - `-env=D:/workspaces`
- -mode
  + [**Required**] Split mode, case-insensitive. `fixed` / `1`: fixed tile size; `avg` / `2`: fixed tile count
  + Example:
    - `-mode=fixed`, `-mode=1`: split by tile pixel size
    - `-mode=avg`, `-mode=2`: split into columns × rows
- -order
  + [**Optional**] Split origin, default `lu` (top-left). `lu` / `leftup` / `1` (top-left), `ld` / `leftdown` / `2` (bottom-left)
  + Example:
    - `-order=LeftUp`, `-order=lu`, `-order=1`: start from the **top-left**
    - `-order=LeftDown`, `-order=ld`, `-order=2`: start from the **bottom-left**
- -size
  + [**Required**] Split parameter, format `mxn`, `mXn`, or `m*n` (case-insensitive). In `fixed` mode this is tile width × height in pixels; in `avg` mode it is column count × row count
  + Example:
    - `-mode=fixed -size=512*512`: tiles of 512×512 pixels
    - `-mode=avg -size=10*10`: a 10×10 grid of tiles
- -trim
  + [**Optional**] Edge trimming, default `off`. `on` / `true` / `enable` / `1`, `off` / `false` / `disable` / `0`
  + `on`: tiles that would extend past the source image are cropped to the remaining pixels
  + `off`: edge tiles keep the full tile size, with empty padding
  + Example:
    - `-trim=on`, `-trim=true`, `-trim=1`: enable trimming
    - `-trim=off`, `-trim=false`, `-trim=0`: disable trimming
- -format
  + [**Optional**] Output format: `png` / `jpg` (`jpeg`). Unspecified or `auto` uses the source image format
  + Example:
    - `-format=png`
- -ratio
  + [**Optional**] JPEG quality, integer, default **80**. Ignored for PNG
  + Example:
    - `-ratio=80`: JPEG quality 80
- -in
  + [**Required**] Source image file
  + Absolute path, or a relative path against the current working directory / executable directory
  + Example:
    - `-in=D:/workspaces/res/In.png`
    - `-in=res/In.png`
- -out
  + [**Required**] Output path; must include **wildcards** so each tile gets a unique file name
  + Absolute path, or a relative path resolved against `-env`
  + Wildcards:
    - `{n0}` / `{N0}`: 0-based linear index
    - `{n1}` / `{N1}`: 1-based linear index
    - `{x0}` / `{X0}`: 0-based **column** index
    - `{x1}` / `{X1}`: 1-based **column** index
    - `{y0}` / `{Y0}`: 0-based **row** index
    - `{y1}` / `{Y1}`: 1-based **row** index
    - `{ext}`: extension of the final output format
  + Example:
    - `-env=D:/workspaces -out=dir/Slice{n1}.png`
    - `-env=D:/workspaces -out=dir/Slice{y1}_{x1}.{ext}`

### <span id="p3.2">Usage Examples

**Note:** relative paths follow the resolution rules above.

- Split into 1024×1024 tiles, trim edges, name files by row and column:

	```sh
	ImageSplitter -mode=fixed -order=LeftUp -size=1024x1024 -trim=on -in=src/In.jpeg -out=tar/slice_{y1}_{x1}.{ext}
	```

- Split into a 5×4 grid, write PNG, name files with a 1-based index:

	```sh
	ImageSplitter -mode=avg -order=2 -size=5x4 -in=src/In.jpeg -out=tar/slice_{n1}.{ext} -format=png
	```

### <span id="p3.3">Example

- Samples live in [demo](/demo).
- On Windows, run [DemoRunWin.bat](/demo/DemoRunWin.bat).
- On macOS, see [DemoRunMac.sh](/demo/DemoRunMac.sh); Linux can use the same script. Adjust the executable path in the script to match your local binary name before running.

  [Command line parameter description](#p3.1)

## <span id="p4">Dependency Library

- infra-go [https://github.com/xuzhuoxi/infra-go](https://github.com/xuzhuoxi/infra-go)
- goxc [https://github.com/laher/goxc](https://github.com/laher/goxc) (optional, legacy build scripts)

## <span id="p5">Contact

xuzhuoxi

<xuzhuoxi@gmail.com> or <mailxuzhuoxi@163.com>

## <span id="p6">License

ImageSplitter source code is available under the MIT [License](/LICENSE).
