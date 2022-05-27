package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func work(generator *AssetConfigGenerator, index int) {
	generationData := generator.generate()
	outputFramesDirPath := fmt.Sprintf("%s/raw/%d", OUTPUT_FRAMES_DIR, index)
	outputVideoDirPath := fmt.Sprintf("%s/%d", OUTPUT_VIDEO_DIR, index)
	outputManifestPath := fmt.Sprintf("%s/%d.json", outputVideoDirPath, index)

	e0 := os.MkdirAll(OUTPUT_DIR, 0755)
	if e0 != nil {
		fmt.Println("error creating directory output", e0)
	}

	e1 := os.MkdirAll(fmt.Sprintf("%s", OUTPUT_FRAMES_DIR), 0755)
	if e1 != nil {
		fmt.Println("error creating directory frames", e1)
	}

	e2 := os.MkdirAll(outputFramesDirPath, 0755)
	if e2 != nil {
		fmt.Println("error creating directory frames", e2)
	}

	e3 := os.MkdirAll(outputVideoDirPath, 0755)
	if e3 != nil {
		fmt.Println("error creating directory video", e3)
	}

	file, err := json.MarshalIndent(generationData.data, " ", " ")

	if err != nil {
		fmt.Println("Couldnt mashal json")
	}

	err = ioutil.WriteFile(outputManifestPath, file, 777)

	if err != nil {
		fmt.Println("Couldnt write file", err)
	}

	combineAttributes(generationData.frames, index)
}
