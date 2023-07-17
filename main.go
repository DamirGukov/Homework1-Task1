package main

import (
	"errors"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	MakeFile()
	ReadFile()

	file, err := os.Open("phoneNumber.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	data := make([]byte, 4)
	regex := regexp.MustCompile(`\b(\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4})\b`)

	for {
		n, err := file.Read(data)
		if errors.Is(err, io.EOF) {
			break
		}

		phoneNumbers := regex.FindAllString(string(data[:n]), -1)
		for _, number := range phoneNumbers {
			fmt.Println(number)
		}
	}
}

func ReadFile() {
	fileContent, err := os.ReadFile("phoneNumber.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(fileContent))
}

func MakeFile() {
	file, err := os.Create("phoneNumber.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer file.Close()

	_, err = file.WriteString("1234567890\n(123) 456-7890\n(123)456-7890\n123-456-7890\n123.456.7890\n123 456 7890")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}
