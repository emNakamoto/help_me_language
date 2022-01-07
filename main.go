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

func connectDb(username string, password string) (*sql.DB, error) {
	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
	var err error
	db, err = sql.Open("mysql", username+":"+password+"@tcp(127.0.0.1:3306)/language_database")
	_ = db

	// if there is an error opening the connection, handle it
	if err != nil {
		printRed("DB connection error")
		panic(err.Error())
		return nil, err
	}
	return db, nil
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

func query_dict(lang, word, tag string) {
	
	query := "SELECT * FROM dictionary WHERE (language='" + lang + "'"
	if word != "" {
		query = query + " and word='" + word + "'"
	}
	if tag != "" {
		query = query + " and tag='" + tag + "'"
	}
	query = query + ")" 
	results, err := db.Query(query)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	
    for results.Next() {
        var word Word
        // for each row, scan the result into our tag composite object
        err = results.Scan(&word.Language, &word.Word, &word.Meaning, &word.Tag)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        print(word.Word)
		print(word.Tag)
		print(word.Meaning)
    }
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
			query_dict("english", "", "")
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

