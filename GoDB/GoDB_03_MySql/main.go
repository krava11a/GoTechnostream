package main

import (
	"database/sql"
	"fmt"
	//не используем импорт, однако в init() пакета зарегает нам драйвер для БД
	_ "github.com/go-sql-driver/mysql"
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
	row := db.QueryRow("SELECT fio,info,score FROM students WHERE id = ?", id)
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
	db, err = sql.Open("mysql", "scot:tiger@tcp(localhost:3306)/university?charset=utf8" +
		"&interpolateParams=true")

	//if necessary we can say how many connects are need
	//db.SetMaxOpenConns(10)

	//fmt.Println("OpenConnections",db.Stats().OpenConnections)

	PanicOnErr(err)

	//check connection
	//err = db.Ping()
	//PanicOnErr(err)

	//fmt.Println("OpenConnections",db.Stats().OpenConnections)

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

	//  ? - placeholder
	result, err := db.Exec("INSERT INTO students(fio,info,score) values(?,?,?)", "KAS", "Manager", "7")
	PanicOnErr(err)
	affected, err := result.RowsAffected()
	PanicOnErr(err)
	lastID, err := result.LastInsertId()
	PanicOnErr(err)
	fmt.Println(affected)
	fmt.Println(lastID)

	PrintByID(1)

	//QueryRow не нужно закрывать row.Close(), внутри реализации уже есть Close()
	var fio string
	row := db.QueryRow("SELECT fio FROM students where id = 1")
	err = row.Scan(&fio)
	PanicOnErr(err)
	fmt.Println("fio:", fio)

	//update
	result, err = db.Exec(
		"UPDATE students SET info = ? WHERE id = ?",
		"test update user",
		lastID,
	)
	PanicOnErr(err)
	affected, err = result.RowsAffected()
	PanicOnErr(err)
	fmt.Println("Update rowsAffected - ",affected)
	PrintByID(lastID)

	//prepared statements
	//prepare query
	stmt, err := db.Prepare("UPDATE students SET info = ?, score = ? WHERE id = ?")
	PanicOnErr(err)
	// exec for prepare statement query
	result, err = stmt.Exec("prepared statements update", 150,lastID)
	PanicOnErr(err)
	PrintByID(lastID)


	return

}
