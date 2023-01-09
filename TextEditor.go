package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func GetFileLine(PathName string, line int) {
	file, err := os.Open(PathName)

	if err != nil {
		fmt.Println(err)
		file.Close()
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

	file.Close()
}

func concatenateDir(dirName string, nextPath string) string {
	var b bytes.Buffer

	b.WriteString(dirName)
	b.WriteString("/")
	b.WriteString(nextPath)

	return b.String()
}

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}

	correctDir := concatenateDir(dir, "FileEditor/Documentos/Animais.txt")

	//fmt.Println(correctDir)

	GetFileLine(correctDir, 5)

}
