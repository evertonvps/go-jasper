package main

import (
	"fmt"

	gojasper "github.com/evertonvps/go-jasper"
)

func main() {

	parms := []gojasper.Parameter{}
	dbConnection := &gojasper.DbConnection{
		Host:   "0.0.0.0",
		Port:   32769,
		User:   "root",
		Pass:   "",
		DbName: "sandbox",
	}

	j := gojasper.NewGoJasper("postgres", dbConnection, parms, "pdf")
	j.Output = "."
	j.Executable = "../jasperstarter/bin/jasperstarter"
	j.DbConnection.JdbcDir = "../jasperstarter/jdbc"
	j.Compile("postgres-report.jrxml")

	if b, err := j.Process("postgres-report.jasper"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
}
