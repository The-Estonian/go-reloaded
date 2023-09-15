package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func capitalizer(str string) string {
	return strings.ToUpper(string(str[0])) + str[1:]
}

func main() {
	if len(os.Args) == 3 {
		inputFile := os.Args[1]
		outputFile := os.Args[2]
		data, error := os.ReadFile(inputFile)
		splitData := strings.Split(string(data), " ")
		if error != nil {
			panic("Error with file")
		} else {
			punctuationStart := false
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
					newStringList[len(newStringList)-1] = capitalizer(string(newStringList[len(newStringList)-1]))
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
						newStringList[len(newStringList)-j-1] = capitalizer(string(newStringList[len(newStringList)-j-1]))
					}
					i++
				} else if splitData[i] == "a" && string(splitData[i+1][0]) == "a" ||
					splitData[i] == "a" && string(splitData[i+1][0]) == "e" ||
					splitData[i] == "a" && string(splitData[i+1][0]) == "i" ||
					splitData[i] == "a" && string(splitData[i+1][0]) == "o" ||
					splitData[i] == "a" && string(splitData[i+1][0]) == "u" ||
					splitData[i] == "a" && string(splitData[i+1][0]) == "h" {
					newStringList = append(newStringList, "an")
				} else if string(splitData[i][0]) == "." && len(string(splitData[i])) == 1 {
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
					if !punctuationStart {
						splitData[i+1] = "'" + splitData[i+1]
						punctuationStart = true
					} else {
						newStringList[len(newStringList)-1] = newStringList[len(newStringList)-1] + "'"
						punctuationStart = false
					}
				} else {
					newStringList = append(newStringList, splitData[i])
				}
			}
			newFile, error := os.Create(outputFile)
			if error != nil {
				panic("Error creating file")
			} else {
				returnSentence := ""
				for i := 0; i < len(newStringList); i++ {
					returnSentence += newStringList[i]
					if len(newStringList[i]) > 0 {
						if i < len(newStringList)-2 {
							returnSentence += " "
						}
					}
				}
				fmt.Println(returnSentence)
				newFile.WriteString(returnSentence)
			}
		}
	} else {
		fmt.Println("Please add input and output file names")
		return
	}
}
