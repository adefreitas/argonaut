package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"os"
	"sync"

	ffmpeg "github.com/adefreitas/go-fluent-ffmpeg"
)

func extractFrame(framePath string) image.Image {
	// Extrated from https://stackoverflow.com/a/49595770
	existingImageFile, err := os.Open(framePath)
	if err != nil {
		fmt.Printf("Error opening image %w\n"+framePath, err)
	}

	defer existingImageFile.Close()

	_, _, decodingError := image.Decode(existingImageFile)
	if decodingError != nil {
		fmt.Printf("Error decoding image %w\n"+framePath, decodingError)
	}
	// fmt.Println(imageData)
	// fmt.Println(imageType)

	existingImageFile.Seek(0, 0)

	loadedImage, err := png.Decode(existingImageFile)
	if err != nil {
		fmt.Printf("Error decoding png image %w\n"+framePath, err)
	}
	// fmt.Println(loadedImage)
	return loadedImage
}

func extractBackground() image.Image {
	return extractFrame(INPUT_DIR + "/background.png")
}

func combineAttributesForFrame(frames Frames, prefix int, frameNumber int, wg *sync.WaitGroup) {
	// defer wg.Done()
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

	path := fmt.Sprintf("%s/raw/%d/%d_%d.jpeg", OUTPUT_FRAMES_DIR, prefix, prefix, frameNumber)
	out, err := os.Create(path)
	if err != nil {
		fmt.Printf("Error creating image file: %s\n", path)
	}

	err = jpeg.Encode(out, bgImage, &jpeg.Options{Quality: 60})
	if err != nil {
		fmt.Printf("Error creating image file: %+v\n", err)
		return
	}
}

func combineAttributes(frames Frames, prefix int) {
	fmt.Println("Generating frames for asset", prefix)
	var wg sync.WaitGroup
	paralelization := 20
	wg.Add(paralelization)
	c := make(chan int)
	lo, hi := 0, 199
	// Creating an array from 0 to 200 for paralelization
	frameNumbers := make([]int, hi-lo+1)
	for i := range frameNumbers {
		frameNumbers[i] = i + lo
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
				combineAttributesForFrame(frames, prefix, v, &wg)
			}
		}(c)
	}
	// for i = 0; i < 200; i++ {
	// 	go combineAttributesForFrame(frames, prefix, i, &wg)
	// }

	// Adding frame numbers to the channel to be consumed by the loop above
	for _, a := range frameNumbers {
		c <- a
	}
	// closing channel
	close(c)
	fmt.Println("Waiting for paralel frame creation to finish for asset", prefix)
	wg.Wait()
	fmt.Println("Paralel frame creation done for asset", prefix)
	fmt.Println("Generating video", prefix)
	fileExtension := "%01d.jpeg"
	framesInputPath := fmt.Sprintf("%s/raw/%d/%d_%s", OUTPUT_FRAMES_DIR, prefix, prefix, fileExtension)
	audioInputPath := fmt.Sprintf("%s/bliptunes.mp3", INPUT_AUDIO_DIR)
	outputVideoPath := fmt.Sprintf("%s/%d/%d_output.webm", OUTPUT_VIDEO_DIR, prefix, prefix)
	ffmpeg.NewCommand("").
		Input(framesInputPath, nil, "", false).
		Input(audioInputPath, nil, "", false).
		OutputPath(outputVideoPath).Run()
	fmt.Println("Video generation finished for asset", prefix)
}
