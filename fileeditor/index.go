package fileeditor

import (
	"capec/utils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func isTextFile(filename string) bool {
	extension := filepath.Ext(filename)
	switch extension {
	case ".txt", ".js", ".go", ".py":
		return true
	default:
		return false
	}
}

func gatherTextFiles(dir string) ([]string, error) {
	var fileList []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && isTextFile(info.Name()) {
			fileList = append(fileList, path)
		}
		return nil
	})
	return fileList, err
}

func writeToFile(fileName string, updatedContent []byte) error {
	err := ioutil.WriteFile(fileName, []byte(updatedContent), 0644)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}
	return nil
}

func randomElement(slice []string) string {
	return slice[rand.Intn(len(slice))]
}

func generateRandomFilename(extension string) string {
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

func deleteAndRecreateFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	newFilename := generateRandomFilename(filepath.Ext(filename))
	newFilePath := filepath.Join(filepath.Dir(filename), newFilename)

	err = os.Remove(filename)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(newFilePath, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

func modifyFile(fileName string) error {
	content, err := ioutil.ReadFile(fileName)

	if err != nil {
		fmt.Println("Error reading file:", err)
		return err
	}
	rand.Seed(time.Now().UnixNano())
	actionCode := rand.Intn(2)
	switch actionCode {
	case 0:
		updatedContent := append(content, content...)
		return writeToFile(fileName, updatedContent)
	case 1:
		updatedContent := content
		if len(content)%2 == 1 {
			updatedContent = append(
				updatedContent,
				updatedContent[len(content)-1],
			)
		}
		a := len(updatedContent) / 2
		return writeToFile(fileName, updatedContent[0:a])
	default:
		return nil
	}

}

func processTextFile(fileName string) (string, error) {
	rand.Seed(time.Now().UnixNano())
	actionCode := rand.Intn(2)
	switch actionCode {
	case 0:
		return "RECREATED", deleteAndRecreateFile(fileName)
	case 1:
		return "UPDATED", modifyFile(fileName)
	default:
		return "NO_ACTION", nil
	}

}

func randomNumNotInArray(arr []int, max int) int {
	rand.Seed(time.Now().UnixNano())

	var newRandomNumber int
	for {
		newRandomNumber = rand.Intn(max)
		if !utils.Contains(arr, newRandomNumber) {
			break
		}
	}
	return newRandomNumber
}

func ModifyFiles(dir string, count int) error {
	fileList, err := gatherTextFiles(dir)
	if err != nil {
		return err
	}

	if len(fileList) == 0 {
		return fmt.Errorf("no text files found in the directory %s", dir)
	}

	var touchedFiles []int
	for i := 0; i < count; i += 1 {
		rand.Seed(time.Now().UnixNano())
		fileIndex := randomNumNotInArray(touchedFiles, count)
		touchedFiles = append(touchedFiles, fileIndex)
		randFile := fileList[fileIndex]
		action, err := processTextFile(randFile)
		if err != nil {
			return fmt.Errorf("error %s file: %s, error %s", action, randFile, err.Error())

		}
		fmt.Printf(".    File with name %s was %s \n", randFile, action)
	}
	return nil
}
