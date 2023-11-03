package main

import (
	counter "1771/pkg/counterStruct"
	"fmt"
	"log"
	"sync"
)

func worker(c *counter.Counter, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		if ok := c.Add(1); !ok {
			break
		}
	}
}

func main() {
	var wg sync.WaitGroup

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

	wg.Add(amountOfThreads)
	for id := 0; id < amountOfThreads; id++ {
		go worker(c, &wg)
	}
	wg.Wait()

	// Печатаем значение счетчика
	fmt.Println("Counter:", c.Value())
	//Закрываем канал
	c.CloseChannel()
}
