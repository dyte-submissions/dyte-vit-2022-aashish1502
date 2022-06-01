package main

import (
	dank "dank/src"
	"flag"
	"fmt"
	"os"
	"strings"
)

type Repo struct {
	name string
	repo string
}

func main() {

	getCmd := flag.NewFlagSet("check", flag.ExitOnError)
	getCSV := getCmd.String("input", "", "Reads a CSV and checks")
	getOutput := getCmd.String("output", "", "The output name of the file <optional>")
	getPack := getCmd.String("pack", "", "The name of the package along with it's versions")
	getUpdate := getCmd.Bool("update", false, "if you need to update after checking <optional>")

	if len(os.Args) < 2 {

		fmt.Println("dank: incorrect use of command")
		fmt.Printf("dank: to know more about them use `dank --help`")
		os.Exit(1)
	}

	switch os.Args[1] {

	case "check":
		handler(getCmd, getCSV, getPack, getOutput, getUpdate)

	default:

	}

}

func handler(getCmd *flag.FlagSet, name *string, pack *string, output *string, update *bool) {

	getCmd.Parse(os.Args[2:])

	data := [][]string{
		{"name", "repo", "version", "version_satisfired"},
	}

	if *name == "" {
		fmt.Println("incorrect input type for <name> file")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	if *pack == "" {
		fmt.Println("incorrect dependency name please check and use it correctly")
		getCmd.PrintDefaults()
		os.Exit(1)
	}

	dependency := strings.Split(*pack, "@")

	csvData := dank.ReadFile(*name)

	for _, line := range csvData {

		repositoryData := dank.Repo{
			Name: line[0],
			Repo: line[1],
		}

		JSONData := dank.GetPackageData(repositoryData)

		currVersion, satisfies := dank.CheckVersion(JSONData, dependency[0], dependency[1])

		var x []string

		if satisfies {
			x = []string{repositoryData.Name, repositoryData.Repo, currVersion, "yes"}
		} else {
			x = []string{repositoryData.Name, repositoryData.Repo, currVersion, "no"}
		}
		data = append(data, x)

	}

	if *output == "" {
		dank.WriteData(*name, data)
	} else {
		fmt.Println(*output)
		if strings.Split(*output, ".")[1] != "csv" {
			dank.WriteData(*output+".csv", data)
		} else {
			dank.WriteData(*output, data)
		}
	}

}
