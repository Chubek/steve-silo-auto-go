package main

import (
	"database/sql"
	"errors"
	"io/fs"
	"log"
	"os"
	"time"

	"github.com/lazureykis/dotenv"
	_ "github.com/mattn/go-sqlite3"
)

func connectToDB() (tx *sql.Tx) {
	dotenv.Go()

	if _, err := os.Stat(os.Getenv("DATABASE_LOC")); errors.Is(err, fs.ErrNotExist) {
		if err != nil {
			db, err := sql.Open("sqlite3", os.Getenv("DATABASE_LOC"))
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			sqlStmt := `
			create table m2web (id integer not null primary key autoincrement, name text, value text, quality text, date text);
			create table chipbins (id integer not null primary key autoincrement, bin_one text, bin_two text, bin_three text, bin_four text, bin_five text, date text);
			create table mybinview_tanks (id integer not null primary key autoincrement, tank_name text, location_name text, state text, percent text, last_read text, measurement text, date text);
			create table mybinview_definitions (id integer not null primary key autoincrement, tank_id integer not null, type text, state text, last_date text, measurement text, date text)
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

	return tx

}

func InsertIntoDb(m2wObj Server, chipbObj ChipBinsOCRRes, binViewObj MeasurementData) {
	tx := connectToDB()

	stmtM2Web, err := tx.Prepare("insert into m2web(name, value, quality, date) values(?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error with ststm2web", err)
	}
	defer stmtM2Web.Close()

	stmtChipBins, err := tx.Prepare("insert into chipbins(bin_one, bin_two, bin_three, bin_four, bin_five, date) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error with ststChipBins", err)
	}
	defer stmtChipBins.Close()
	stmtTanks, err := tx.Prepare("insert into mybinview_tanks(tank_name, location_name, state, percent, last_read, measurement, date) values(?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error with stmtTanks", err)
	}
	defer stmtTanks.Close()

	stmtDefs, err := tx.Prepare("insert into mybinview_definitions (tank_id, type, state, last_date, measurement, date) values(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal("Error with stmtDef", err)
	}
	defer stmtDefs.Close()

	for _, v := range m2wObj.Vars.Var {
		currentTime := time.Now()
		_, errmw := stmtM2Web.Exec(v.Text, v.Value, v.Quality, currentTime.String())
		log.Println("Inserted into M2Web.")
		if errmw != nil {
			log.Fatal(errmw)
		}
	}

	listOfBins := make([]string, 0)

	for _, l := range chipbObj.ParsedResults[0].TextOverlay.Lines {
		listOfBins = append(listOfBins, l.LineText)
	}

	currentTime := time.Now()
	_, errcb := stmtChipBins.Exec(listOfBins[0],
		listOfBins[1],
		listOfBins[2],
		listOfBins[3],
		listOfBins[4],
		currentTime.String())
	if errcb != nil {
		log.Fatal(errcb)
	}
	log.Println("Inserted into ChipBins.")

	for _, m := range binViewObj.Model {
		currentTime := time.Now()
		sqlRes, errt := stmtTanks.Exec(m.TankName,
			m.LocationName,
			m.State,
			m.Percent,
			m.LatestReadingDate,
			m.Measurement,
			currentTime.String())

		if errt != nil {
			log.Fatal("Error with tanks", errt)
		}
		log.Println("Inserted into Tanks.")
		lastId, err := sqlRes.LastInsertId()
		if err != nil {
			log.Fatal(err)
		}
		for _, d := range m.Definitions {
			currentTime := time.Now()
			_, errd := stmtDefs.Exec(lastId,
				d.Type,
				d.State,
				d.LatestReadingDate,
				d.Measurement,
				currentTime.String())
			if errd != nil {
				log.Fatal("Error with defs", errd)
			}
			log.Println("Inserted into Definitions.")

		}

	}

	tx.Commit()
}
