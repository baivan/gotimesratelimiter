package main

import (
	"fmt"
	"sync"
	"time"
)

type RequestsCount struct {
	Lock  sync.Mutex
	Count int
}

func (r *RequestsCount) Incr() {
	r.Lock.Lock()
	r.Count += 1
	r.Lock.Unlock()
}
func (r *RequestsCount) Decr() int {
	r.Lock.Lock()
	r.Count -= 1
	// use a saturating sub version
	if r.Count < 0 {
		r.Count = 0
	}
	value := r.Count
	r.Lock.Unlock()
	return value
}

func (r *RequestsCount) Reset() {
	r.Lock.Lock()
	r.Count = 290 //this should be defined in a config file somewhere
	r.Lock.Unlock()
}
func (r *RequestsCount) isZero() bool {
	r.Lock.Lock()
	value := r.Count == 0
	r.Lock.Unlock()
	return value
}
func initCounter() RequestsCount {
	return RequestsCount{
		Lock:  sync.Mutex{},
		Count: 0,
	}
}

func resetCount(counter *RequestsCount) {
	timer := time.NewTicker(time.Second * 2) //the 2 seconds to tick,
	for range timer.C {
		fmt.Println("+++++++++++++ticker ticked  ")
		counter.Reset()
	}
}

//simulates processign a request like sending
func sendRequest(counter *RequestsCount) {

	// we have sent one request so decrement
	for {
		if counter.isZero() {
			continue
		} else {
			value := counter.Decr()
			fmt.Printf("Counter value %d\n", value)
			//time.Sleep(30 * time.Millisecond)
		}
	}

}

//the solution is to limit sending requests to
//290 per 2 seconds

func main() {
	counter := initCounter()
	done := make(chan bool)
	go resetCount(&counter)
	go sendRequest(&counter)

	<-done
}
