package main

import (
	"errors"
	"fmt"
	"os"
	"utils"
	"strings"
	"path"
	"path/filepath"
	"github.com/freenerd/go-import-extractor/extractor"
)

var graphTemplate = `digraph {
	{{- if eq .direction "horizontal" -}}
	rankdir=LR;
	{{ end -}}
	node [shape=box];
	{{ range $mod, $modId := .mods -}}
	{{ $modId }} [label="{{ $mod }}"];
	{{ end -}}
	{{- range $modId, $depModIds := .dependencies -}}
	{{- range $_, $depModId := $depModIds -}}
	{{ $modId }} -> {{ $depModId }};
	{{  end -}}
	{{- end -}}
	}
	`

type Mapping struct {
	packageFrom string
	pacakgeTo   string
}

type DAG struct {
	parentPackage *DAG
	toPackages    []DAG
}

func initDAG() *DAG {
	dag := DAG{nil, []}
	return &dag
}

func extractPackageAndMapping(goFilePath string) []Mapping {
	// ??????
	imports, err = extractor.FileImportCalls(goFilePath)
	if err != nil {
		errorHandler(err)
	}

	packageFrom := filepath.Dir(goFilePath)

	mappings := make([]Mapping, len(imports))
	for i, _ :range mappings {
		mappings[i] = Mapping{packageFrom: packageFrom, packageTo: imports[i]}
	}

	return mappings
}

func updateMappingDAG(mapping Mapping, dag DAG) {

}

func genDot(dag) string {

}

func genPng(pngFilePath string, dot string) {

}

func getParam() (string, string, error) {
	if len(os.Args) <= 2 {
		return "", "", errors.New("参数个数错误")
	}
	dir := os.Args[0]
	imgFilePath := os.Args[1]

	isDir, err = utils.IsDir(dir)
	if (!isDir || err) {
		errorHandler(err)
	}

	filenameWithSuffix := path.Base(imgFilePath)
	fileSuffix = path.Ext(filenameWithSuffix)
	if fileSuffix != ".png" {
		errorHandler(errors.New("图片必须为png图片"))
	}

	if (utils.Exists(imgFilePath) && utils.IsFile(imgFilePath)) {
		err := os.Remove(imgFilePath)
	}

	return dir, imgFilePath, nil
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dir, imgFilePath, err := getParam()
	errorHandler(err)

	dag = initDAG()

	files, err := utils.GetAllFilesOfExt(dir, ".go")
	errorHandler(err)

	for i, goFile := range files {
		mappings := extractPackageAndMapping(goFile)
		for j, mapping :range mappings {
			updateMappingDAG(mapping, dag)
		}	
	}

	dot := genDot(dag)

	genPng(imgFilePath, dot)
}
