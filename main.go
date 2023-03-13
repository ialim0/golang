package main

import (
	"fmt"
	"io/ioutil"
	"math/big"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var input__file, output__file, trim, stringg string
	var new_string []string
	var tr, ver, c, fin int
	ver = 0

	fin = 0

	if len(os.Args) == 3 {
		input__file = os.Args[1]
		output__file = os.Args[2]

		content_file, _ := readFile(input__file)

		if containsString(content_file, "(cap)") {
			content_file = replaceWord(content_file, "(cap)", "(cap, 1)")
		}
		if containsString(content_file, "(low)") {
			content_file = replaceWord(content_file, "(low)", "(low, 1)")
		}
		if containsString(content_file, "(up)") {
			content_file = replaceWord(content_file, "(up)", "(up, 1)")
		}

		tab_string := splitIntoWords(content_file)
		if containsString(tab_string[len(tab_string)-2], ")") {
			fin = 1

		}
		//fmt.Println(tab_string)

		for i, world := range tab_string {

			if containsString(world, "(low,") || containsString(world, "(cap,") || containsString(world, "(up,") {
				ver++
				if containsString(world, "(low,") {

					tr++
					trim = strings.Trim(tab_string[i+1], ")")

					int_trim := stringToInt(trim)

					for nb := int_trim; nb > 0; nb-- {

						tab_string[i-nb] = makeLowercase(tab_string[i-nb])

						int_trim--
					}
					deleteElement(tab_string, i)
					deleteElement(tab_string, i)

				} else if containsString(world, "(cap,") {
					tr++
					trim = strings.Trim(tab_string[i+1], ")")

					int_trim := stringToInt(trim)

					for nb := int_trim; nb > 0; nb-- {

						tab_string[i-nb] = capitalize(tab_string[i-nb])

						int_trim--
					}
					deleteElement(tab_string, i)
					deleteElement(tab_string, i)

				} else if containsString(world, "(up,") {
					tr++
					trim = strings.Trim(tab_string[i+1], ")")

					int_trim := stringToInt(trim)

					for nb := int_trim; nb > 0; nb-- {

						tab_string[i-nb] = makeUppercase(tab_string[i-nb])

						int_trim--
					}
					deleteElement(tab_string, i)
					deleteElement(tab_string, i)

				}

			} else {

				if world == "(bin)" {
					tr++
					tab_string[i-1] = binaryToDecimal(tab_string[i-1])
					deleteElement(tab_string, i)

				} else if world == "(hex)" {
					tr++
					tab_string[i-1], _ = hexToDec(tab_string[i-1])
					deleteElement(tab_string, i)

				}

			}

		}
		if tr > 0 {

			if ver > 0 {
				c = len(tab_string) - (2 * tr) - 1

			} else {
				c = len(tab_string) - (2 * tr) + 1
			}

			for j := 0; j <= c; j++ {
				new_string = append(new_string, tab_string[j])

			}

			if fin > 0 {
				stringg = tabbTostring(new_string)

			} else {
				stringg = tabTostring(new_string)
			}

		} else {
			stringg = tabTostring(tab_string)
		}

		//fmt.Println(stringg)
		if countExpression(stringg, "'") >= 2 {
			stringg = replaceWord(stringg, "' ", "'")
			stringg = replaceWord(stringg, " '", "'")

		}

		writeToFile(output__file, stringg)

	}
	fmt.Println(formatText("I was sitting over there ,and ' then BAMM ! ! I was thinking ... You ' were right!?"))

}

func readFile(filename string) (string, error) {
	// Read the contents of the file
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}

	// Convert the content to a string and return it
	return string(content), nil
}

func writeToFile(filename string, text string) error {
	// Convert the string to a byte slice
	data := []byte(text)

	// Write the byte slice to the file
	err := ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return err
	}

	return nil
}

func hexToDec(hex string) (string, error) {
	// Convert the hex string to a big.Int
	i, success := new(big.Int).SetString(hex, 16)
	if !success {
		return "", fmt.Errorf("%s is not a valid hexadecimal string", hex)
	}

	// Convert the big.Int to a decimal string
	dec := i.String()

	return dec, nil
}

func binaryToDecimal(binary string) string {
	decimal := 0
	power := 1

	for i := len(binary) - 1; i >= 0; i-- {
		if binary[i] == '1' {
			decimal += power
		}
		power *= 2
	}

	return strconv.Itoa(decimal)
}

func makeUppercase(s string) string {
	return strings.ToUpper(s)
}

func makeLowercase(s string) string {
	return strings.ToLower(s)
}

func capitalize(s string) string {
	if len(s) == 0 {
		return s
	}
	firstChar := strings.ToUpper(string(s[0]))
	rest := s[1:]
	return firstChar + rest
}

// Function that take a string and convert it in array of word
func splitIntoWords(text string) []string {
	words := strings.Fields(text)
	return words
}

// Function to know if a string contain a giving string
func containsString(str, substr string) bool {
	return strings.Contains(str, substr)
}

// Function that convert string in number
func stringToInt(str string) int {
	i, _ := strconv.Atoi(str)

	return i
}

// Delete in array
func deleteElement(arr []string, index int) []string {
	// Remove element by appending elements before and after the one to be deleted
	return append(arr[:index], arr[index+1:]...)
}

func tabTostring(words []string) string {
	txt := ""
	for i, _ := range words {

		if i < len(words)-1 && startsWith(words[i+1], ",") {

			txt = txt + words[i]
		} else if startsWith(words[i], ",") && !startsWith(words[i+1], ",") {

			words[i] = getSuffix(words[i], ",")

			txt = txt + "," + " " + words[i] + " "

		} else if i < len(words)-1 && startsWith(words[i+1], ".") && !startsWith(words[i], "..") {

			txt = txt + words[i]
		} else if startsWith(words[i], ".") && !startsWith(words[i], "..") {

			words[i] = getSuffix(words[i], ".")

			txt = txt + "." + " " + words[i]

		} else if i < len(words)-1 && startsWith(words[i+1], "...") {

			txt = txt + words[i]
		} else if startsWith(words[i], "...") {

			words[i] = getSuffix(words[i], "...")

			txt = txt + "..." + " " + words[i] + " "

		} else if i < len(words)-1 && startsWith(words[i+1], ":") {

			txt = txt + words[i]
		} else if startsWith(words[i], ":") && !startsWith(words[i+1], ":") {

			words[i] = getSuffix(words[i], ":")

			txt = txt + ":" + " " + words[i] + " "

		} else if i == len(words)-1 {

			txt = txt + words[i]
		} else if i < len(words)-1 && startsWith(words[i+1], "!?") {

			txt = txt + words[i]
		} else if startsWith(words[i], "!?") {

			words[i] = getSuffix(words[i], "!?")

			txt = txt + "!?" + " " + words[i] + " "

		} else if words[i] == "a" || words[i] == "A" && i < len(words)-1 {
			if startsWithVowel(words[i+1]) {
				txt = txt + words[i] + "n" + " "

			} else {
				txt = txt + words[i] + " "
			}

		} else {

			txt = txt + words[i] + " "

		}

	}
	return txt
}
func tabbTostring(words []string) string {
	txt := ""

	for i, _ := range words {

		if i < len(words)-2 {
			txt = txt + words[i] + " "

		} else {
			txt = txt + words[i]
		}

	}

	return txt
}

func splitSentenceIntoWords(sentence string) []string {
	// Split the sentence into words using whitespace as the delimiter
	words := strings.Fields(sentence)

	return words
}

// start with
func startsWith(s, prefix string) bool {
	return len(s) >= len(prefix) && s[0:len(prefix)] == prefix
}

// Function to get the suffix
func getSuffix(word string, prefix string) string {
	if strings.HasPrefix(word, prefix) {
		return word[len(prefix):]
	}
	return ""
}

// Help know if giving instance start with
func startsWithVowel(word string) bool {
	vowels := []string{"a", "e", "i", "o", "u", "h"}
	for _, v := range vowels {
		if strings.HasPrefix(strings.ToLower(word), v) {
			return true
		}
	}
	return false
}

// Function that count how many time a giving expression is in sentence
func countExpression(sentence string, expression string) int {
	count := 0
	words := strings.Split(sentence, " ")
	for _, word := range words {
		if strings.Contains(word, expression) {
			count++
		}
	}
	return count
}

// Function
func replaceWord(sentence, oldWord, newWord string) string {
	// Replace all occurrences of `oldWord` with `newWord` in the `sentence`.
	return strings.ReplaceAll(sentence, oldWord, newWord)
}

func formatText(text string) string {
	// Close punctuations to the previous word and separate with a space
	pattern := regexp.MustCompile(`\s*([.,!?;:])\s*`)
	text = pattern.ReplaceAllString(text, "$1 ")

	// Remove extra spaces
	text = strings.TrimSpace(text)
	text = strings.ReplaceAll(text, "  ", " ")

	// Format groups of punctuations
	text = strings.ReplaceAll(text, "...", "...")
	text = strings.ReplaceAll(text, "!?", "!?")
	text = strings.ReplaceAll(text, "! !", "!!")
	text = strings.ReplaceAll(text, "? ?", "??")
	text = strings.ReplaceAll(text, ". .", "..")
	text = strings.ReplaceAll(text, "' ", "'")
	text = strings.ReplaceAll(text, " '", "'")

	return text
}
