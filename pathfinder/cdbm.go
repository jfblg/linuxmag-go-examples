package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"os/user"
	"path"

	"github.com/manifoldco/promptui"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	addMode := flag.Bool("add", false, "addition mode")
	flag.Parse()

	db, err := sql.Open("sqlite3", dbPath())
	panicOnErr(err)
	defer db.Close()

	_, err := os.Stat(dbPath())
	if os.IsNotExist(err) {
		create(db)
	}

	dir, err := os.Getwd()
	panicOnErr(err)

	if *addMode {
		dirInsert(db, dir)
	} else {
		items := dirList(db)
		prompt := promptui.Select{
			Label: "Pick a directory",
			Items: items,
		}

		_, result, err := prompt.Run()
		panicOnErr(err)

		fmt.Fprintf(os.Stderr, "%s\n", result)
	}
}

func dirList(db *sql.DB) []string {
	items := []string{}

	rows, err := db.Query(`SELECT dir FROM dirs ORDER BY date DESC LIMIT 10`)
	panicOnErr(err)

	usr, err := user.Current()
	panicOnErr(err)

	for rows.Next() {
		var dir string
		err = rows.Scan(&dir)
		panicOnErr(err)
		items = append(items, dir)
	}

	if len(items) == 0 {
		items = append(items, usr.HomeDir)
	} else if len(items) > 1 {
		items = items[1:] // skip first item i.e. users HOME
	}

	return items
}

func dirInsert(db *sql.DB, dir string) {
	stmt, err := db.Prepare(`REPLACE INTO dirs(dir, date) VALUES(?, datetime('now'))`)
	panicOnErr(err)

	_, err := stmt.Exec(dir)
	panicOnErr(err)
}

func create(db *sql.DB) {
	_, err := db.Exec(`CREATE TABLE dirs (dir text, date text)`)
	panicOnErr(err)

	_, err := db.Exec(`CREATE UNIQUE INDEX idx ON dirs (dir)`)
	panicOnErr(err)
}

func dbPath() {
	var dbFile = ".cdbm.db"

	usr, err := user.Current()
	panicOnErr(err)
	return path.Join(usr.HomeDir, dbFile)
}

func panicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}
