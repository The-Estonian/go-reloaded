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
	data, error := os.ReadFile("sample.txt")
	splitData := strings.Split(string(data), " ")
	if error != nil {
		panic("Error with file")
	} else {
		var newStringList []string
		for i := 0; i < len(splitData); i++ {
			if splitData[i] == "(up)" {
				newStringList[len(newStringList)-1] = strings.ToUpper(newStringList[len(newStringList)-1])
			} else if splitData[i] == "(hex)" {
				number, _ := strconv.ParseInt(newStringList[len(newStringList)-1], 16, 64)
				newStringList[len(newStringList)-1] = strconv.FormatInt(number, 10)
			} else if splitData[i] == "(bin)" {
				number, _ := strconv.ParseInt(newStringList[len(newStringList)-1], 2, 64)
				newStringList[len(newStringList)-1] = strconv.FormatInt(number, 10)
			} else if splitData[i] == "(low)" {
				newStringList[len(newStringList)-1] = strings.ToLower(newStringList[len(newStringList)-1])
			} else if splitData[i] == "(cap)" {
				newStringList[len(newStringList)-1] = cases.Title(language.Und, cases.NoLower).String(newStringList[len(newStringList)-1])
			} else if strings.HasPrefix(splitData[i], "(low,") {
				iteratorCount, _ := strconv.Atoi(string(splitData[i+1][0]))
				for j := iteratorCount - 1; j >= 0; j-- {
					newStringList[len(newStringList)-j-1] = strings.ToLower(newStringList[len(newStringList)-j-1])
				}
				i++
			} else if strings.HasPrefix(splitData[i], "(up,") {
				iteratorCount, _ := strconv.Atoi(string(splitData[i+1][0]))
				for j := iteratorCount - 1; j >= 0; j-- {
					newStringList[len(newStringList)-j-1] = strings.ToUpper(newStringList[len(newStringList)-j-1])
				}
				i++
			} else if strings.HasPrefix(splitData[i], "(cap,") {
				iteratorCount, _ := strconv.Atoi(string(splitData[i+1][0]))
				for j := iteratorCount - 1; j >= 0; j-- {
					newStringList[len(newStringList)-j-1] = cases.Title(language.Und, cases.NoLower).String(newStringList[len(newStringList)-j-1])
				}
				i++
			} else if splitData[i] == "a" && string(splitData[i+1][0]) == "a" ||
				splitData[i] == "a" && string(splitData[i+1][0]) == "e" ||
				splitData[i] == "a" && string(splitData[i+1][0]) == "i" ||
				splitData[i] == "a" && string(splitData[i+1][0]) == "o" ||
				splitData[i] == "a" && string(splitData[i+1][0]) == "u" ||
				splitData[i] == "a" && string(splitData[i+1][0]) == "h" {
				newStringList = append(newStringList, "an")
			} else if string(splitData[i][0]) == "." && string(splitData[i][1]) != "." {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + "."
				newStringList = append(newStringList, splitData[i][1:])
			} else if string(splitData[i][0]) == "," {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + ","
				newStringList = append(newStringList, splitData[i][1:])
			} else if string(splitData[i][0]) == "!" && len(splitData[i]) == 1 {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + "!"
				newStringList = append(newStringList, splitData[i][1:])
			} else if string(splitData[i][0]) == "?" && len(splitData[i]) == 1 {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + "?"
				newStringList = append(newStringList, splitData[i][1:])
			} else if string(splitData[i][0]) == "!" && string(splitData[i][1]) == "!" ||
				string(splitData[i][0]) == "!" && string(splitData[i][1]) == "?" ||
				string(splitData[i][0]) == "?" && string(splitData[i][1]) == "!" ||
				string(splitData[i][0]) == "?" && string(splitData[i][1]) == "?" {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + splitData[i]
			} else if string(splitData[i][0]) == ":" {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + ":"
				newStringList = append(newStringList, splitData[i][1:])
			} else if string(splitData[i][0]) == ";" {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + ";"
				newStringList = append(newStringList, splitData[i][1:])
			} else if string(splitData[i][0]) == "." && string(splitData[i][1]) == "." {
				newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + "..."
			} else if string(splitData[i]) == "'" {
				
			} else {
				newStringList = append(newStringList, splitData[i])
			}
		}
		newFile, error := os.Create("result.txt")
		if error != nil {
			panic("Error creating file")
		} else {
			returnSentence := ""
			for i := 0; i < len(newStringList); i++ {
				returnSentence += newStringList[i]
				if i < len(newStringList)-2 {
					returnSentence += " "
				}
			}
			fmt.Println(returnSentence)
			newFile.WriteString(returnSentence)
		}
	}
}
