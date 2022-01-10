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
	insert.Close()
}

// function to update dictionary, given language and word
func update_dict(db *sql.DB, lang, word, meaning, tag string) {
	var update_string string
	// if tag is NULL, we won't need to add '' around it	

	if meaning == "" && tag == "" {
		return
	}
	if meaning != "" {
		update_string = "meaning = '"+ meaning + "'"
	}
	if tag != "" {
		// if we're only updating tag, we won't need to add a comma
		if update_string != "" {
			update_string = update_string + ", "
		}
		update_string = update_string + "tag = '"+ tag + "'"
	}

	query := "UPDATE dictionary SET " + update_string + "WHERE (language='" + lang + "' and word='" + word + "')"
	    //Exec executes a database query, but it does
	    // not return any row as result.
	_, err := db.Exec(query)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }	
}

// add to dictionary, if adding a new definition to a word that already exists, just update
func add_to_dict(db *sql.DB, lang, new_word, meaning, tag string) {
	word := Word{"", "", "", ""}

	query := "SELECT * FROM dictionary WHERE (language='" + lang + "' and word='" + new_word + "')"

	results, err := db.Query(query)
	if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
	for results.Next() {
        // there should only be one entry but this should ensure the row is valid
        err = results.Scan(&word.Language, &word.Word, &word.Meaning, &word.Tag)
        if err != nil {
            panic(err.Error()) 
        }
    }
	// if there is no entry yet
	if word.Word == "" {
		
		query = "INSERT INTO dictionary (language, word, meaning, tag) VALUES ('" + lang + "', '" + new_word + "', '" + meaning + "', '" + tag + "') "
		_, err = db.Exec(query)
		if err != nil {
        	panic(err.Error()) // proper error handling instead of panic in your app
    	}
	} else {
		// there should always be a meaning if we're adding to the dictionary
		if word.Meaning != "" {
			word.Meaning = word.Meaning + "\n\n.\n" + meaning
		}
		// if we're adding a new tag, just append it to the old one, null tags don't work with scan
		// default for tags is none if none provided. If both none, tag doesn't change
		if word.Tag == "none" && tag != "none"{
			word.Tag = tag
		} else if word.Tag != "none" && tag != "none" {
			word.Tag = word.Tag + ", " + tag
		}
		update_dict(db, word.Language, word.Word, word.Meaning, word.Tag)
	}
	print("new entry added :)")
}

func delete_entry(db *sql.DB, table,condition string) {
	query_string := "DELETE FROM " + table + " WHERE " + condition

	_, err := db.Exec(query_string)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }	
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
		print("-------------------------")
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