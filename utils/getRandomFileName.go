package utils

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}

func randomElement(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func getRandomFilename(extension string) string {
	adjectives := []string{
		"happy", "friendly", "lovely", "graceful", "bold", "mysterious",
	}
	nouns := []string{
		"lion", "eagle", "panda", "dolphin", "fox", "elephant",
	}

	rand.Seed(time.Now().UnixNano())
	randomAdjective := randomElement(adjectives)
	randomNoun := randomElement(nouns)
	randomNumber := rand.Intn(1000)

	filename := fmt.Sprintf("%s_%s_%03d%s", randomAdjective, randomNoun, randomNumber, extension)
	return filename
}

func GetRandomFilename(filename string) string {
	var newFilePath string

	for {
		newFilename := getRandomFilename(filepath.Ext(filename))
		newFilePath = filepath.Join(filepath.Dir(filename), newFilename)
		if !fileExists(newFilePath) {
			break
		}
	}
	return newFilePath
}
