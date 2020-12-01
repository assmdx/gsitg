package dep

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/assmdx/gsitg/utils/error_"
	"github.com/assmdx/gsitg/utils/io"
	"github.com/freenerd/go-import-extractor/extractor"
	"github.com/goccy/go-graphviz"
)

type Mapping struct {
	from string
	to   string
}

type DAG struct {
	importedCount int
	module        string
	importModules []*DAG
}

var packageName string = ""
var sourceCodeDir string = ""
var dagSet map[string]*DAG = make(map[string]*DAG)

func genDAG(goFiles []string) {
	for _, goFile := range goFiles {
		mappings := ExtractMappings(goFile)
		for _, mapping := range mappings {
			dagFrom := dagSet[mapping.from]
			dagTo := dagSet[mapping.to]

			if dagFrom == nil {
				dagFrom = &DAG{importedCount: 0, module: mapping.from, importModules: make([]*DAG, 0)}
				dagSet[mapping.from] = dagFrom
			}

			if dagTo == nil {
				dagTo = &DAG{importedCount: 0, module: mapping.to, importModules: make([]*DAG, 0)}
				dagSet[mapping.to] = dagTo
			}

			var ans bool = true
			for _, d := range dagFrom.importModules {
				if mapping.to == d.module {
					ans = false
					break
				}
			}

			if ans == true {
				dagFrom.importModules = append(dagFrom.importModules, dagTo)
				dagTo.importedCount++
			}
		}
	}
}

func ExtractMappings(goFilePath string) []Mapping {
	imports, err := extractor.FileImportCalls(goFilePath)
	if err != nil {
		error_.ErrorHandler(err)
	}

	dir := filepath.Dir(goFilePath)

	from := strings.ReplaceAll(dir, sourceCodeDir, packageName)

	if err != nil {
		error_.ErrorHandler(err)
	}

	// 优化: 数组去重算法????
	mappings := make([]Mapping, 0)
	for _, to := range imports {
		if to == "" || from == to {
			continue
		}
		ans := false
		for _, m := range mappings {
			if to == m.to {
				ans = true
				break
			}
		}
		if ans == false {
			mappings = append(mappings, Mapping{from: from, to: to})
		}
	}

	return mappings
}

func genPng(pngFilePath string) {
	g := graphviz.New()
	graph, err := g.Graph()
	if err != nil {
		graph.Close()
		g.Close()
		error_.ErrorHandler(err)
	}

	defer func() {
		graph.Close()
		g.Close()
	}()

	rootNode, err := graph.CreateNode("___")
	error_.ErrorHandler(err)

	for module, dag := range dagSet {
		if dag.importedCount == 0 {
			rootToNode, err := graph.CreateNode(dag.module)
			error_.ErrorHandler(err)
			e, err := graph.CreateEdge("___"+dag.module, rootNode, rootToNode)
			error_.ErrorHandler(err)
			e.SetLabel("___" + dag.module)
		}

		if len(dag.importModules) == 0 {
			continue
		}

		fromNode, err := graph.CreateNode(module)
		error_.ErrorHandler(err)

		for _, to := range dag.importModules {
			toNode, err := graph.CreateNode(to.module)
			error_.ErrorHandler(err)

			e, err := graph.CreateEdge(module+to.module, fromNode, toNode)
			error_.ErrorHandler(err)
			e.SetLabel(module + to.module)

			fmt.Println("edge: ", module, " -> ", to.module)
		}
	}

	if err := g.RenderFilename(graph, graphviz.PNG, pngFilePath); err != nil {
		error_.ErrorHandler(err)
	}

}

func Analysis(packageName_ string, sourceCodeDir_ string, imgFilePath string) {
	packageName = packageName_
	sourceCodeDir = sourceCodeDir_

	isDir := io.IsDir(sourceCodeDir)
	if !isDir {
		error_.ErrorHandler(errors.New("source code dir is not a directory"))
	}

	filenameWithSuffix := path.Base(imgFilePath)
	fileSuffix := path.Ext(filenameWithSuffix)
	if fileSuffix != ".png" {
		error_.ErrorHandler(errors.New("must be png file"))
	}

	if io.Exists(imgFilePath) && io.IsFile(imgFilePath) {
		err := os.Remove(imgFilePath)
		error_.ErrorHandler(err)
	}

	goFiles, err := io.GetAllFilesOfExt(sourceCodeDir, ".go")
	error_.ErrorHandler(err)

	genDAG(goFiles)

	genPng(imgFilePath)
}
