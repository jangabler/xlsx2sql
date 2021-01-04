package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jangabler/xlsx2sql/pkg/extractor"
	"github.com/jangabler/xlsx2sql/pkg/mapping"
)

var (
	// Mapping is the mapping definition file in XML format.
	Mapping = flag.String("m", "", "Mapping definition file in XML format.")
)

func main() {
	os.Exit(run())
}

func run() int {
	// Input
	flag.Usage = usage
	flag.Parse()
	if len(*Mapping) == 0 {
		usage()
		return 2
	}

	// Process
	m := mapping.New()
	err := m.ReadXML(*Mapping)
	if err != nil {
		log.Fatal(err)
	}
	e := extractor.New()
	e.Run(*m)

	// Output
	for _, s := range e.Results {
		stmt := s.GenerateInsertIntoStmt()
		fmt.Println(stmt)
	}
	return 0
}

func usage() {
	flag.PrintDefaults()
}
