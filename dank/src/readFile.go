package dank

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type Repo struct {
	Name string
	Repo string
}

/*

The readFile function will take a file name and a channel

input: filename: string
	   ch: channel

proccess:
		it will read the csv files and pass it into the channel which will be read by another function which will do other things

*/

func ReadFile(file string) [][]string {

	csvFile, err := os.Open(file)

	if err != nil {
		fmt.Println("can't read the csv file make sure it is of the right encoding")
		os.Exit(1)
	}

	defer csvFile.Close()

	csvData, err := csv.NewReader(csvFile).ReadAll()

	if err != nil {
		fmt.Println("something's wrong i can feel it")
		os.Exit(1)
	}

	csvData = csvData[1:]
	return csvData

}

func WriteData(file string, data [][]string) {

	f, err := os.Create(file)
	defer f.Close()

	if err != nil {

		log.Fatalln("failed to open file", err)
	}

	w := csv.NewWriter(f)
	defer w.Flush()

	for _, record := range data {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
}
