# go-jasper

*go-jasper*
is a Go library provides an interface for compiling and processing JasperReports. It enables seamless integration with JasperReports, allowing users to generate reports dynamically from Go applications. The library supports executing .jrxml and .jasper report templates, handling parameters, and exporting reports in various formats such as PDF and HTML. Designed for efficiency and ease of use, it simplifies report generation without requiring manual interaction with JasperStarter or Java processes.

[Download jasperstarter here](https://sourceforge.net/projects/jasperstarter/files/)

[Download postgressql jdbc driver here](https://jdbc.postgresql.org/download/postgresql-42.2.24.jar), and move to folder *jasperstarter/jdbc*


Examples
--------


### postgresql

[Postegres exameple](examples/postgres-datasource.go)
```go
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
```
### Json datasource
[Json Datasource](examples/json-datasource.go)
```go
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
```

### parameter datasource
[parameter datasource](examples/no-datasource.go)
```go
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

	j.Verbose = true
	if err := j.Compile("hello_world.jrxml"); err != nil {
		fmt.Println(err.Error())
	}

	j.Output = fmt.Sprintf("./tmp/%d", time.Now().UnixMilli())
	if b, err := j.Process("hello_world.jasper"); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(string(b))
	}
```