package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func connectDb(username, password string) (*sql.DB, error) {
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

func insert(db *sql.DB, table, values string) {
	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO " + table + " VALUES ( " + values + " )")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
}

func query_dict(db *sql.DB, lang, word, tag string) {
	
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

func get_languages(db *sql.DB, table string) []string {
	var lang_array []string
	query := "SELECT DISTINCT language FROM " + table
	
	results, err := db.Query(query)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	
    for results.Next() {
        var lang string
        // for each row, scan the result into our tag composite object
        err = results.Scan(&lang)
        if err != nil {
            panic(err.Error()) // proper error handling instead of panic in your app
        }
                // and then print out the tag's Name attribute
        print(lang)
		lang_array = append(lang_array, lang)
    }
	return lang_array
}