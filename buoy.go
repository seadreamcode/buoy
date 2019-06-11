package main

import "github.com/seadreamcode/buoy/db"

func newDatabase(filename string) (*db.Database, error) {
	return &db.Database{}, nil
}

func main() {
	db, _ := newDatabase("test.db")
	defer db.Close()

	db.Exec(`SELECT * FROM users WHERE username = ? LIMIT 1`)
}
