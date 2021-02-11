package controller_base

import (
	"fmt"
	"sync"
)

func Controller() {
	processing_input1 := make(chan string)
	processing_input2 := make(chan string)
	processing_routing := make(chan string)
	routing_output1 := make(chan string)
	routing_output2 := make(chan string)

	var wg sync.WaitGroup
	wg.Add(6)
	go Input1(&wg, processing_input1)
	go Input2(&wg, processing_input2)
	go Output1(&wg, routing_output1)
	go Output2(&wg, routing_output2)
	go Processing(&wg, processing_input1, processing_input2, processing_routing)
	go Routing(&wg, routing_output1, routing_output2, processing_routing)
	wg.Wait()
}

func Input1(wg *sync.WaitGroup, processing_chan chan string) {
	defer wg.Done()
	processing_chan <- "input1"
t0:
	for {
		processing_chan <- "input1"
		continue t0
	}
}

func Input2(wg *sync.WaitGroup, processing_chan chan string) {
	defer wg.Done()
	processing_chan <- "input2"
t0:
	for {
		processing_chan <- "input2"
		continue t0
	}
}

func Output1(wg *sync.WaitGroup, routing_chan chan string) {
	defer wg.Done()
	<-routing_chan
t0:
	for {
		<-routing_chan
		continue t0
	}
}

func Output2(wg *sync.WaitGroup, routing_chan chan string) {
	defer wg.Done()
	<-routing_chan
t0:
	for {
		<-routing_chan
		continue t0
	}
}

func Processing(wg *sync.WaitGroup, input1_chan, input2_chan, routing_chan chan string) {
	defer wg.Done()
	select {
	case msg := <-input1_chan:
		fmt.Println("processing: packet from", msg)
	t0:
		for {
			routing_chan <- msg
			fmt.Println("processing: packet to routing")
			select {
			case msg := <-input1_chan:
				fmt.Println("processing: packet from:", msg)
				continue t0
			case msg := <-input2_chan:
				fmt.Println("processing: packet from:", msg)
				continue t0
			}
		}
	case msg := <-input2_chan:
		fmt.Println("processing: packet from:", msg)
	t1:
		for {
			routing_chan <- msg
			fmt.Println("processing: packet to routing")
			select {
			case msg := <-input1_chan:
				fmt.Println("processing: packet from:", msg)
				continue t1
			case msg := <-input2_chan:
				fmt.Println("processing: packet from:", msg)
				continue t1
			}
		}
	}
}

func Routing(wg *sync.WaitGroup, output1_chan, output2_chan, processing_chan chan string) {
	defer wg.Done()
	msg := <-processing_chan
	fmt.Println("routing: msg from processing -", msg)
t0:
	for {
		select {
		case output1_chan <- msg:
			fmt.Println("routing: msg to output1")
			msg = <-processing_chan
			fmt.Println("routing: msg from processing -", msg)
			continue t0
		case output2_chan <- msg:
			fmt.Println("routing: msg to output1")
			msg = <-processing_chan
			fmt.Println("routing: msg from processing -", msg)
			continue t0
		}
	}
}
