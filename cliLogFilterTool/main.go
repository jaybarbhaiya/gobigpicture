package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	path := flag.String("path", "myapp.log", "The path of the log file to be analyzed")
	level := flag.String("level", "ERROR", "The log level that must be filtered. Options DEBUG, INFO, WARNING, ERROR or CRITICAL")

	flag.Parse()

	file, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		sString, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		if strings.Contains(sString, *level) {
			fmt.Println(sString)
		}
	}
}
