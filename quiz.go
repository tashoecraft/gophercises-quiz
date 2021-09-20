package main

import (
        "encoding/csv"
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
        "strings"
        "io"
)

func main() {
	//default to opening problems.csv, but if a flag is provided, use the passed in filepath

	// os package allows for opening file "os.Open(filePath)"

	filePath := flag.String("f", "./problems.csv", "CSV file to parse")

	f, err := os.Open(*filePath)
	check(err)

	scannedFile := bufio.NewScanner(f)

        success := 0
        total := 0
	for scannedFile.Scan() {
           r := csv.NewReader(strings.NewReader(scannedFile.Text()))

          for {
            record, err := r.Read()

            if err == io.EOF {
              break;
            }

            check(err)

            fmt.Printf("Question: %s ", record[0])

            buf := bufio.NewReader(os.Stdin)
            answerWithNewLine, err := buf.ReadString('\n')
            check(err)
            
            answer := strings.ToLower(strings.TrimSpace(strings.TrimSuffix(answerWithNewLine, "\n")))
            if answer == strings.ToLower(record[1]) {
              success++
            }

            total++
          }
	}

        fmt.Printf("You got %d right out of %d\n", success, total)
        f.Close()
}

func check(e error) {
	if e != nil {
		log.Fatal(e)
		panic(e)
	}
}
