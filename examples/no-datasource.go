package main

import (
	"fmt"

	gojasper "github.com/evertonvps/go-jasper"
)

func main() {

	output := "."
	parms := []gojasper.Parameter{
		{Key: "go_jasper_hello", Value: "Tonho !!"},
	}

	j := &gojasper.GoJasper{
		Executable:     "../JasperStarter/bin/jasperstarter",
		Format:         "pdf",
		Parameters:     parms,
		DatasourceType: "none",

		Output: output,
	}

	j.Compile("hello_world.jrxml")

	if b, err := j.Process("hello_world.jasper"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
}
