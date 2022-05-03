package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("SP// Backend Developer Test - Input Processing")
	fmt.Println()

	// Read STDIN into a new buffered reader
	reader := bufio.NewReader(os.Stdin)

	for {
		var (
			isPrefix       = true
			err            error
			line, fullLine []byte
		)

		for isPrefix && err == nil {
			line, isPrefix, err = reader.ReadLine()
			fullLine = append(fullLine, line...)
		}

		if err != nil {
			if err == io.EOF {
				log.Println("reached end of file")
				break
			}
			log.Fatalf("error reading from reader: %v\n", err)
		}

		str := string(fullLine)
		if strings.Contains(str, "error") {
			log.Println(str)
		}
	}
}
