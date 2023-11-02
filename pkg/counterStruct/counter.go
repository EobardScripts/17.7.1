package counterStruct

import (
	"fmt"
	"sync"
)

type Counter struct {
	value    int
	maxValue int
	ch       chan int
	mutex    sync.RWMutex
	wg       sync.WaitGroup
}

func NewCounter(maxValue int) *Counter {
	return &Counter{
		value:    0,
		maxValue: maxValue,
		ch:       make(chan int, maxValue),
		mutex:    sync.RWMutex{},
		wg:       sync.WaitGroup{},
	}
}

func (c *Counter) GetValue() int {
	return c.value
}

func (c *Counter) Wait() {
	c.wg.Wait()
}

func (c *Counter) Add(amount int) {
	c.wg.Add(amount)
}

func (c *Counter) Increment() {
	defer c.wg.Done()
	for {
		if c.isMax() {
			return
		}
		c.ch <- c.value
		c.mutex.Lock()
		c.value++
		fmt.Println(c.value)
		c.mutex.Unlock()
	}
}

func (c *Counter) isMax() bool {
	c.mutex.RLock()
	value := c.value
	c.mutex.RUnlock()
	if value >= c.maxValue {
		fmt.Println("is max")
		_, ok := <-c.ch
		if ok {
			c.mutex.Lock()
			c.Close()
			c.ch = nil
			c.mutex.Unlock()
		}
		return true
	}
	return false
}

func (c *Counter) Close() {
	close(c.ch)
}
