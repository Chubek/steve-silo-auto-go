package main

import (
	"database/sql"
	"errors"
	"io/fs"
	"log"
	"os"

	"github.com/lazureykis/dotenv"
	_ "github.com/mattn/go-sqlite3"
)

var tx *sql.Tx

func connectToDB() {
	dotenv.Go()

	if _, err := os.Stat("/path/to/file"); errors.Is(err, fs.ErrNotExist) {
		if err != nil {
			db, err := sql.Open("sqlite3", os.Getenv("DATABASE_LOC"))
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			sqlStmt := `
			create table m2web (id integer not null primary key, name text, value text, quality text);
			create table chipbins (id integer not null primary key, bin_one text, bin_two text, bin_three text, bin_four text, bin_five text);
			create table mybinview_tanks (id integer not null primary key, tan_name text, location_name text, state text, percent text, last_read text, measurement text, full text);
			create table mybinview_definitions (id integer not null primary key, tank_id integer not null foreign key references mybinview_tanks(id), type text, state text, last_date text, measurement text)
		`
			_, err = db.Exec(sqlStmt)
			if err != nil {
				log.Printf("%q: %s\n", err, sqlStmt)
				return
			}
			tx, err = db.Begin()
			if err != nil {
				log.Fatal(err)
			}
		}
	} else {
		db, err := sql.Open("sqlite3", os.Getenv("DATABASE_LOC"))
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		tx, err = db.Begin()
		if err != nil {
			log.Fatal(err)
		}
	}

}
