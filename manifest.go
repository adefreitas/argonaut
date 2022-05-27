package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readManifest() NamedManifest {
	fmt.Println("Reading manifest")
	dat, err := os.Open(INPUT_MANIFEST_DIR + "/manifest.json")

	if err != nil {
		fmt.Print("Error reading file", err)
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(dat)

	var attributeManifests []AttributeManifest

	json.Unmarshal(byteValue, &attributeManifests)

	var namedManifest NamedManifest

	for i := 0; i < len(attributeManifests); i++ {
		attributeManifest := attributeManifests[i]
		key := attributeManifest.Attribute
		switch key {
		case HANDS:
			namedManifest.hands = attributeManifest
		case AURA:
			namedManifest.aura = attributeManifest
		case WATCHERS:
			namedManifest.watchers = attributeManifest
		case STAIRS:
			namedManifest.stairs = attributeManifest
		case ARCHES:
			namedManifest.arches = attributeManifest
		case GEMS:
			namedManifest.gems = attributeManifest
		case BLIPS:
			namedManifest.blips = attributeManifest
		default:
			fmt.Println("Unknown attribute name", key)
		}
	}

	return namedManifest
}
