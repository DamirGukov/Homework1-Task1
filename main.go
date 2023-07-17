package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
)

func main() {
	MakeFile()

	file, err := os.Open("phoneNumber.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	regex := regexp.MustCompile(`\(?\b(\(?\d{3}\)?[-.\s]?\d{3}[-.\s]?\d{4})\b`)
	phoneNumbers := regex.FindAllString(string(content), -1)

	for _, number := range phoneNumbers {
		fmt.Println(number)
	}
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
