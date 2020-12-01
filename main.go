package main

import (
	"errors"
	"os"

	"github.com/assmdx/gsitg/dep"
	"github.com/assmdx/gsitg/utils/error_"
)

func getParam() (string, string, string, error) {
	if len(os.Args) <= 3 {
		return "", "", "", errors.New("参数个数错误")
	}

	packageName := os.Args[0]
	sourceDir := os.Args[1]
	imgFilePath := os.Args[2]

	return packageName, sourceDir, imgFilePath, nil
}

func main() {
	packageName_, sourceDir, imgFilePath, err := getParam()
	error_.ErrorHandler(err)

	dep.Analysis(packageName_, sourceDir, imgFilePath)
}
