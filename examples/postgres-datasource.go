package main

import (
	"fmt"
	"time"

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
	j.Verbose = true
	j.Output = "."
	j.Executable = "../jasperstarter/bin/jasperstarter"
	j.DbConnection.JdbcDir = "../jasperstarter/jdbc"

	if err := j.Compile("postgres-report.jrxml"); err != nil {
		fmt.Println(err.Error())
	}

	j.Output = fmt.Sprintf("./tmp/%d", time.Now().UnixMilli())
	if b, err := j.Process("postgres-report.jasper"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
}
