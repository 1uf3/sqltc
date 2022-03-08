package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"sqltc"
)

func main() {
	file := flag.String("file", "sample.sql", "Pass the sql file to get variable name and type.")
	flag.Parse()
	var sqlfile sqltc.SqlFile
	if err := sqlfile.File(*file); err != nil {
		log.Fatalf("sqlfile.File Error: %s", err)
	}
	columns := sqltc.Columns{}
	for _, query := range sqlfile.Queries {
		columns = sqltc.Convert(query)
	}
	json, err := json.Marshal(columns)
	if err != nil {
		log.Fatalf("Json.Marshal Error: %s", err)
	}

	fmt.Println(string(json))
}
