package main

import (
	"log"
	"reflect"
	"sync"
)

func makeChan(chp interface{}, bufferSize int) {
	chpv := reflect.ValueOf(chp)
	if chpv.Kind() != reflect.Ptr {
		log.Panic("Первый аргумент должен быть ссылкой на канал!")
	}
	chv := chpv.Elem()
	if chv.Kind() != reflect.Chan {
		log.Panic("Первый аргумент должен быть ссылкой на канал!")
	}

	chantype := chv.Type()
	chv.Set(reflect.MakeChan(chantype, bufferSize))
}

func send(ch interface{}, value interface{}) {
	chv := reflect.ValueOf(ch)
	if chv.Kind() != reflect.Chan {
		log.Panic("first argument must be chan")
	}

	v := reflect.ValueOf(value)
	if chv.Type().Elem() != v.Type() {
		log.Panic("chan and value don't match type")
	}
	if chv.Type().ChanDir() == reflect.RecvDir {
		log.Panic("first argument must be sendable chan")
	}

	chv.Send(v)
}

//получение данных из произвольного канала
func recv(ch interface{}, p interface{}) bool {
	chv := reflect.ValueOf(ch)
	if chv.Kind() != reflect.Chan {
		log.Panic("first argument must be chan")
	}
	pv := reflect.ValueOf(p)
	if pv.Kind() != reflect.Ptr {
		log.Panic("second argument must be pointer")
	}

	if chv.Type().Elem() != pv.Type().Elem() {
		log.Panic("chan and value don't match type")
	}

	if chv.Type().ChanDir() == reflect.SendDir {
		log.Panic("first argument must be recievable chan")
	}

	v, ok := chv.Recv()
	pv.Elem().Set(v)

	return ok

}

func selectCase(recvCh interface{}, recvCase func(v interface{}, ok bool), sendCh interface{}, sendValue interface{},
	sendCase func(), defaultCase func()) {

	recvChv := reflect.ValueOf(recvCh)
	if recvChv.Kind() != reflect.Chan || recvChv.Type().ChanDir() == reflect.SendDir {
		log.Panic("first argument must be recievable chan")
	}

	sendChv := reflect.ValueOf(sendCh)
	if sendChv.Kind() != reflect.Chan || sendChv.Type().ChanDir() == reflect.RecvDir {
		log.Panic("third argument must be sendable chan")
	}

	v := reflect.ValueOf(sendValue)
	if sendChv.Type().Elem() != v.Type() {
		log.Panic("sendCh and sendValue don't match type")
	}

	if sendChv.Type().ChanDir() == reflect.RecvDir {
		log.Panic("first argument must be sendable chan")
	}

	cases := []reflect.SelectCase{
		reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: recvChv,
			Send: reflect.ValueOf(nil),
		},
		reflect.SelectCase{
			Dir:  reflect.SelectSend,
			Chan: sendChv,
			Send: v,
		},
		reflect.SelectCase{
			Dir:  reflect.SelectDefault,
			Chan: reflect.ValueOf(nil),
			Send: reflect.ValueOf(nil),
		},
	}

	chosen, recv, recvOK := reflect.Select(cases)
	switch chosen {
	case 0:
		recvCase(recv.Interface(), recvOK)
	case 1:
		sendCase()
	case 2:
		defaultCase()

	}

}

func main() {
	var ch chan int
	makeChan(&ch, 10)
	send(ch, 1)
	log.Printf("%d%d", len(ch), cap(ch))
	send(ch, 2)
	log.Printf("%d%d", len(ch), cap(ch))
	var n1 int
	recv(ch, &n1)
	log.Println(n1)

	ch2 := make(chan int)

	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go selectCase(
			//case recieve
			ch,
			func(v interface{}, ok bool) {
				log.Println(v, ok)
				wg.Done()
			},
			//case send
			ch2,
			10,
			func() {
				log.Println("send")
				wg.Done()
			},
			//default
			func() {
				log.Println("default", <-ch2)
				wg.Done()
			},
		)
	}
	wg.Wait()
}
