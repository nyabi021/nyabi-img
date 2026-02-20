package main

import (
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	_ "image/png"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func getDownloadPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("Failed to get user home directory: %v\n", err)
	}

	folderPath := filepath.Join(homeDir, "Downloads", "twitter_media_harvest")

	return folderPath
}

func processImage(filePath string, outputPath string, quality int) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("Failed to close file: %v\n", err)
		}
	}()

	img, _, err := image.Decode(file)
	if err != nil {
		return fmt.Errorf("Failed to decode image: %v\n", err)
	}

	baseName := filepath.Base(filePath)
	ext := filepath.Ext(baseName)
	nameWithoutExt := strings.TrimSuffix(baseName, ext)
	newFileName := nameWithoutExt + ".jpg"

	newFilePath := filepath.Join(outputPath, newFileName)
	out, err := os.Create(newFilePath)
	if err != nil {
		return fmt.Errorf("Failed to create file: %v\n", err)
	}

	defer func() {
		_ = out.Close()
	}()

	err = jpeg.Encode(out, img, &jpeg.Options{Quality: quality})
	if err != nil {
		return fmt.Errorf("Failed to encode image: %v\n", err)
	}

	return nil
}

func main() {
	var inputDir string
	var outputDir string
	var quality int

	defaultInput := getDownloadPath()
	defaultOutput := filepath.Join(defaultInput, "output")

	flag.StringVar(&inputDir, "input", defaultInput, "Input directory")
	flag.StringVar(&outputDir, "output", defaultOutput, "Output directory")
	flag.IntVar(&quality, "q", 90, "Output image quality (1-100)")

	flag.Parse()

	fmt.Println("Target folder:", inputDir)

	validExts := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	entries, err := os.ReadDir(inputDir)
	if err != nil {
		log.Fatalf("Failed to read directory: %v\n", err)
	}

	var imagePaths []string

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		filename := entry.Name()

		ext := strings.ToLower(filepath.Ext(filename))

		if validExts[ext] {
			fullPath := filepath.Join(inputDir, filename)
			imagePaths = append(imagePaths, fullPath)
		}
	}

	fmt.Printf("Total images: %d\n", len(imagePaths))

	outputPath := filepath.Join(inputDir, "output")
	err = os.MkdirAll(outputPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v\n", err)
	}

	for _, imgPath := range imagePaths {
		err := processImage(imgPath, outputPath, quality)
		if err != nil {
			log.Printf("Failed to process %s: %v\n", imgPath, err)
			continue
		}
		fmt.Printf("Successfully processed %s\n", filepath.Base(imgPath))
	}

	fmt.Println("Done")
}
