package jasper

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
)

type Parameter struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type DbConnection struct {
	// 	-u <dbuser> database user
	User string `json:"dbuser"`
	// 	-p <dbpasswd> database password
	Pass string `json:"dbpasswd"`
	// 	-H <dbhost> database host
	Host string `json:"dbhost"`
	// 	-n <dbname>  database name
	DbName string `json:"dbname"`
	// 	--db-sid <sid> oracle sid
	Sid string `json:"db-sid"`
	// 	--db-port <port>database port
	Port int `json:"db-port"`
	// 	--db-driver <name> jdbc driver class name for use with type: generic
	JdbcDriver string `json:"db-driver"`
	// 	--db-url <jdbcUrl>     jdbc url without user, passwd with type:generic
	JdbcUrl string `json:"db-url"`
	// 	--jdbc-dir  <dir>       directory where jdbc driver jars are located. Defaults to ./jdbc
	JdbcDir string `json:"jdbc-dir"`
}

type GoJasper struct {
	Executable string
	//-f view, print, pdf, rtf, xls, xlsMeta, xlsx, docx, odt, ods, pptx, csv, csvMeta, html, xhtml, xml, jrprint
	Format string `json:"format"`
	Locale string
	Output string

	Parameters   []Parameter `json:"parameters"`
	Resources    string
	DbConnection *DbConnection `json:"db_connection"`
	//-t <dstype> datasource type: none, csv, xml, json, mysql, postgres, oracle, generic (jdbc)
	DatasourceType string `json:"dstype"`
	// 	--data-file <file>     input file for file based datasource
	DataFile string `json:"data-file"`

	// 	--csv-first-row        first row contains column headers
	// 	--csv-columns <list>   Comma separated list of column names
	// 	--csv-record-del <delimiter>
	// 						   CSV Record Delimiter - defaults to line.separator
	// 	--csv-field-del <delimiter>
	// 						   CSV Field Delimiter - defaults to ","
	// 	--csv-charset <charset>
	// 						   CSV charset - defaults to "utf-8"
	// 	--xml-xpath <xpath>    XPath for XML Datasource
	// 	--json-query <jsonquery>
	// 						   JSON query string for JSON Datasource
	JsonQuery string `json:"json_query"`

	//	  output options:
	//		-N <printername>       name of printer
	//		-d                     show print dialog when printing
	//		-s <reportname>        set internal report/document name when printing
	//		-c <copies>            number of copies. Defaults to 1
	//		--out-field-del <delimiter>
	//							   Export CSV (Metadata) Field Delimiter - defaults to ","
	//		--out-charset <charset>
}

func NewGoJasperJsonData(dataFile, jsonquery string, parameters []Parameter, format string, output string) *GoJasper {
	if len(jsonquery) == 0 {
		jsonquery = "."
	}
	return &GoJasper{
		Executable:     "jasperstarter",
		Format:         format,
		Parameters:     parameters,
		DatasourceType: "json",
		JsonQuery:      jsonquery,
		DataFile:       dataFile,
		Output:         output,
	}
}

func NewGoJasper(datasourceType string, dbConnection *DbConnection, parameters []Parameter, format string) *GoJasper {

	return &GoJasper{
		Executable:     "jasperstarter",
		Format:         format,
		Parameters:     parameters,
		DatasourceType: datasourceType,
		DbConnection:   dbConnection,
	}
}

func (g *GoJasper) Compile(reportSource string) error {

	if _, err := os.Stat(reportSource); os.IsNotExist(err) {
		return errors.New("Invalid input file")
	}
	args := []string{"compile", filepath.Clean(reportSource)}

	if len(g.Output) > 0 {
		args = append(args, "-o")
		args = append(args, g.Output)
	}
	_, err := g.execute(args...)
	return err
}

func (g *GoJasper) Process(input string) ([]byte, error) {

	args := []string{"process", input}

	if len(g.Locale) > 0 {
		args = append(args, "--locale")
		args = append(args, g.Locale)
	}
	if len(g.Output) > 0 {
		args = append(args, "-o")
		args = append(args, g.Output)
	}

	args = append(args, "-f")
	args = append(args, g.Format)

	if len(g.Parameters) > 0 {
		args = append(args, "-P")

		for _, parm := range g.Parameters {
			if reflect.TypeOf(parm.Value).Kind() == reflect.String {
				args = append(args, fmt.Sprintf("%s=\"%s\"", parm.Key, parm.Value))
			} else {
				args = append(args, fmt.Sprintf("%s=%s", parm.Key, parm.Value))
			}

		}

	}

	if len(g.DatasourceType) == 0 {
		args = append(args, "-t")
		args = append(args, "none")
	} else {
		args = append(args, "-t")
		args = append(args, g.DatasourceType)
	}

	if g.DatasourceType == "json" {
		args = append(args, "--json-query")
		args = append(args, g.JsonQuery)

		if len(g.DataFile) > 0 {
			args = append(args, "--data-file")
			args = append(args, g.DataFile)
		}

	}

	if g.DbConnection != nil {

		if len(g.DbConnection.User) > 0 {
			args = append(args, "-u")
			args = append(args, g.DbConnection.User)
		}

		if len(g.DbConnection.Pass) > 0 {
			args = append(args, "-p")
			args = append(args, g.DbConnection.Pass)
		}

		if len(g.DbConnection.Host) > 0 {
			args = append(args, "-H")
			args = append(args, g.DbConnection.Host)
		}

		if len(g.DbConnection.DbName) > 0 {
			args = append(args, "-n")
			args = append(args, g.DbConnection.DbName)
		}

		if g.DbConnection.Port > 0 {
			args = append(args, "--db-port", strconv.Itoa(g.DbConnection.Port))
		}

		if len(g.DbConnection.JdbcDriver) > 0 {
			args = append(args, "--db-driver")
			args = append(args, g.DbConnection.JdbcDriver)
		}

		if len(g.DbConnection.JdbcUrl) > 0 {
			args = append(args, "--db-url")
			args = append(args, g.DbConnection.JdbcUrl)
		}

		if len(g.DbConnection.JdbcDir) > 0 {
			args = append(args, "--jdbc-dir")
			args = append(args, g.DbConnection.JdbcDir)
		}

		if len(g.DbConnection.Sid) > 0 {
			args = append(args, "--db-sid")
			args = append(args, g.DbConnection.Sid)
		}

	}

	return g.execute(args...)
}

func (g *GoJasper) execute(args ...string) ([]byte, error) {
	if len(args) < 2 {
		return nil, errors.New("Invalid command executable")
	}

	if _, err := os.Stat(g.Executable); os.IsNotExist(err) {
		return nil, errors.New("Invalid resource directory")
	}
	fmt.Println(args)
	cmd := exec.Command(g.Executable, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error on execute JasperStarter:", err)
		fmt.Println("error:", stderr.String())

	}

	return cmd.Output()
}
