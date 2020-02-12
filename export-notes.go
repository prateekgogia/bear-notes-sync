package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	if err := printNotes(); err != nil {
		log.Fatal("Error printing notes info ", err)
	}

}

var bearDB = "/Users/pgogia/workspace/golang/src/github.com/prateekgogia/bear-notes-sync/database.sqlite"

func printNotes() error {
	database, err := sql.Open("sqlite3", bearDB)
	if err != nil {
		return err
	}

	rows, err := database.Query("SELECT ZCREATIONDATE, ZMODIFICATIONDATE, ZTEXT, ZUNIQUEIDENTIFIER FROM `ZSFNOTE`")
	if err != nil {
		return fmt.Errorf("failed to create query err: %v", err)
	}
	defer rows.Close()

	var ZCREATIONDATE, ZMODIFICATIONDATE, ZTEXT, ZUNIQUEIDENTIFIER string
	for rows.Next() {
		rows.Scan(&ZCREATIONDATE, &ZMODIFICATIONDATE, &ZTEXT, &ZUNIQUEIDENTIFIER)
		fmt.Println(ZTEXT)
	}

	return nil
}
