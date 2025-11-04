package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"strings"
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

func main() {
	var params Params
	params.getParams()

	foldersToEvaluate := []string{params.Filename}
	folders := folderInspector(foldersToEvaluate)

	for _, f := range folders {
		content := "---\nsorting-spec: |-\n  sortspec"
		sort.Strings(f.Children)
		for _, i := range f.Children {
			content += "\n  " + strings.Split(i, ".")[0]
		}
		content += "\n---"
		touch(f.Path, "sortspec", *params.Ext, content, *params.Dry)
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
