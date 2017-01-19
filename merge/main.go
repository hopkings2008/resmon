package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"

	"github.com/zouyu/resmon/parser"
)

var files []string

var regMatch = regexp.MustCompile("_result.json")

func visit(path string, f os.FileInfo, err error) error {
	if f.IsDir() {
		return nil
	}
	fmt.Printf("Visited: %s\n", path)
	if regMatch.MatchString(path) {
		files = append(files, path)
	}
	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, visit)
	if len(files) == 0 {
		fmt.Printf("no specified file found on %s\n", root)
		return
	}
	p, _ := parser.CreateParser("(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S+)\\s+(\\S).*", 0, "v1/image/get")
	err := p.Import(files...)
	if err != nil {
		fmt.Printf("failed to parse files, err: %v\n", err)
		return
	}

	err = p.Save("./final_result.json")
	if err != nil {
		fmt.Printf("failed to save result, err: %v\n", err)
		return
	}
	fmt.Printf("succeed to save result.\n")
}
