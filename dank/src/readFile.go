package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type repo struct {
	name string
	repo string
}

/*

The readFile function will take a file name and a channel

input: filename: string
	   ch: channel

proccess:
		it will read the csv files and pass it into the channel which will be read by another function which will do other things

*/

func readFile(file string) {

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

	for _, line := range csvData {

		repositoryData := repo{
			name: line[0],
			repo: line[1],
		}

		//TODO: channel it to a function that takes the data and uses the repo link.
		fmt.Println(repositoryData)

	}

}

func main() {

	readFile("test.csv")

}
