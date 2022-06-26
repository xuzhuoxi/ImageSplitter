#!/usr/bin/env bash

ImageSplitter -mode=1 -order=1 -size=1024x1024 -in=./src.jpeg -out=./out/a_{y1}_{x1}.png -format=png

ImageSplitter -mode=2 -order=1 -size=3x4 -in=./src.jpg -out=./out/a_{y0}_{x0}.jpg -ratio=90