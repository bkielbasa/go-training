package counter

type counter struct {
	count int
}

var v int

//go:noinline
func (c *counter) Inc() {
	c.count++
}

//go:noinline
func (c *counter) Dec() {
	c.count--
}

//go:noinline
func (c *counter) Val() int {
	return c.count
}

func increment(c *counter) {
	for i := 0; i < 1000; i++ {
		c.Inc()
	}
}

func decrement(c *counter) {
	for i := 0; i < 1000; i++ {
		c.Dec()
	}
}

func val(c *counter) {
	for i := 0; i < 1000; i++ {
		v += c.Val()
	}
}

//go:noinline
func StressTest() int {
	c := counter{}

	done := make(chan int, 3)
	go func() {
		decrement(&c)
		done <- 0
	}()

	go func() {
		val(&c)
		done <- 0
	}()

	go func() {
		increment(&c)
		done <- 0
	}()

	for i := 0; i < 3; i++ {
		<-done
	}

	return c.count
}
