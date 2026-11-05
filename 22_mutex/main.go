package main

import (
	"fmt"
	"sync"
)

type post struct {
	views int
	mutex sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
	defer func() {
		p.mutex.Unlock()
		wg.Done()
	}()
	p.mutex.Lock()
	p.views++
}

func main() {
	var wg sync.WaitGroup
	my_post := post{
		views: 0,
	}

	var i int = 0

	for i = 0; i < 100; i++ {
		wg.Add(1)
		go my_post.inc(&wg)
	}
	wg.Wait()
	fmt.Println(my_post.views)
}
