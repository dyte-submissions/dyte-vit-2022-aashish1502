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

	name: "dyte-react-sample-app",
	repo: "https://github.com/dyte-in/react-sample-app/",
}

func GetPackageData(repository Repo) []byte {

	url := "https://raw.githubusercontent.com"
	extention := strings.Split(repository.repo, "github.com")[1]
	url += extention

	packageJSONurl := url + "main/package.json"

	res, err := http.Get(packageJSONurl)

	if err != nil {
		fmt.Println("The Url for the repository", repository.repo, "is incorrect")
		os.Exit(1)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	var JSONData map[string]interface{}

	json.Unmarshal(data, &JSONData)

	DependenciesJSON, _ := json.Marshal(JSONData["dependencies"])

	return DependenciesJSON
}

func GetPackageLockData(repository Repo) {

	url := "https://raw.githubusercontent.com"
	extention := strings.Split(repository.repo, "github.com")[1]
	url += extention

	packageLockJSONurl := url + "main/package-lock.json"

	res, err := http.Get(packageLockJSONurl)

	if err != nil {
		fmt.Println("The Url for the repository", repository.repo, "is incorrect")
		os.Exit(1)
	}

	defer res.Body.Close()

	data, _ := ioutil.ReadAll(res.Body)

	var JSONData map[string]interface{}

	json.Unmarshal(data, &JSONData)

	test, _ := json.Marshal(JSONData["dependencies"])

	fmt.Printf(string(test))

}

func CheckVersion(JSONData []byte, name string, version string) {

	var data map[string]interface{}

	json.Unmarshal(JSONData, &data)

	if data[name] == "^"+version {
		fmt.Println("Correct Version")
	} else {
		fmt.Println("incorrect version")
	}

}
