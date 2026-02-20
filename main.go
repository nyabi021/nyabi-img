package main

import (
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

func processImage(filePath string, outputPath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()

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
	defer out.Close()

	err = jpeg.Encode(out, img, &jpeg.Options{Quality: 90})
	if err != nil {
		return fmt.Errorf("Failed to encode image: %v\n", err)
	}

	return nil
}

func main() {
	targetFolder := getDownloadPath()
	fmt.Println("Target folder:", targetFolder)

	validExts := map[string]bool{
		".png":  true,
		".jpg":  true,
		".jpeg": true,
	}

	entries, err := os.ReadDir(targetFolder)
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
			fullPath := filepath.Join(targetFolder, filename)
			imagePaths = append(imagePaths, fullPath)
		}
	}

	fmt.Printf("Total images: %d\n", len(imagePaths))

	outputPath := filepath.Join(targetFolder, "output")
	err = os.MkdirAll(outputPath, 0755)
	if err != nil {
		log.Fatalf("Failed to create output directory: %v\n", err)
	}

	for _, imgPath := range imagePaths {
		err := processImage(imgPath, outputPath)
		if err != nil {
			log.Printf("Failed to process %s: %v\n", imgPath, err)
			continue
		}
		fmt.Printf("Successfully processed %s\n", filepath.Base(imgPath))
	}

	fmt.Println("Done")
}
