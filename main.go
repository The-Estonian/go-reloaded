package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello!")
	data, error := os.ReadFile("sample.txt")
	splitData := strings.Split(string(data), " ")
	if error != nil {
		panic("Error with file")
	} else {
		var newStringList []string
		counter := 0
		for i := 0; i < len(splitData)-1; i++ {
			if splitData[i+1] == "(up)" {
				newStringList = append(newStringList, strings.ToUpper(splitData[i]))
				i++
			} else if strings.HasPrefix(splitData[i+1], "(low,") {
				counter, _ = strconv.Atoi(string(splitData[i+2][0]))
				fmt.Println(counter)
			} else if splitData[i+1] == "(low)" {
				newStringList = append(newStringList, strings.ToLower(splitData[i]))
				i++
			} else if splitData[i+1] == "(bin)" {
				number, _ := strconv.ParseInt(splitData[i], 2, 64)
				newStringList = append(newStringList, strconv.FormatInt(number, 10))
				i++
			} else if splitData[i+1] == "(hex)" {
				number, _ := strconv.ParseInt(splitData[i], 16, 64)
				newStringList = append(newStringList, strconv.FormatInt(number, 10))
				i++
			} else if splitData[i+1] == "(cap)" {
				newStringList = append(newStringList, cases.Title(language.Und, cases.NoLower).String(splitData[i]))
				i++
			} else {
				newStringList = append(newStringList, splitData[i])
			}
		}
		newStringList = append(newStringList, splitData[len(splitData)-1])
		fmt.Println(newStringList)
		newFile, error := os.Create("result.txt")
		if error != nil {
			panic("Error creating file")
		} else {
			newFile.WriteString("newStringList")
		}
	}
}
