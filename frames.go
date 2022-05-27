package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"os"
)

func extractFrame(framePath string) image.Image {
	// Extrated from https://stackoverflow.com/a/49595770
	existingImageFile, err := os.Open(framePath)
	if err != nil {
		fmt.Printf("Error opening image %w\n"+framePath, err)
	}

	defer existingImageFile.Close()

	imageData, imageType, err := image.Decode(existingImageFile)
	if err != nil {
		fmt.Printf("Error decoding image %w\n"+framePath, err)
	}
	fmt.Println(imageData)
	fmt.Println(imageType)

	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		fmt.Printf("Error decoding png image %w\n"+framePath, err)
	}
	fmt.Println(loadedImage)
	return loadedImage
}

func extractBackground() image.Image {
	return extractFrame(INPUT_DIR + "/background.png")
}

func combinteAttributesForFrame(frames Frames, prefix int16, frameNumber int16) {
	generatedFrames := make([]image.Image, 0)
	generatedFrames = append(generatedFrames, extractFrame(frames.aura[frameNumber]))
	generatedFrames = append(generatedFrames, extractFrame(frames.blips[frameNumber]))
	generatedFrames = append(generatedFrames, extractFrame(frames.gems[frameNumber]))
	generatedFrames = append(generatedFrames, extractFrame(frames.hands[frameNumber]))
	generatedFrames = append(generatedFrames, extractFrame(frames.stairs[frameNumber]))
	generatedFrames = append(generatedFrames, extractFrame(frames.watchers[frameNumber]))

	bgImage := image.NewRGBA(image.Rect(0, 0, 1080, 1920))

	for _, img := range generatedFrames {
		draw.Draw(bgImage, img.Bounds(), img, image.ZP, draw.Over)
	}

	path := fmt.Sprintf("%s/raw/%d/%d_%d.png", OUTPUT_FRAMES_DIR, prefix, prefix, frameNumber)
	out, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating image file: %s\n", path)
	}

	err = png.Encode(out, bgImage)
	if err != nil {
		fmt.Printf("Error creating image file: %+v\n", err)
		return
	}

}
