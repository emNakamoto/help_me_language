package main

import (
	"bufio"
	"database/sql"
	"os"

	"github.com/emNakamoto/help_me_language/secrets"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

/*
 * Dictionary... - a very simple struct
 */
type Word struct {
	Language string `json:"language"`
	Word     string `json:"word"`
	Meaning  string `json:"meaning"`
	Tag      string `json:"tag"`
}


func main() {
	db, err := connectDb(secrets.GetUser(), secrets.GetPassword())
	// Execute the query
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	printBlue("==========================================")
	printBlue("\t\tDICTIONARY							")
	printBlue("==========================================")
	printBlue("Welcome! Please choose a language.")
	print("1. English")
	print("2. Japanese")
	print("3. French")
	print("4. Korean")

	for { // main loop
		reader := bufio.NewReader(os.Stdin)
		char, _, err := reader.ReadRune()
		switch char {
		case '1':
			print("1 pressed")
			query_dict(db, "english", "", "")
		default:
			print("1 not pressed")
		}

		if err != nil {
			printRed(err.Error())
			printRed("Exiting.")
			break
		}
	}
	printBlue("Choose an option:")
	print("1. Browse")
	print("2. Add word")
	print("3. Flashcard")
	print("4. Quiz")

}

