package main

import (
	"fmt"

	gojasper "github.com/evertonvps/go-jasper"
)

func main() {
	jsonFile := "posts.json"
	jsonQuery := ""
	output := "."
	parms := []gojasper.Parameter{}

	j := gojasper.NewGoJasperJsonData(jsonFile, jsonQuery, parms, "pdf", output)

	j.Executable = "../JasperStarter/bin/jasperstarter"

	j.Compile("post-json.jrxml")

	if b, err := j.Process("post-json.jasper"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
}
