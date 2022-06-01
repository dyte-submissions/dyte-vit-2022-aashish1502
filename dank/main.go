package main

import dank "dank/src"

func main() {

	//dank.CheckVersion(dank.GetPackageData(dank.Test), "axios", "0.24.0")
	dank.UpdateDependency(dank.Test, "axios", "This string was inserte by dank :)")

}
