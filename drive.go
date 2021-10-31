package main

import (
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

func main() {
	textBytes, err := ioutil.ReadFile("vm.txt")
	if err != nil {
		panic(err)
	}

	text := string(textBytes)

	text = strings.ReplaceAll(text, "\n", "")
	text = strings.ReplaceAll(text, "\r", "")
	text = strings.ReplaceAll(text, "\t", "")
	text = strings.ReplaceAll(text, ",", " ")
	text = strings.ReplaceAll(text, ".", " ")
	text = strings.ReplaceAll(text, "(", " ")
	text = strings.ReplaceAll(text, ")", " ")
	text = strings.ReplaceAll(text, "!", " ")
	text = strings.ReplaceAll(text, "?", " ")
	text = strings.ReplaceAll(text, "\"", " ")
	text = strings.ReplaceAll(text, "«", " ")
	text = strings.ReplaceAll(text, "»", " ")
	text = strings.ReplaceAll(text, "'", " ")
	text = strings.ReplaceAll(text, ":", " ")
	text = strings.ReplaceAll(text, ";", " ")
	text = strings.ReplaceAll(text, "  ", " ")
	text = strings.ReplaceAll(text, "   ", " ")
	text = strings.ReplaceAll(text, "    ", " ")
	text = strings.ToLower(text)

	data := []byte(sortMapToString(stringAnalysis(text)))

	ioutil.WriteFile("res.txt", data, 0644)
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

func sortMapToString(m map[string]int) string {
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

	return outString
}
