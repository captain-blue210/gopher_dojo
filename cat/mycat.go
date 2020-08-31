package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var printLineNum bool

	flag.BoolVar(&printLineNum, "n", printLineNum, "通し番号を付与する")

	flag.Parse()

	index := 1
	for _, name := range flag.Args() {
		sf, err := os.Open(name)
		if err != nil {
			log.Fatal(err)
		}

		scanner := bufio.NewScanner(sf)
		for scanner.Scan() {
			if printLineNum {
				fmt.Fprintln(os.Stdout, strconv.Itoa(index)+": ", scanner.Text())
			} else {
				fmt.Println(scanner.Text())
			}
			index++
		}
	}
}
