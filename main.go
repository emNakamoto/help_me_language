package main

import (
	"bufio"
	"database/sql"
	"os"

	"strings"
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

func mod_dict(db *sql.DB, lang string) {
	print("What would you like to do?")
	print("1. add a new word/definition")
	print("2. edit existing entries")
	print("3. delete entries")
	
	// reading as a string because the rest of the input will be grabbed as string rather than rune
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	// remove trailing newline
    text = strings.Replace(text, "\n", "", -1)

	switch text {
	case "1":
		print("Enter word")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		word := strings.ToLower(text)
		if word == "" {
			print("invalid word entry")
			return
		}

		print("Enter definition (end with new line)")
		text, _ = reader.ReadString('\n')
		definition := strings.Replace(text, "\n", "", -1)
		
		print("Add a tag? (enter empty line for no)")
		text, _ = reader.ReadString('\n')
		text := strings.Replace(text, "\n", "", -1)
		var tag string
		if text == ""  {
			tag = "none"
		} else {
			tag = text
		}
		add_to_dict(db, lang, word, definition, tag)

	case "2":
		print("Edit which word?")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		edited_word := strings.ToLower(text)
		print("\n")
		print("-------------------------")
		print("current entry: (word, tag, definition)")
		query_dict(db, lang, text, "")
		print("-------------------------")
		print("\n")

		print("edit definition? (enter new definition or empty line for no)")
		text, _ = reader.ReadString('\n')
		new_meaning := strings.Replace(text, "\n", "", -1)

		print("edit tag? (enter new tag or empty line for no)")
		text, _ = reader.ReadString('\n')
		new_tag := strings.Replace(text, "\n", "", -1)

		update_dict(db, lang, edited_word, new_meaning, new_tag)
		print("entry edited")

	case "3":
		print("delete which word?")
		text, _ = reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		delete_string := "language='" + lang + "' and word='" + text + "'"
		delete_entry(db, "dictionary", delete_string)
		print("entry deleted")
		
	default:
		print("not a valid option")
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
			query_dict(db, "english", "", "")
		default:
			print("1 not pressed")
			mod_dict(db,"english")
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

