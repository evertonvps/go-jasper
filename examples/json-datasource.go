package main

import (
	"fmt"
	"time"

	gojasper "github.com/evertonvps/go-jasper"
)

func main() {
	jsonFile := "posts.json"
	jsonQuery := ""
	output := "."
	parms := []gojasper.Parameter{}

	j := gojasper.NewGoJasperJsonData(jsonFile, jsonQuery, parms, "pdf", output)

	j.Executable = "../JasperStarter/bin/jasperstarter"
	j.Verbose = true

	if err := j.Compile("post-json.jrxml"); err != nil {
		fmt.Println(err.Error())
	}

	j.Output = fmt.Sprintf("./tmp/%d", time.Now().UnixMilli())
	if b, err := j.Process("post-json.jasper"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
}
