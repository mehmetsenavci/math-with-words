package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// var path string
	// fmt.Scanln(&path)
	f, err := os.Open("sample.in")
	if err != nil {
		fmt.Print(err.Error())
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	wordToInt := make(map[string]int)
	intToWord := make(map[int]string)

	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		if words[0] == "def" {
			val, err := strconv.Atoi(words[2])
			if err != nil {
				fmt.Println(err.Error())
			}

			wordToInt[words[1]] = val
			for k, v := range intToWord {
				if v == words[1] {
					delete(intToWord, k)
				}
			}
			intToWord[val] = words[1]

		} else if words[0] == "calc" {
			total := 0
			totalStr := ""
			math := make(map[string]int)
			for i := 1; i < len(words); i++ {
				if !isOperant(words[i]) {
					math[words[i]] = wordToInt[words[i]]
				}
			}
			for i := 1; i < len(words); i++ {
				if isOperant(words[i]) {
					switch words[i] {
					case "-":
						math[words[i+1]] = math[words[i+1]] * -1
						i++
					case "+":
					case "=":
						for _, v := range math {
							total += v
						}
						totalStr, _ = getValueFromIntMap(intToWord, total)
					}
				} else {
					_, message := getValueFromWordMap(wordToInt, words[i])
					if message != "" {
						totalStr = message
						break
					}

				}
			}
			fmt.Println(totalStr)
		} else if words[0] == "clear" {
			for k := range wordToInt {
				delete(wordToInt, k)
			}
			for k := range intToWord {
				delete(intToWord, k)
			}
		}
	}
	fmt.Println(wordToInt)
	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}

func isOperant(word string) bool {
	return word == "-" || word == "+" || word == "="
}

func getValueFromWordMap(wordToInt map[string]int, word string) (int, string) {
	val, ok := wordToInt[word]
	if ok {
		return val, ""
	}
	return 0, "unknown"
}

func getValueFromIntMap(intToWord map[int]string, in int) (string, string) {
	val, ok := intToWord[in]
	if ok {
		return val, ""
	}
	return "unknown", "unknown"
}
