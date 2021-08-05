package multithreading

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

type FizzBuzz struct {
	n          int
	numCh      chan struct{}
	fizzCh     chan int
	buzzCh     chan int
	fizzBuzzCh chan int
	q          chan struct{}
}

// n%3 == 0
func (f *FizzBuzz) fizz() {
	defer wg.Done()
	for {
		select {
		case <-f.fizzCh:
			fmt.Println("fizz")
			f.numCh <- struct{}{}
		case <-f.q:
			return
		}
	}
}

// n%5 == 0
func (f *FizzBuzz) buzz() {
	defer wg.Done()
	for {
		select {
		case <-f.buzzCh:
			fmt.Println("buzz")
			f.numCh <- struct{}{}
		case <-f.q:
			return
		}
	}
}

// n%3 == 0 && && n%5 == 0
func (f *FizzBuzz) fizzbuzz() {
	defer wg.Done()
	for {
		select {
		case <-f.fizzBuzzCh:
			fmt.Println("fizzbuzz")
			f.numCh <- struct{}{}
		case <-f.q:
			return
		}
	}
}

func (f *FizzBuzz) number() {
	defer wg.Done()

	for i := 1; i <= f.n; i++ {
		<-f.numCh
		switch {
		case i%3 == 0 && i%5 == 0:
			f.fizzBuzzCh <- 1
		case i%3 == 0:
			f.fizzCh <- 1
		case i%5 == 0:
			f.buzzCh <- 1
		default:
			fmt.Println(i)
			f.numCh <- struct{}{}
		}
	}

	f.q <- struct{}{}
	f.q <- struct{}{}
	f.q <- struct{}{}
}

func NewFizzBuzz(n int) *FizzBuzz {
	return &FizzBuzz{
		n:          n,
		numCh:      make(chan struct{}, 1), //多一个缓冲，避免下轮循环死锁
		fizzCh:     make(chan int),
		buzzCh:     make(chan int),
		fizzBuzzCh: make(chan int),
		q:          make(chan struct{}, 1), //用来通知另外三个线程退出，避免泄露
	}
}

func main() {
	fb := NewFizzBuzz(15)
	wg.Add(4)
	go fb.fizz()
	go fb.buzz()
	go fb.fizzbuzz()
	go fb.number()
	fb.numCh <- struct{}{}

	wg.Wait()
}
