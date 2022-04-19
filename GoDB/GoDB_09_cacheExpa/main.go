package main

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"reflect"
	"strconv"
	"time"
)

var (
	c  redis.Conn
	db *gorm.DB
)

type CachedItem struct {
	Data interface{}
	Tags map[string]int
}

type Article struct {
	Name string
	From string
}

type Articles []Article

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
	if u.Score > 10 {
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

func PrintByKeyFromCache(mkey string) {
	fio, err := getCachedRecord(mkey)
	if err == redis.ErrNil {
		fmt.Println("Record not found in cahce")
	} else if err != nil {
		log.Fatal(err)
	}
	fmt.Println(" From cache table. Fio is ", fio)
}

func GetFioFromDB(id uint) (string, error) {
	st := Student{}
	err := db.Find(&st, id).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("record not found", id)
		return "", nil
	} else {
		PanicOnErr(err)
	}
	return st.Fio, nil
}

func getCachedRecord(mkey string) (string, error) {
	println("get", mkey)
	//get row with //https://redis.io/commands/get
	data, err := c.Do("GET", mkey)
	item, err := redis.String(data, err)

	if err == redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		return "", redis.ErrNil
	} else if err != nil {
		PanicOnErr(err)
	}
	return item, nil

}

func main() {

	var err error
	ttl := 5

	c, err = redis.DialURL("redis://user:@localhost:6379/0")
	PanicOnErr(err)
	defer c.Close()

	var top = Articles{
		Article{
			Name: "Как взрываются базовые станции",
			From: "geektimes",
		},
		Article{
			Name: "Java and Docker. It's must have",
			From: "habr",
		},
	}

	item := CachedItem{
		Data: top,
		Tags: map[string]int{
			"habr": 1,
			"GT":   1,
		},
	}
	jsonMarshal, _ := json.Marshal(item)
	println(string(jsonMarshal))
	keyNews := "top_news"
	result, err := redis.String(c.Do("SET", keyNews, jsonMarshal))
	if result != "OK" {
		panic("result not ok: " + result)
	}

	c.Do("SET", "habr", 1)
	c.Do("SET", "GT", 1)
	c.Do("INCR","habr")


	topNewsFromCache, err := getCachedRecord(keyNews)

	cItems := CachedItem{}
	json.Unmarshal([]byte(topNewsFromCache), &cItems)

	fmt.Printf("FROM CACHE: \n %s \n", topNewsFromCache)
	fmt.Printf("FROM VARIABLE AFTER UNMARSHALLING: \n %+v \n", cItems)
	fmt.Printf("FROM VARIABLE TAGS: \n %+v \n", cItems.Tags)

	keys := make([]interface{},0)
	toCompare := make([]int,0)
	for key, val := range cItems.Tags{
		keys = append(keys,key)
		toCompare= append(toCompare,val)
	}

	fmt.Println(" keys \n",keys,"\n","toCompare\n",toCompare)

	reply, err := redis.Ints(c.Do("MGET", keys...))
	PanicOnErr(err)
	fmt.Println(reply)

	if reflect.DeepEqual(toCompare,reply){
		println("cahce valid")
	}else {
		println("cache not valid")
	}

	return

	db, err = gorm.Open("mysql", "scot:tiger@tcp(localhost:3306)/university?charset=utf8")
	PanicOnErr(err)
	defer db.Close()
	//connect
	db.DB()
	db.DB().Ping()

	userID := 4
	mkey := "top_user_" + strconv.Itoa(userID)
	fio, err := getCachedRecord(mkey)
	if err == redis.ErrNil {
		println("Create cache data")
		result, err := redis.String(c.Do("SET", mkey+"_lock", fio, "EX", 3, "NX"))
		if result != "OK" {
			fmt.Println("somebody already build this cahce")
			for i := 0; i < 3; i++ {
				println("sleep ", i)
				time.Sleep(time.Second)
			}
		}

		fio, err = GetFioFromDB(uint(userID))
		result, err = redis.String(c.Do("SET", mkey, fio, "EX", ttl))
		PanicOnErr(err)
		if result != "OK" {
			panic("result not ok: " + result)
		}
		fmt.Println("Cahced data is created. Fio is " + fio)

		n, err := redis.Int(c.Do("DEL", mkey+"_lock"))
		PanicOnErr(err)
		println("lock deleted:", n)
	}

	PrintByKeyFromCache(mkey)
}
