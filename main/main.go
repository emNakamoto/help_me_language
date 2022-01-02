package main

import (
	"bufio"
	"database/sql"
	"os"

	"github.com/emNakamoto/help_me_language/ui"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
 * Dictionary... - a very simple struct
 */
type Word struct {
	Word     string `json:"word"`
	Meaning  string `json:"meaning"`
	Language string `json:"language"`
	Tag      string `json:"tag"`
}

func connectDb(username string, password string) (*sql.DB, error) {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/dict")
	_ = db

	// if there is an error opening the connection, handle it
	if err != nil {
		ui.PrintRed("DB connection error")
		panic(err.Error())
	}
}

func insert(table string, values string) {
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO " + table + " VALUES ( " + values + " )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func main() {
	connectDb("root", "tEqJtT)7GLJ8KfmA", db)
	// Execute the query
	results, err := db.Query("SELECT * FROM dictionary")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()
	for results.Next() {
		var word Word
		// for each row, scan the result into our tag composite object
		err = results.Scan(&word.Word, &word.Meaning, &word.Language, &word.Tag)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the tag's Name attribute
		ui.Print(word.Word)
	}

	ui.PrintBlue("==========================================")
	ui.PrintBlue("\t\tDICTIONARY							")
	ui.PrintBlue("==========================================")
	ui.PrintBlue("Welcome! Please choose a language.")
	ui.Print("1. English")
	ui.Print("2. Japanese")
	ui.Print("3. French")
	ui.Print("4. Korean")

	for { // main loop
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		switch char {
		case '1':
			ui.Print("1 pressed")
		default:
			ui.Print("1 not pressed")
		}

		if err != nil {
			ui.PrintRed(err.Error())
			ui.PrintRed("Exiting.")
			break
		}
	}
	ui.PrintBlue("Choose an option:")
	ui.Print("1. Browse")
	ui.Print("2. Add word")
	ui.Print("3. Flashcard")
	ui.Print("4. Quiz")

}
