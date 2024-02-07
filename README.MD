# gfdoc-md

https://www.hailaz.cn/gfdoc-md/

<https://docusaurus.io/zh-CN/docs>

## docs 文档下载方式

```shell
// 先安装工具
go install github.com/hailaz/doc2pdf/cmd/doc2pdf@latest
// 运行下载
doc2pdf gf -m=md
// 复制生成的文件到docs和static
cp -r -f output/goframe-latest-md/* docs/
cp -r -f output/goframe-latest-md-static/* static/

// windows
// xcopy /s /y A\* B\

```

## 代码块支持语言

https://prismjs.com/#supported-languages

https://docusaurus.io/zh-CN/docs/markdown-features/code-blocks