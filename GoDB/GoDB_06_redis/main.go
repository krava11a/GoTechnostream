package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"log"
	"time"
)

var (
	c redis.Conn
)

func PanicOnErr(err error) {
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

func getRecord(mkey string) string {
	println("get", mkey)
	//get row with //https://redis.io/commands/get
	data, err := c.Do("GET", mkey)
	item, err := redis.String(data, err)

	if err == redis.ErrNil {
		fmt.Println("Record not found in redis (return value is nil)")
		return ""
	} else if err != nil {
		PanicOnErr(err)
	}
	return item

}

func main() {

	var err error
	c, err = redis.DialURL("redis://user:@localhost:6379/0")
	PanicOnErr(err)
	defer c.Close()

	mkey := "record_33"

	item := getRecord(mkey)

	fmt.Println(item)
	fmt.Printf("first get %+v\n", item)

	ttl := 5
	//add row
	result, err := redis.String(c.Do("SET", mkey, 1, "EX", ttl))
	PanicOnErr(err)
	if result != "OK" {
		panic("result not ok: " + result)
	}
	time.Sleep(time.Microsecond)
	//time.Sleep(7 * time.Second)

	item = getRecord(mkey)
	fmt.Printf("second get %+v\n", item)

	//https://redis.io/commands/incr/ - хороошо описан рейтлимитер
	//https://redis.io/commands/incrby
	n, _ := redis.Int(c.Do("INCRBY", mkey, 2))
	fmt.Println("INCRBY by 2", mkey, "is", n)

	//https://redis.io/commands/decrby/
	n, _ = redis.Int(c.Do("DECRBY", mkey, 1))
	fmt.Println("DECRY by 1 ", mkey, "is", n)

	//если записи не было - редис создаст
	n, err = redis.Int(c.Do("INCR", mkey+"_not_exist"))
	fmt.Println("INCR (default by 1) ", mkey+"_not_exist", "is", n)

	PanicOnErr(err)

	keys := []interface{}{mkey, mkey + "_not_exist", "sure_not_exist"}

	reply, err := redis.Strings(c.Do("MGET", keys...))
	PanicOnErr(err)
	fmt.Println(reply)

}
