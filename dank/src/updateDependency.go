package dank

import (
	"encoding/json"
	"fmt"
	"github.com/tidwall/sjson"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func UpdateDependency(repo Repo, name string, version string) {

	repository := strings.Split(repo.repo, "/")
	repoName := repository[len(repository)-1]
	repoName = strings.ReplaceAll(repoName, ".git", "")
	fmt.Println(repoName)
	x, _ := os.Getwd()
	fmt.Println(x)

	git := repo.repo
	cmd := exec.Command("git", "clone", git)
	fmt.Println("cloning")
	err := cmd.Run()

	if err != nil {
		fmt.Println("unable to clone repository", repo.repo)
		os.Exit(1)
	}

	err = os.Chdir(x + "/" + repoName)

	//cmd = exec.Command("cd", repoName)
	//err = cmd.Run()

	if err != nil {
		fmt.Println("error occured!")
		os.Exit(1)
	}

	fmt.Println(git)

	str := "dank: updated " + name + " to version " + version
	str_1 := "dank/update-1"
	cmd = exec.Command("git", "checkout", str_1, "||", "git", "checkout", "-b", str_1)
	err = cmd.Run()

	content, e := ioutil.ReadFile("package.json")

	if e != nil {
		fmt.Println("cannot read package.json")
		os.Exit(1)
	}

	data, _ := sjson.Set(string(content), "dependencies."+name, "^"+version)

	fmt.Println(data)

	JSONData, _ := json.Marshal(data)

	fmt.Println(JSONData)

	check := json.Valid(JSONData)

	if check {
		fmt.Println("JSON valid")
	} else {
		fmt.Println("JSON invalid")
	}

	ioutil.WriteFile("package.json", []byte(data), fs.FileMode(0644))

	cmd = exec.Command("git", "add", ".")
	cmd.Run()

	cmd = exec.Command("git", "commit", "-m", str)
	cmd.Run()

	cmd = exec.Command("git", "push", "-u", "origin", str_1)
	cmd.Run()

	fmt.Println(x)
	err = os.Chdir(x)

	if err != nil {
		fmt.Println("cannot go back to the root dir")
		os.Exit(1)
	}

	Clean(x + repoName)

}

func Clean(repoName string) {

	fmt.Println(repoName)
	err := os.RemoveAll(repoName)

	if err != nil {
		fmt.Println("There was an error in deleting a file")
		os.Exit(1)
	}

}
