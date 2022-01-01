package main

import (
    "bufio"
	"github.com/emNakamoto/help_me_language/ui"
    // _ "database/sql"
    // _ "github.com/go-sql-driver/mysql"
)

func main() {
        for { // main loop
        ui.PrintBlue("==========================================")
        ui.PrintBlue("\t\tDICTIONARY							")
        ui.PrintBlue("==========================================")
        ui.PrintBlue("Welcome! Please choose a language.")
        ui.Print("1. English")
        ui.Print("2. Japanese")
        ui.Print("3. French")
        ui.Print("4. Korean")

        reader := bufio.NewReader(os.Stdin)
        char, _, err := reader.ReadRune()

        if err != nil {
            ui.PrintRed(err)
            ui.PrintRed("Exiting.")
            break
        }

        ui.PrintBlue("Choose an option:")
        ui.Print("1. Browse")
        ui.Print("2. Add word")
        ui.Print("3. Flashcard")
        ui.Print("4. Quiz")
    } 

    // // Open up our database connection.
    // // I've set up a database on my local machine using phpmyadmin.
    // // The database is called testDb
    // db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")

    // // if there is an error opening the connection, handle it
    // if err != nil {
    //     panic(err.Error())
    // }

    // // defer the close till after the main function has finished
    // // executing
    // defer db.Close()

}