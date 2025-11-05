package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

// come al solito ci sono eccezioni alle regole. qui visto che interagisci con l'input utente e very likely there is nothing to test here puoi exitare diretamente qui. At this point you haven't done much os no problem in ending the program.
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

	// qui non ero sicuro in realta' tu cosa ti aspetti di passare come content. se un header diverso o proprio un contenuto passato al fly
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

func createSortspecFile(path, name, ext, content string, dry bool) error {
	// Cleanly handle the extension.
	if ext != "" && !strings.HasPrefix(ext, ".") {
		ext = "." + ext
	}

	// Use filepath.Join for cross platform safety. Se cambi da linux a windows a mac it still works
	fileName := filepath.Join(path, name+ext)

	// avoid nested ifs it's clearer to read.
	if dry {
		fmt.Printf("[DRY RUN] Would create file: \"%s\" with content:\n%s\n---\n", fileName, content)
		return nil
	}

	// Write the file with 0644 permissions (standard for data files).
	err := os.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		return err // Return the error, don't exit.
	}

	fmt.Printf("Created file: \"%s\"\n", fileName)
	return nil
}
