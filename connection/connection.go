package connection

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	// "strings"

	_ "github.com/lib/pq"
	"github.com/xo/dburl"
)

type DB struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

var (
	CONFIG_FILE string = "./config.json" // relative path from main.go
	dbs         []*DB
	currentDB   *sql.DB
)

func GetDBs() []*DB {
	if len(dbs) != 0 {
		return dbs
	}
	bytes, err := os.ReadFile(CONFIG_FILE)
	if err != nil {
		log.Fatal(err)
	}

	if err := json.Unmarshal(bytes, &dbs); err != nil {
		log.Fatal(err)
	}

	return dbs
}

func GetSchema(name string) ([]string, error) {
	res := []string{}
	// check if DB has already asigned
	if currentDB == nil {
		for _, db := range GetDBs() {
			if name == db.Name {
				url, err := dburl.Parse(db.Url)
				if err != nil {
					return res, err
				}
				db, err := sql.Open(url.Driver, url.DSN)
				currentDB = db
				// res = append(res, url.Driver)
				// res = append(res, url.DSN)
				if err != nil {
					return res, err
				}
				break
			}
		}
	}

	rows, err := currentDB.Query("SELECT schema_name from information_schema.schemata")
	// rows, err := currentDB.Query("SELECT 1;")
	if err != nil {
		return res, err
	}
	fieldsList := rowsToStrings(rows)[1:] // remove header line
	for _, fields := range fieldsList {
		res = append(res, strings.Join(fields, ", "))
	}

	return res, nil
}

func rowsToStrings(rows *sql.Rows) [][]string {
	cols, err := rows.Columns()
	if err != nil {
		panic(err)
	}
	pretty := [][]string{cols}
	results := make([]interface{}, len(cols))
	for i := range results {
		results[i] = new(interface{})
	}
	for rows.Next() {
		if err := rows.Scan(results[:]...); err != nil {
			panic(err)
		}
		cur := make([]string, len(cols))

		for i := range results {
			val := *results[i].(*interface{})
			var str string
			if val == nil {
				str = "NULL"
			} else {
				switch v := val.(type) {
				case []byte:
					str = string(v)
				default:
					str = fmt.Sprintf("%v", v)
				}
			}
			cur[i] = str
		}
		pretty = append(pretty, cur)
	}
	return pretty
}
