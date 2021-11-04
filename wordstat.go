package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
)

var inputFile string
var outputFile string
var ignoreNumbers bool
var helpFlag bool

func printHelp() {
	helpText := `
GoWordStat by wachttijd (https://github.com/wachttijd/go-wordstat)

Available flags:
	-i Input file for statistical word analysis (default: "input.txt")
	-o Output file for results of analysis (default: "output.json")
	-ignums Ignore numbers (default: false)
	-h Print this help
`
	fmt.Println(helpText)
}

func main() {
	flag.StringVar(&inputFile, "i", "input.txt", "Input file for statistical analysis")
	flag.StringVar(&outputFile, "o", "output.json", "Output file for results of analysis")
	flag.BoolVar(&ignoreNumbers, "ignums", false, "Ignore numbers")
	flag.BoolVar(&helpFlag, "h", false, "Print help")

	flag.Parse()

	if helpFlag {
		printHelp()
		os.Exit(0)
	}

	textBytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}
	text := string(textBytes)

	fmt.Println("Preparing text for analysis...")

	text = cleanText(text)

	fmt.Println("Text prepared, analysis started...")

	data := []byte(sortMapToJSONString(stringAnalysis(text)))
	ioutil.WriteFile(outputFile, data, 0644)

	fmt.Println("Analysis complete!")
}

func stringAnalysis(text string) map[string]int {
	allWords := strings.Fields(text)
	wordsStat := map[string]int{}

	for _, word := range allWords {
		if inKeys(wordsStat, word) {
			continue
		}

		for _, compaire_word := range allWords {
			if word == compaire_word {
				wordsStat[word] += 1
			}
		}
	}

	return wordsStat
}

func inKeys(mymap map[string]int, val string) bool {
	keys := make([]string, 0, len(mymap))
	for k := range mymap {
		keys = append(keys, k)
	}

	for _, key := range keys {
		if key == val {
			return true
		}
	}

	return false
}

func sortMapToJSONString(m map[string]int) string {
	type key_value struct {
		Key   string
		Value int
	}

	var sorted_struct []key_value
	var outString string

	for key, value := range m {
		sorted_struct = append(sorted_struct, key_value{key, value})
	}

	sort.Slice(sorted_struct, func(i, j int) bool {
		return sorted_struct[i].Value > sorted_struct[j].Value
	})

	for _, key_value := range sorted_struct {
		outString += "\"" + key_value.Key + "\":" + strconv.Itoa(key_value.Value) + ",\n"
	}

	outString = "{\n" + outString + "}"

	outString = strings.Replace(outString, ",\n}", "\n}", 1)

	return outString
}

func cleanText(text string) string {
	if ignoreNumbers {
		text = strings.ReplaceAll(text, "0", "")
		text = strings.ReplaceAll(text, "1", "")
		text = strings.ReplaceAll(text, "2", "")
		text = strings.ReplaceAll(text, "3", "")
		text = strings.ReplaceAll(text, "4", "")
		text = strings.ReplaceAll(text, "5", "")
		text = strings.ReplaceAll(text, "6", "")
		text = strings.ReplaceAll(text, "7", "")
		text = strings.ReplaceAll(text, "8", "")
		text = strings.ReplaceAll(text, "9", "")
	}

	text = strings.ReplaceAll(text, "\n", " ")
	text = strings.ReplaceAll(text, "\r", " ")
	text = strings.ReplaceAll(text, "\t", " ")
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, "(", " ")
	text = strings.ReplaceAll(text, ")", " ")
	text = strings.ReplaceAll(text, "[", " ")
	text = strings.ReplaceAll(text, "]", " ")
	text = strings.ReplaceAll(text, "{", " ")
	text = strings.ReplaceAll(text, "}", " ")
	text = strings.ReplaceAll(text, "!", " ")
	text = strings.ReplaceAll(text, "?", " ")
	text = strings.ReplaceAll(text, "\"", " ")
	text = strings.ReplaceAll(text, "«", " ")
	text = strings.ReplaceAll(text, "»", " ")
	text = strings.ReplaceAll(text, "“", " ")
	text = strings.ReplaceAll(text, "”", " ")
	text = strings.ReplaceAll(text, "'", " ")
	text = strings.ReplaceAll(text, ":", " ")
	text = strings.ReplaceAll(text, ";", " ")

	text = strings.ReplaceAll(text, "  ", " ")
	text = strings.ReplaceAll(text, "   ", " ")
	text = strings.ReplaceAll(text, "    ", " ")
	text = strings.ReplaceAll(text, "     ", " ")
	return strings.ToLower(text)
}
