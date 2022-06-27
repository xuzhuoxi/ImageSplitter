ImageSplitter.exe -mode=fixed -order=LeftUp -size=1024x1024 -in=src/In.jpeg -out=tar/slice_{y1}_{x1}.{ext}
:: ImageSplitter -mode=2 -order=2 -size=5x4 -in=src/In.jpeg -out=tar/slice_{n1}.{ext} -format=png
pause