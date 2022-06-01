package dank

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Repo struct {
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

	var myMap map[string]Repo

	for _, line := range csvData {

		repositoryData := Repo{
			name: line[0],
			repo: line[1],
		}

		myMap[line[1]] = repositoryData

		//TODO: channel it to a function that takes the data and uses the repo link.
		fmt.Println(repositoryData)

	}

}

func WriteData(file string, check string, repository Repo, myMap map[string]Repo, version_satisfied bool) {

	csvFile, err := os.Create(file)

	if err != nil {
		fmt.Println("can't write in the csv file make sure the file is correct")
		os.Exit(1)
	}
	defer csvFile.Close()

}
