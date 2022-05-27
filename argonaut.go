package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	t1 := time.Now()
	fmt.Println("Hello there! Starting processing")
	manifest := readManifest()
	var generator AssetConfigGenerator
	var total int = 10
	generator.init(int16(total), manifest)

	// fmt.Println(generator.generate())
	// generator.generate()
	var wg sync.WaitGroup
	paralelization := 3
	wg.Add(paralelization)
	c := make(chan int)
	lo, hi := 0, total
	// Creating an array from 0 to 200 for paralelization
	assetNumbers := make([]int, hi-lo+1)
	for i := range assetNumbers {
		assetNumbers[i] = i + lo
	}
	// List creating ends

	for i := 0; i < paralelization; i++ {
		go func(c chan int) {
			for {
				v, more := <-c
				if more == false {
					wg.Done()
					return
				}
				work(&generator, v)
			}
		}(c)
	}

	// Adding asset numbers to the channel to be consumed by the loop above
	for _, a := range assetNumbers {
		c <- a
	}
	// closing channel
	close(c)
	fmt.Println("Waiting for paralel asset generation to finish")
	wg.Wait()
	fmt.Println("Paralel asset generation done!")
	t2 := time.Now()

	fmt.Println("Generation started at", t1)
	fmt.Println("Generation ended at", t2)

	diff := t2.Sub(t1)
	fmt.Println("Generation took", diff)
}
