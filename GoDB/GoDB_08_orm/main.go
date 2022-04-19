package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	db *gorm.DB
)

type Student struct {
	ID    uint `SQL:"AUTO_INCREMENT" gorm:"primary_key"`
	Fio   string
	Info  string
	Score int
}

func (u *Student) TableName() string {
	return "students"
}

func (u *Student) BeforeSave() (err error) {
	fmt.Println("Trigger on before save")
	if u.Score>10 {
		fmt.Println("Test passed")
	}
	return
}

func PanicOnErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func PrintById(id uint) {
	st := Student{}
	err := db.Find(&st, id).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("record not found", id)
	} else {
		PanicOnErr(err)
	}
	fmt.Printf("PrintById : %+v, data: %+v\n", id, st)
}

func main() {
	var err error
	db, err = gorm.Open("mysql", "scot:tiger@tcp(localhost:3306)/university?charset=utf8")
	PanicOnErr(err)
	defer db.Close()
	//connect
	db.DB()
	db.DB().Ping()

	//select by id
	PrintById(1)
	PrintById(100)

	//select all
	all := []Student{}
	db.Find(&all)
	for i, v := range all {
		fmt.Printf("students[%d] %+v\n", i, v)
	}

	//insert student
	newStudent := Student{
		Fio:   "DON DIMON",
	}
	db.Create(&newStudent)
	PrintById(newStudent.ID)

	//update student
	newStudent.Info = "occupation - programmer"
	newStudent.Score = 15
	db.Save(&newStudent)
	PrintById(newStudent.ID)


	return

}
