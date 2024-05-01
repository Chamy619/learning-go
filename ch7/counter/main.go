package main

import (
	"fmt"
	"time"
)

type Counter struct {
	total       int
	lastUpdated time.Time
}

func (c *Counter) Increment() {
	c.total++
	c.lastUpdated = time.Now()
}

func (c Counter) String() string {
	return fmt.Sprintf("total: %d, last updated: %v", c.total, c.lastUpdated)
}

func doUpdateWrong(c Counter) {
	c.Increment()
	fmt.Println("in doUpdateWrong:", c.String())
}

func doUpdateRight(c *Counter) {
	c.Increment()
	fmt.Println("in doUpdateRight:", c.String())
}

func main() {
	var c1 Counter
	fmt.Println(c1.String())
	c1.Increment()
	fmt.Println(c1.String())

	var c2 Counter
	doUpdateWrong(c2)
	fmt.Println("in main:", c2.String())
	doUpdateRight(&c2)
	fmt.Println("in main:", c2.String())
}
