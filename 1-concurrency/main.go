package main

import (
	"fmt"
	"math"
	"math/rand"
)

type ss struct {
}

const RangeSlice = 10
const RandMAX = 100

func randSl(ch chan int) {
	sl := make([]int, RangeSlice)
	for range sl {
		ch <- rand.Intn(RandMAX)
	}
}
func powSl(ch chan int, chQ chan float64) {
	for i := 0; i < RangeSlice; i++ {
		res := <-ch
		res2 := math.Pow(float64(res), 2)
		chQ <- res2
	}
}

func main() {
	fmt.Println("1-concurrency")
	chSlice := make(chan int)
	chQ := make(chan float64)
	go randSl(chSlice)
	go powSl(chSlice, chQ)
	for i := 0; i < RangeSlice; i++ {
		res := <-chQ
		fmt.Printf(" %v", res)
	}
	close(chSlice)
	close(chQ)
}