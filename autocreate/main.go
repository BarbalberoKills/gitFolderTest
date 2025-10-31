package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

type Node struct {
	Name     string
	Indent   int
	Path     string
	Children []*Node
}

type Params struct {
	Dry      *bool
	Content  *string
	Ext      *string
	Filename string
}

func main() {
	params := Params{}
	params.getParams()
	list, rootPath := readFile(os.Args[len(os.Args)-1])

	root := Node{
		Name:   "root",
		Indent: -1,
		Path:   *rootPath,
	}

	stack := []*Node{&root}

	for _, line := range *list {
		name, indent := indentCount(line)

		for indent <= stack[len(stack)-1].Indent {
			stack = stack[:len(stack)-1]
		}

		parent := stack[len(stack)-1]

		newNode := &Node{
			Name:   name,
			Indent: indent,
			Path:   parent.Path + "/" + parent.Name,
		}
		if parent.Indent == -1 {
			newNode.Path = parent.Path
		}

		parent.Children = append(parent.Children, newNode)

		stack = append(stack, newNode)
	}

	for _, child := range root.Children {
		child.Maker(&params)
	}
}

func (p *Params) getParams() {
	p.Dry = flag.Bool("dry", false, "Simulation mode")
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

func readFile(passedFile string) (*[]string, *string) {
	file, err := os.Open(passedFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	filePath := filepath.Dir(file.Name())

	var list []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	defer file.Close()

	return &list, &filePath
}

func indentCount(line string) (string, int) {
	indent := 0
	var name string
	for _, c := range line {
		if string(c) != " " {
			name = line[indent:]
			break
		}
		indent += 1
	}
	return name, indent
}

func (n *Node) Maker(p *Params) {
	if len(n.Children) == 0 {
		touch(n.Path, n.Name, *p.Ext, *p.Content, *p.Dry)
	} else {
		makeFolder(n.Path, n.Name, *p.Dry)
	}
	for _, child := range n.Children {
		child.Maker(p)
	}
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

func makeFolder(path, name string, dry bool) {
	folderName := path + "/" + name
	if !dry {
		err := os.MkdirAll(folderName, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	fmt.Println("Created folder: \"" + folderName + "\"")
}
