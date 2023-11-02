package main

import (
	"1771/pkg/counterStruct"
	"fmt"
	"log"
)

func main() {
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

	counter := counterStruct.NewCounter(maxValue)
	for i := 0; i < amountOfThreads; i++ {
		counter.Add(1)
		go counter.Increment()
	}

	counter.Wait()
}
