package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func produce(queue chan int, wg sync.WaitGroup) {
	fmt.Println("begin to produce")
	timer := time.NewTimer(time.Second)
	for  {
		select {
			case <-timer.C:
				data := rand.Intn(10)
				queue <- data
				fmt.Printf("send data  %d to queue\n", data)
				timer.Reset(time.Second)
		}
	}
	wg.Done()
}

func consume(queue chan int, wg sync.WaitGroup, name string) {
	fmt.Println("begin to consume")
	for i:= 0; i < 10;i++{
		select {
		case	data := <-queue:
			fmt.Printf("%s read data %d from queue\n",name, data)
		}
	}
	fmt.Println("consumer begin to sleep")
	time.Sleep(time.Second*10)
	fmt.Println("consumer resumed")
	for {
		select {
		case	data := <-queue:
			fmt.Printf("read data %d from queue\n", data)
		}
	}
	wg.Done()
}

func main(){
	queue := make(chan int,10)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go produce(queue,wg)
	wg.Add(1)
	go consume(queue,wg,"a")
	wg.Add(1)
	go consume(queue,wg,"b")
	wg.Wait()
}