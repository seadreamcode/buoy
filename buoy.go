package main

import (
	"fmt"

	db "github.com/seadreamcode/buoy/core"
)

func newDatabase(filename string) (*db.Database, error) {
	return &db.Database{}, nil
}

func main() {
	db, _ := newDatabase("test.db")
	defer db.Close()

	rows, _ := db.Exec(`SELECT * FROM users WHERE username = ?`)
	fmt.Printf("%v\n", rows)
}
