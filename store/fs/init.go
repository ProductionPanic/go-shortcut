package fs

import (
	"os"
	"regexp"
	"strings"
)

func EnsureFile(filePath string) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		_, err = os.Create(filePath)
		if err != nil {
			panic(err)
		}
	}
}

func FileHasString(filePath string, str string) bool {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// now we check if the file contains the string
	return strings.Contains(string(file), str)
}

func AppendToFile(filePath string, str string) {
	f, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}
}

func RemoveFromFile(filePath string, regexStr string) {
	file, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}
	// now we check if the file contains the string
	re := regexp.MustCompile(regexStr)
	newFile := re.ReplaceAll(file, []byte(""))
	err = os.WriteFile(filePath, newFile, 0644)
	if err != nil {
		panic(err)
	}
}

func Home() string {
	home, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return home
}

func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return !os.IsNotExist(err)
}
