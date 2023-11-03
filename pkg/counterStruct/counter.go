package counter

type Counter struct {
	value  int
	limit  int
	txDone chan bool
}

func NewCounter(limit int) *Counter {
	c := Counter{
		limit:  limit,
		value:  0,
		txDone: make(chan bool, 1),
	}
	c.txDone <- true
	return &c
}

func (c *Counter) Add(amount int) bool {
	<-c.txDone

	if c.value >= c.limit {
		c.txDone <- true
		return false
	}

	c.value += amount
	c.txDone <- true
	return true
}

func (c *Counter) Value() int {
	<-c.txDone
	v := c.value
	c.txDone <- true
	return v
}
