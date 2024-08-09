package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sqlc/pkg/sqlc"

	"github.com/gocarina/gocsv"
	_ "github.com/mattn/go-sqlite3"
)

//go:embed sql/schema.sql
var dbstring string

func main() {
	db, err := sql.Open("sqlite3", "database.db")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec(dbstring)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	q := sqlc.New(db)

	file, err := os.Open("data.csv")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	input := []sqlc.NewUserParams{}

	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = ';'
		return r // Allows use semicolon as delimiter
	})

	err = gocsv.Unmarshal(file, &input)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, value := range input {
		err = q.NewUser(context.Background(), value)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}
}
