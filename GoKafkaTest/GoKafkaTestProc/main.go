package GoKafkaTestProc

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main()  {
	var wg sync.WaitGroup
	c := 0 //counter

	for {

		// создайм объект контекста с таймаутом в 15 секунд для чтения сообщений
		ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancel()

		// читаем очередное сообщение из очереди
		// поскольку вызов блокирующий - передаём контекст с таймаутом
		m, err := r.ReadMessage(ctx)
		if err != nil {
			fmt.Println("3")
			fmt.Println(err)
			break
		}

		wg.Add(1)
		// создайм объект контекста с таймаутом в 10 миллисекунд для каждой вычислительной горутины
		goCtx, goCcancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
		defer goCcancel()

		// вызываем функцию обработки сообщения (факторизации)
		go process(goCtx, c, &wg, m)
		c++
	}

	// ожидаем завершения всех горутин
	wg.Wait()
}
