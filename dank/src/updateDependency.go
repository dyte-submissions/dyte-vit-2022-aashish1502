package dank

import (
	"fmt"
	"github.com/tidwall/sjson"
	"io/fs"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

func UpdateDependency(repo Repo, name string, version string) {

	repository := strings.Split(repo.Repo, "/")
	repoName := repository[len(repository)-1]
	repoName = strings.ReplaceAll(repoName, ".git", "")
	fmt.Println(repoName)
	x, _ := os.Getwd()
	fmt.Println(x)

	git := repo.Repo
	cmd := exec.Command("git", "clone", git)
	fmt.Println("cloning")
	err := cmd.Run()

	if err != nil {
		fmt.Println("unable to clone repository", repo.Repo)
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
	cmd = exec.Command("git", "checkout", "-b", str_1)
	err = cmd.Run()

	if err != nil {
		cmd = exec.Command("git", "checkout", str_1)
		err = cmd.Run()

		if err != nil {
			fmt.Println("Cannot change into the branch")
			os.Exit(1)
		}
	}

	content, e := ioutil.ReadFile("package.json")

	if e != nil {
		fmt.Println("cannot read package.json")
		os.Exit(1)
	}

	data, _ := sjson.Set(string(content), "dependencies."+name, "^"+version)

	ioutil.WriteFile("package.json", []byte(data), fs.FileMode(0644))

	cmd = exec.Command("git", "add", ".")
	err = cmd.Run()

	if err != nil {
		fmt.Println("cannot git add")
		os.Exit(1)
	}

	cmd = exec.Command("git", "commit", "-m", str)
	err = cmd.Run()

	if err != nil {
		fmt.Println("cannot commit")
		os.Exit(1)
	}

	cmd = exec.Command("git", "push", "-u", "origin", str_1)
	err = cmd.Run()

	if err != nil {
		fmt.Println("unable to push")
		os.Exit(1)
	}

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
