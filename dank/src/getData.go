package dank

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

var Test = Repo{

	Name: "dyte-react-sample-app",
	Repo: "https://github.com/aashish1502/react-sample-app.git",
}

func GetPackageData(repository Repo) []byte {

	url := "https://raw.githubusercontent.com"
	extention := strings.Split(repository.Repo, "github.com")[1]
	url += extention

	packageJSONurl := url + "/main/package.json"

	res, err := http.Get(packageJSONurl)

	if err != nil {
		fmt.Println("The Url for the repository", repository.Repo, "is incorrect")
		os.Exit(1)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	var JSONData map[string]interface{}

	json.Unmarshal(data, &JSONData)

	DependenciesJSON, _ := json.Marshal(JSONData["dependencies"])

	//fmt.Printf(string(DependenciesJSON))
	return DependenciesJSON
}

func GetPackageLockData(repository Repo) []byte {

	url := "https://raw.githubusercontent.com"
	extention := strings.Split(repository.Repo, "github.com")[1]
	url += extention

	packageLockJSONurl := url + "main/package-lock.json"

	res, err := http.Get(packageLockJSONurl)

	if err != nil {
		fmt.Println("The Url for the repository", repository.Repo, "is incorrect")
		os.Exit(1)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	var JSONData map[string]interface{}

	json.Unmarshal(data, &JSONData)

	test, _ := json.Marshal(JSONData["dependencies"])

	return test

}

func CheckVersion(JSONData []byte, name string, version string) (string, bool) {

	var data map[string]interface{}

	json.Unmarshal(JSONData, &data)

	if data[name] == "^"+version {
		return data[name].(string), true
	} else {
		return data[name].(string), false
	}

}
