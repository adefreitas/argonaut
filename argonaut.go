package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello there! Starting processing")
	manifest := readManifest()
	var generator AssetConfigGenerator
	generator.init(4, manifest)

	// fmt.Println(generator.generate())
	// generator.generate()
	var wg sync.WaitGroup
	paralelization := 2
	wg.Add(paralelization)
	c := make(chan int)
	lo, hi := 0, 4
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
}
