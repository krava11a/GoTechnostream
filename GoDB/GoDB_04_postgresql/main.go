package main

import (
	"database/sql"
	"fmt"
	//не используем импорт, однако в init() пакета зарегает нам драйвер для БД
	_ "github.com/lib/pq"
	"log"
)

var (
	db *sql.DB
)

func PrintByID(id int64) {
	var fio string
	var info sql.NullString
	// var info string
	var score int
	row := db.QueryRow("SELECT fio,info,score FROM students WHERE id = $1", id)
	err := row.Scan(&fio, &info, &score)
	PanicOnErr(err)
	fmt.Println("PrintByID:", id, "fio:", fio, "info:", info, "score:", score)
}

func PanicOnErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func main() {
	var err error
	//create structure db
	//but connection start only in first query
	//db, err = sql.Open("mysql", "scot:tiger@tcp(localhost:3306)/university?charset=utf8")
	//several parametres
	db, err = sql.Open("postgres", "user=postgres dbname=university password=example" +
		" host=localhost sslmode=disable")

	defer db.Close()
	//if necessary we can say how many connects are need
	//db.SetMaxOpenConns(10)

	fmt.Println("OpenConnections",db.Stats().OpenConnections)

	PanicOnErr(err)


	//check connection
	err = db.Ping()
	PanicOnErr(err)

	fmt.Println("OpenConnections",db.Stats().OpenConnections)

	rows, err := db.Query("SELECT fio,info,score from students")
	PanicOnErr(err)
	for rows.Next() {
		var fio string
		var info string
		var score string
		err := rows.Scan(&fio, &info, &score)
		PanicOnErr(err)
		fmt.Println("fio:", fio, "info:", info, "score:", score)
	}


	//закрывать обязательно, иначе потекут соединения
	rows.Close()


	//  &1... - placeholder
	var lastID int64
	err = db.QueryRow("INSERT INTO students(fio,info,score) values($1,$2,$3) RETURNING id",
		"KAS",
		"Manager",
		"7",
	).Scan(&lastID)
	PanicOnErr(err)
	fmt.Println("last id = ",lastID)

	PrintByID(lastID)


	//QueryRow не нужно закрывать row.Close(), внутри реализации уже есть Close()
	var fio string
	row := db.QueryRow("SELECT fio FROM students where id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("fio:", fio)

	//update
	result, err := db.Exec(
		"UPDATE students SET info = $1 WHERE id = $2",
		"test update user",
		lastID,
	)
	PanicOnErr(err)
	affected, err := result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update rowsAffected - ",affected)
	PrintByID(lastID)

	//prepared statements
	//prepare query
	stmt, err := db.Prepare("UPDATE students SET info = $1, score = $2 WHERE id = $3")
	PanicOnErr(err)
	// exec for prepare statement query
	result, err = stmt.Exec("prepared statements update", 150,lastID)
	PanicOnErr(err)
	PrintByID(lastID)

	return


}
