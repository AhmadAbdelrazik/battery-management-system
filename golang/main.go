package main

import "os"

func main() {
	inputRaw, _ := os.ReadFile("ocv-vs-soc.csv")

	data := PopulateData(string(inputRaw))

}
