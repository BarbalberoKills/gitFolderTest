package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
)

func main() {
	cpuprofile := flag.String("cpuprofile", "", "write cpu profile to `file`")
	memprofile := flag.String("memprofile", "", "write memory profile to `file`")
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	var newString = "hello"
	fmt.Println(newString)
	reversed := reverseStringHeaviest(newString)
	fmt.Println(reversed)

	if *memprofile != "" {
		f, err := os.Create(*memprofile)
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		runtime.GC()    // get up-to-date statistics
		// Lookup("allocs") creates a profile similar to go test -memprofile.
		// Alternatively, use Lookup("heap") for a profile
		// that has inuse_space as the default index.
		if err := pprof.Lookup("allocs").WriteTo(f, 0); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}
}

// version written by a new hired guy
func reverseStringHeaviest(s string) string {
	var newString string

	for i := len(s) - 1; i >= 0; i-- {
		newString = newString + string(s[i])
	}

	return newString
}

// version written by be
func reverseStringSmart(s string) string {
	runes := []rune(s)

	i := 0
	j := len(runes) - 1

	for i < j {
		temp := runes[i]
		runes[i] = runes[j]
		runes[j] = temp

		i = i + 1
		j = j - 1
	}

	return string(runes)
}

// version that ended up in prod after I got a review from a senior dev
func reverseStringSmarter(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {

		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}
