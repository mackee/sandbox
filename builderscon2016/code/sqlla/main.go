package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	printSQL()
	mappingSQL()
}

func printSQL() {
	q := NewExampleSQL().Select().ID(100)
	fmt.Println(q.ToSql())
}

func mappingSQL() {
	db, _ := sql.Open("sqlite3", "./foo.db")
	row, _ := NewExampleSQL().Select().ID(100).Single(db)
	fmt.Printf("%+v\n", row)
}
