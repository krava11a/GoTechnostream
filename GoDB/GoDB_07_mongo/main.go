package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"math/rand"
)

var (
	sess *mgo.Session
)

type student struct {
	ID    bson.ObjectId `json:"id" bson:"_id"`
	Fio   string        `json:"fio" bson:"fio"`
	Info  string        `json:"info" bson:"info"`
	Score int           `json:"score"" bson: "score"`
}

func PanicOnErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func main() {
	var err error
	sess, err := mgo.Dial("mongodb://localhost:27017/")
	PanicOnErr(err)

	//Если коллекции не будет, то она создастся автоматически
	collection := sess.DB("university").C("students")
	index := mgo.Index{
		Key: []string{"fio"},
	}
	err = collection.EnsureIndex(index)
	PanicOnErr(err)

	if n, _ := collection.Count(); n == 0 {
		firstStudent := &student{ID: bson.NewObjectId(), Fio: "KAN", Info: "Speciala Agent", Score: 1550}
		err := collection.Insert(firstStudent)
		PanicOnErr(err)
	}

	var allStudents []student
	//bson.M{} - это типо условие для поиска
	err = collection.Find(bson.M{}).All(&allStudents)
	PanicOnErr(err)

	for i, v := range allStudents {
		fmt.Printf("student[%d]: %v\n", i, v)
	}

	//generate id
	id := bson.NewObjectId()
	//bson.M{"_id":id}
	var nonExistingStudent student
	err = collection.Find(bson.M{"_id": id}).One(&nonExistingStudent)
	if err == mgo.ErrNotFound {
		fmt.Println("Row not found!!!")
	} else if err != nil {
		PanicOnErr(err)
	}

	fmt.Println("Not exist - ", nonExistingStudent)

	secondStudent := &student{
		ID:    id,
		Fio:   fmt.Sprintf("Ivan %v", rand.Intn(100)),
		Info:  "",
		Score: rand.Intn(100),
	}
	//err = collection.Insert(secondStudent)
	//PanicOnErr(err)
	//
	//err = collection.Find(bson.M{"_id": id}).One(&nonExistingStudent)
	//PanicOnErr(err)
	//
	//fmt.Println("But now exist - ", nonExistingStudent)


	collection.UpdateAll(
		bson.M{"fio": "KAN"},
		bson.M{"$set": bson.M{"Info": "all ivans info!"}},
	)

	err = collection.Find(bson.M{"_id": id}).One(&nonExistingStudent)

	secondStudent.Info = "all records!!!!"
	collection.Update(bson.M{"_id":secondStudent.ID},&secondStudent)


	err = collection.Find(bson.M{}).All(&allStudents)
	PanicOnErr(err)

	for i, v := range allStudents {
		fmt.Printf("student[%d]: %v\n", i, v)
	}
}
