#!/usr/bin/env bash
CRTDIR=$(cd `dirname $0`; pwd)

${CRTDIR}/ImageSplitter -mode=2 -order=0 -size=3x4 -in=./1NdxvUJ.jpg -out=./out/a_{y1}_{x1}.jpg -format=png