package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

var (
	// Number is the number of generated Excel files.
	Number = flag.Int(
		"n",
		10,
		"Number of generated Excel files.",
	)
)

func main() {
	os.Exit(run())
}

func run() int {
	flag.Parse()
	workingDir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	for i := 1; i <= *Number; i++ {
		xlsx := excelize.NewFile()
		xlsx.NewSheet("Sheet1")
		err = xlsx.SetCellStr("Sheet1", "A1", getRandString())
		if err != nil {
			panic(err)
		}
		err = xlsx.SetCellInt("Sheet1", "B1", getRandInt())
		if err != nil {
			panic(err)
		}
		err = xlsx.SetCellBool("Sheet1", "C1", getRandBool())
		if err != nil {
			panic(err)
		}
		fileName := fmt.Sprintf(
			"%s/test/data/spreadsheet%d.xlsx",
			workingDir,
			i,
		)
		err := xlsx.SaveAs(fileName)
		if err != nil {
			panic(err)
		}
	}
	return 0
}

func getRandString() string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, 10)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func getRandInt() int {
	return rand.Intn(100)
}

func getRandBool() bool {
	return rand.Intn(100) > 50
}
