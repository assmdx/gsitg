# gsitg

A static source code analysis tool, which is used to analysis and visualize the module dependency of your go package source code.

## Installtion

```
go get github.com/assmdx/gsitg
```

## Usage: Analysis your source code dependency

```
gsitg.dep.Analysis(yourPackageName, sourceCodeDir, resultPngFilePath)
```
You will get the result like this:

![](https://s3.ax1x.com/2020/12/08/DzRMB4.png)

## Test

```
    make test
```
