//Более сложный пример, с использованием пула обработчиков для типовых задач
package main

import (
	"fmt"
	"sync"
)

//Task - описание интерфейса работы
type Task interface {
	Execute()
}



//Pool - структура, нам потребуется Мутекс, для гарантий атомарности изменений самого объекта
//Канал входящих задач
//Канал отмены, для завершения работы
//WaitGroup для контроля завершения работ
type Pool struct {
	mu    sync.Mutex
	size  int
	tasks chan Task
	kill  chan struct{}
	wg    sync.WaitGroup
}

func (p *Pool) Resize(size int) {
	//Захватываем лок, чтобы избежать одновременного изменения состояния
	p.mu.Lock()
	defer p.mu.Unlock()
	for p.size < size {
		p.size++
		p.wg.Add(1)
		go p.worker()
	}
	for p.size > size {
		p.size--
		p.kill <- struct{}{}
	}
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for true {
		select {
		//Если есть задача, то ее нужно обработать
		//Блокируется пока канал не будет закрыт или не поступит новая задача
		case task, ok := <-p.tasks:
			if !ok {
				return
			}
			task.Execute()
		//Если пришел сигнал умирать - выходим
		case <-p.kill:
			return

		}
	}

}

//Скроем внутренне устройство за конструктором, пользователь может влиять только на размер пула
func NewPool(size int) *Pool {
	pool := &Pool{
		//Канал задач - буферизованный, чтобы основная программа не блочилась при постановке задач
		tasks: make(chan Task, 128),
		//Канал для убийства "лишних воркеров"
		kill: make(chan struct{}),
	}
	pool.Resize(size)
	return pool
}

func (p *Pool) Close() {
	close(p.tasks)
}

func (p *Pool) Wait() {
	p.wg.Wait()
}

func (p *Pool) Exec(task Task) {
	p.tasks <- task
}

type ExampleTask string

func (e ExampleTask) Execute() {
	fmt.Println("executing: ", string(e))
}

func main() {
	pool := NewPool(5)
	pool.Exec(ExampleTask("foo"))
	pool.Exec(ExampleTask("bar"))

	pool.Resize(3)
	pool.Resize(6)

	for i := 0; i < 20; i++ {
		pool.Exec(ExampleTask(fmt.Sprintf("additional_%d", i+1)))
	}

	pool.Close()
	pool.Wait()



	//Почитать про sync.Pool{}, использование для снятия нагрузке с gc
}
