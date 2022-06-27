ImageSplitter_win_amd64.exe -mode=fixed -order=LeftUp -size=1024x1024 -trim=on -in=src/In.jpeg -out=tar/slice_{y1}_{x1}.{ext}
:: ImageSplitter_win_amd64.exe -mode=2 -order=2 -size=5x4 -in=src/In.jpeg -out=tar/slice_{n1}.{ext} -format=png
pause