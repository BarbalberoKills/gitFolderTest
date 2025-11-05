package main

import (
	"flag"
	"fmt"
	"os"
)

// Essendo una roba valida per tutto e' buona idea dichiararla come costante. Easy to manipulate
const DefaultHeader = "---\nsorting-spec: |-\n sortspec"

type Params struct {
	Dry       *bool
	Recursive *bool
	// Ho tolto il ptr qui perche' cosi possiamo fare della validation quando parsiamo i parametri.
	// Lo farei anche per gli altri field ma al momento c'e' poco da validare
	Content  string
	Ext      *string
	Filename string
}

type Folder struct {
	Path     string
	Children []string
}

func main() {
	var params Params
	params.getParams()

	// a very good practice is to have all of your functions return errors and have a os.Exit() only in the main.
	// that way your program doesn't exit a minchia di cane in mezzo alla sua esecuzione mentre altre robe magari stanno runnando
	// inoltre cosi e' piu testabile. se fai un exit in una funzione e vuoi testare quella funzione sei fregato perche la funzione exita nel mezzo del tuo test e va tutto a disoneste.
	if err := processFolders(params); err != nil {
		fmt.Printf("couldn't process folders: %v", err)
		os.Exit(1)
	}
}

// wip
func processFolders(p Params) error {
	//foldersToEvaluate := []string{params.Filename}
	//folders := folderInspector(foldersToEvaluate)

	//for _, f := range folders {
	//	content := "---\nsorting-spec: |-\n  sortspec"
	//	sort.Strings(f.Children)
	//	for _, i := range f.Children {
	//		content += "\n  " + strings.Split(i, ".")[0]
	//	}
	//	content += "\n---"
	//	touch(f.Path, "sortspec", *params.Ext, content, *params.Dry)
	//}
	return nil
}

func (p *Params) getParams() {
	p.Dry = flag.Bool("dry", false, "Simulation mode")
	p.Recursive = flag.Bool("recursive", false, "Recursive run")
	p.Ext = flag.String("ext", "", "File extension")

	contentPtr := flag.String("content", "", "Content to write")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [-dry] [-recursive] [-content <text>] [-ext <ext>] <directory>\n", os.Args[0])
		fmt.Fprintln(os.Stderr, "\nOptions:")
		flag.PrintDefaults()
	}

	flag.Parse()

	if *contentPtr != "" {
		p.Content = *contentPtr
	} else {
		p.Content = DefaultHeader
	}

	if len(flag.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Error: Exactly one directory argument is required.")
		flag.Usage()
		os.Exit(1)
	}

	p.Filename = flag.Arg(0)

	// Qui potresti aggiungere altre validation se ne hai anche per altre flags. qui ho implementato un check
	// per accertarmi che quella passata e' veramente una directory
	info, err := os.Stat(p.Filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: Could not stat %s: %v\n", p.Filename, err)
		os.Exit(1)
	}
	if !info.IsDir() {
		fmt.Fprintf(os.Stderr, "Error: Argument %s is not a directory.\n", p.Filename)
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
