package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
)

// Define a custom type for a slice of strings
type stringSlice []string

func readOfFile(inp string, nee bool) {
	file, err := os.Open(inp)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scan := bufio.NewScanner(file)
	var inputLines stringSlice
	for scan.Scan() {
		inputLines = append(inputLines, scan.Text())
	}

	if nee {
		sortedList := stringSlice(inputLines).sorted()
		for _, val := range sortedList {
			fmt.Println(val)
		}
	} else {
		for _, val := range inputLines {
			fmt.Println(val)
		}
	}
}

func (ss stringSlice) sorted() stringSlice {
	uniqueMap := make(map[string]struct{})
	var uniqueList stringSlice
	for _, val := range ss {
		if _, ok := uniqueMap[val]; !ok {
			uniqueMap[val] = struct{}{}
			uniqueList = append(uniqueList, val)
		}
	}
	sort.Strings(uniqueList)
	return uniqueList
}

func main() {
	bif := flag.String("r", "", "input file name")
	nee := flag.Bool("s", false, "whether to print sorted lines")
	flag.Parse()

	if *bif != "" {
		readOfFile(*bif, *nee)
	}
}
