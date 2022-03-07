package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"

	"sqltc"
)

func main() {
	file := flag.String("file", "sample.sql", "SQL file")
	flag.Parse()
	sqlfile := sqltc.SqlFile{}
	sqlfile.File(*file)
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
