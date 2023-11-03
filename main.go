package main

import (
	counter "1771/pkg/counterStruct"
	"context"
	"fmt"
	"log"
	"sync"
)

func worker(ctx context.Context, cancel context.CancelFunc, c *counter.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		c.Add(1, ctx, cancel)
		select {
		case <-ctx.Done():
			return
		default:

		}
	}
}

func main() {
	var wg sync.WaitGroup
	var wgChan sync.WaitGroup
	ctx, cancel := context.WithCancel(context.Background())

	amountOfThreads := 0
	maxValue := 0

	fmt.Println("Укажите количество горутин: ")
	_, err := fmt.Scanln(&amountOfThreads)
	if err != nil {
		log.Fatalln("Неверное значение")
	}

	if amountOfThreads < 1 {
		log.Fatalln("Количество горутин не может быть меньше 1")
	}

	fmt.Println("Укажите максимальное значение счетчика: ")
	_, err = fmt.Scanln(&maxValue)
	if err != nil {
		log.Fatalln("Неверное значение")
	}

	if maxValue < 1 {
		log.Fatalln("Максимальное значение счетчика не может быть меньше 1")
	}

	c := counter.NewCounter(maxValue)

	for id := 0; id < amountOfThreads; id++ {
		wg.Add(1)
		go worker(ctx, cancel, c, &wg)
	}
	go c.Increment(&wgChan, cancel)
	wgChan.Add(1)
	wg.Wait()
	c.CloseChannel()
	wgChan.Wait()

	// Печатаем значение счетчика
	fmt.Println("Counter:", c.Value())
	//Закрываем канал
}
