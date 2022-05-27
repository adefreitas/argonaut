package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello there! Starting processing")
	manifest := readManifest()
	var generator AssetConfigGenerator
	generator.init(4, manifest)
	work(generator, 1)
	// fmt.Println(generator.generate())
	// generator.generate()
}
