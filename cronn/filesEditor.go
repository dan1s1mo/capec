package cronn

import (
	"capec/utils"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"strconv"
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

func deleteAndRecreateFile(filename string) error {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}

	newFilePath := utils.GetRandomFilename(filename)

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

func exists(arr []int, num int) bool {
	for _, val := range arr {
		if val == num {
			return true
		}
	}
	return false
}

func randomNumNotInArray(arr []int, max int) int {
	rand.Seed(time.Now().UnixNano())

	var newRandomNumber int
	for {
		newRandomNumber = rand.Intn(max)
		if !exists(arr, newRandomNumber) {
			break
		}
	}
	return newRandomNumber
}

func EditFiles() {
	args := os.Args

	if len(args) == 1 {
		fmt.Println("No directory path provided.")
		return
	}
	dir := args[1]
	count := 1
	if len(args) == 3 {
		num, err := strconv.Atoi(args[2])
		count = num
		if err != nil {
			fmt.Println("Error unable to convert number to:", err)
			return
		}
	}

	fileList, err := gatherTextFiles(dir)
	if err != nil {
		fmt.Println("Error reading directory:", err)
		return
	}

	if len(fileList) == 0 {
		fmt.Println("No text files found in the directory")
		return
	}

	var touchedFiles []int
	for i := 0; i < count; i += 1 {
		rand.Seed(time.Now().UnixNano())
		fileIndex := randomNumNotInArray(touchedFiles, count)
		touchedFiles = append(touchedFiles, fileIndex)
		randFile := fileList[fileIndex]
		action, err := processTextFile(randFile)
		if err != nil {
			fmt.Printf("Error %s file: %s, error %s", action, randFile, err.Error())
			continue
		}
		fmt.Printf(".    File with name %s was %s \n", randFile, action)
	}
}
