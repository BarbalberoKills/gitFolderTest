package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"gopkg.in/yaml.v3"
)

type Params struct {
	Dry       *bool
	Recursive *bool
	Content   *string
	Ext       *string
	Filename  string
}

type Folder struct {
	Path     string
	Children []string
}

type SortContent struct {
	SortingSpec string `yaml:"sorting-spec"`
}

type MyBuffer struct {
	buf []byte
}

func (m *MyBuffer) Write(p []byte) (int, error) {
	m.buf = p
	return 0, nil
}

func main() {
	var params Params
	params.getParams()

	foldersToEvaluate := []string{params.Filename}
	folders := folderInspector(foldersToEvaluate)

	var mybuf MyBuffer

	var buffer bytes.Buffer
	for _, f := range folders {
		enc := yaml.NewEncoder(&mybuf)
		enc.SetIndent(2)
		var sortContent SortContent
		for j, i := range f.Children {
			sortContent.SortingSpec += strings.Split(i, ".")[0]
			if len(f.Children) > j+1 {
				sortContent.SortingSpec += "\n"
			}
		}

		err := enc.Encode(&sortContent)
		if err != nil {
			fmt.Errorf("ERRORE GROSSO")
		}
		enc.Close()

		fmt.Println(buffer.String())
		yamlData := buffer.Bytes()
		prefix := []byte("---\n")
		sufix := []byte("---")

		finalData := append(prefix, yamlData...)
		finalData = append(finalData, sufix...)

		err = os.WriteFile(f.Path+"/sortspec.md", finalData, 0644)
		if err != nil {
			fmt.Errorf("ALTRO ERRORE GROSSO")
		}
		buffer.Reset()
	}
}

func (p *Params) getParams() {
	p.Dry = flag.Bool("dry", false, "Simulation mode")
	p.Recursive = flag.Bool("recursive", false, "Recursive run")
	p.Content = flag.String("content", "", "Content to write")
	p.Ext = flag.String("ext", "", "File extension")
	flag.Parse()
	p.Filename = flag.Arg(0)

	if len(flag.Args()) != 1 {
		fmt.Println("Exactly one argument is required")
		fmt.Println("Usage: treebuild [-dry] [-content] [-ext] <filename>")
		os.Exit(1)
	}
}

func folderInspector(foldersToEvaluate []string) []*Folder {
	folders := []*Folder{}
	for len(foldersToEvaluate) != 0 {
		files, err := os.ReadDir(foldersToEvaluate[len(foldersToEvaluate)-1])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		folders = append(folders, &Folder{Path: foldersToEvaluate[len(foldersToEvaluate)-1]})

		foldersToEvaluate = foldersToEvaluate[:len(foldersToEvaluate)-1]
		var children []string

		for _, f := range files {
			children = append(children, f.Name())
			if f.Type().IsDir() {
				foldersToEvaluate = append(foldersToEvaluate, folders[len(folders)-1].Path+"/"+f.Name())
			}
		}
		folders[len(folders)-1].Children = children
	}
	return folders
}

func touch(path, name, ext, content string, dry bool) {
	if ext != "" {
		ext = "." + ext
	}
	fileName := path + "/" + name + ext
	if !dry {
		err := os.WriteFile(fileName, []byte(content), 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if content != "" {
		fmt.Println("Created file: \"" + fileName + "\" with content: \"" + content + "\"")
	} else {
		fmt.Println("Created file: \"" + fileName + "\"")
	}
}
