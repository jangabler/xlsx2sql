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
	Mapping = flag.String("m", "", "Mapping definition file in XML format.")
)

func main() {
	// Input
	flag.Usage = usage
	flag.Parse()
	if len(*Mapping) == 0 {
		usage()
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
}

func usage() {
	flag.PrintDefaults()
	os.Exit(2)
}
