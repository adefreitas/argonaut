package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readManifest() {
	fmt.Println("Reading manifest")
	dat, err := os.Open(INPUT_MANIFEST_DIR + "/manifest.json")

	if err != nil {
		fmt.Print("Error reading file", err)
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(dat)

	var attributeManifests []AttributeManifest

	fmt.Println(string(byteValue))
	json.Unmarshal(byteValue, &attributeManifests)
	fmt.Println("Unmarshalled")
	for i := 0; i < len(attributeManifests); i++ {
		fmt.Println("Attribute: " + attributeManifests[i].Attribute)
		for j := 0; j < len(attributeManifests[i].Categories); j++ {
			fmt.Println("Attribute category name: " + attributeManifests[i].Categories[j].Name)
			fmt.Println("Attribute category rarity: ", attributeManifests[i].Categories[j].Rarity)
		}
	}
	// fmt.Println(attributeManifests)
}
