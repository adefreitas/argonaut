package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readManifest() *NamedManifest {
	fmt.Println("Reading manifest")
	dat, err := os.Open(INPUT_MANIFEST_DIR + "/manifest.json")

	if err != nil {
		fmt.Print("Error reading file", err)
		panic(err)
	}

	byteValue, _ := ioutil.ReadAll(dat)

	var attributeManifests []AttributeManifest

	// fmt.Println(string(byteValue))
	json.Unmarshal(byteValue, &attributeManifests)

	var namedManifest NamedManifest

	for i := 0; i < len(attributeManifests); i++ {
		attributeManifest := attributeManifests[i]
		key := attributeManifest.Attribute
		switch key {
		case HAND_TOP_LEFT:
			namedManifest.handTopLeft = attributeManifest
		case HAND_TOP_RIGHT:
			namedManifest.handTopRight = attributeManifest
		case HAND_BOTTOM_LEFT:
			namedManifest.handBottomLeft = attributeManifest
		case HAND_BOTTOM_RIGHT:
			namedManifest.handBottomRight = attributeManifest
		case AURAS:
			namedManifest.auras = attributeManifest
		case BLIPS_AURA:
			namedManifest.blipAura = attributeManifest
		case ELEMENTS:
			namedManifest.elements = attributeManifest
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
		case MUSIC:
			namedManifest.music = attributeManifest
		default:
			fmt.Println("Unknown attribute name", key)
		}
	}
	fmt.Println("Hand bottom left manifest", &namedManifest.handBottomLeft)

	return &namedManifest
}
