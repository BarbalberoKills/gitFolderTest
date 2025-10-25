package main

import (
	"bufio"
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
	if len(os.Args) < 2 {
		fmt.Println("Missing arguments")
		fmt.Println("Usage: bulklist-md <filename> [optional-content]")
		os.Exit(1)
	}

	// if len(os.Args) == 3 {
	// 	content = os.Args[2]
	// }

	list, path := getFile(os.Args[1])

	// filtered := []string{}
	// ignore := []string{"---", "sorting-spec: |-", "  sortspec"}
	// for i := 0; i < len(list); i++ {
	// 	if !slices.Contains(ignore, list[i]) {
	// 		filtered = append(filtered, list[i][2:])
	// 	}
	// }

	root := Node{
		Name:   "root",
		Indent: -1,
		Path:   path,
	}

	stack := []*Node{}

	stack = append(stack, &root)

	for i := 0; i < len(list); i++ {
		indent := 0
		var name string
		for _, c := range list[i] {
			if string(c) != " " {
				name = list[i][indent:]
				break
			}
			indent += 1
		}

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
		child.Maker()
	}
}

func (n *Node) Maker() {
	if len(n.Children) == 0 {
		touch(n.Path, n.Name, "md", "")
	} else {
		makeFolder(n.Path, n.Name)
	}
	for _, child := range n.Children {
		child.Maker()
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

func makeFolder(path, name string) {
	folderName := path + "/" + name
	err := os.MkdirAll(folderName, 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("")
	fmt.Println("Created folder: " + folderName)
}

func touch(path, name, ext, content string) {
	fileName := path + "/" + name + "." + ext
	err := os.WriteFile(fileName, []byte(content), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	if content != "" {
		fmt.Println("Created file: " + fileName + " with content" + content)
	} else {
		fmt.Println("")
		fmt.Println("Created file: " + fileName)
	}
}
