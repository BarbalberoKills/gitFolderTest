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

func main() {
	list, path, content, fileType, dry := getParams()

	root := Node{
		Name:   "root",
		Indent: -1,
		Path:   path,
	}

	stack := []*Node{&root}

	for _, line := range list {
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
		child.Maker(content, fileType, dry)
	}
}

func getParams() ([]string, string, string, string, bool) {
	dry := flag.Bool("dry", false, "Simulation mode")
	content := flag.String("content", "", "Content to write")
	fileType := flag.String("ext", "", "File extension")
	flag.Parse()

	os.Args = os.Args[1:]

	if *dry {
		os.Args = os.Args[1:]
	}
	if *content != "" {
		os.Args = os.Args[1:]
	}
	if *fileType != "" {
		os.Args = os.Args[1:]
		*fileType = "." + *fileType
	}

	if len(os.Args) < 1 {
		fmt.Println("Missing arguments")
		fmt.Println("Usage: treebuild [-dry] [-content] [-ext] <filename>")
		os.Exit(1)
	} else if len(os.Args) > 1 {
		fmt.Println("Too many arguments")
		fmt.Println("Usage: treebuild [-dry] [-content] [-ext] <filename>")
		os.Exit(1)
	}

	list, path := getFile(os.Args[len(os.Args)-1])

	return list, path, *content, *fileType, *dry

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

func (n *Node) Maker(content, fileType string, dry bool) {
	if len(n.Children) == 0 {
		touch(n.Path, n.Name, fileType, content, dry)
	} else {
		makeFolder(n.Path, n.Name, dry)
	}
	for _, child := range n.Children {
		child.Maker(content, fileType, dry)
	}
}

func getFile(passedFile string) ([]string, string) {
	file, err := os.Open(passedFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	list := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}
	defer file.Close()

	return list, filepath.Dir(file.Name())
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
	fmt.Println("")
	fmt.Println("Created folder: " + folderName)
}

func touch(path, name, ext, content string, dry bool) {
	fileName := path + "/" + name + ext
	if !dry {
		err := os.WriteFile(fileName, []byte(content), 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}
	if content != "" {
		fmt.Println("Created file: " + fileName + " with content" + content)
	} else {
		fmt.Println("")
		fmt.Println("Created file: " + fileName)
	}
}
