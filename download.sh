#!/bin/bash

# 先安装工具
go install github.com/hailaz/doc2pdf/cmd/doc2pdf@latest
# 运行下载
doc2pdf gf -m=md
# 复制生成的文件到docs和static
cp -r -f output/goframe-latest-md/* docs/
cp -r -f output/goframe-latest-md-static/* static/
rm docs/首页.md
