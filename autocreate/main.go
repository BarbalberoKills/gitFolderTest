package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"slices"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing arguments")
		fmt.Println("Usage: bulklist-md <filename> [optional-content]")
		os.Exit(1)
	}

	content := ""
	if len(os.Args) == 3 {
		content = os.Args[2]
	}

	file := getFile(os.Args[1])
	defer file.Close()

	list := []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		list = append(list, scanner.Text())
	}

	filtered := []string{}
	ignore := []string{"---", "sorting-spec: |-", "  sortspec"}
	for i := 0; i < len(list); i++ {
		if !slices.Contains(ignore, list[i]) {
			filtered = append(filtered, list[i][2:])
		}
	}

	for _, note := range filtered {
		touch(filepath.Dir(file.Name()), note, "md", content)
	}
}

func getFile(passedFile string) *os.File {
	file, err := os.Open(passedFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return file
}

func touch(path, name, ext, content string) {
	fileName := path + "/" + name + "." + ext
	err := os.WriteFile(fileName, []byte(content), 0755)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Created file: " + fileName + "with content" + content)
}
