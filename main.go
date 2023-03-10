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
	var input__file, output__file, trim string
	var new_string []string
	var tr int

	if len(os.Args) == 3 {
		input__file = os.Args[1]
		output__file = os.Args[2]

		content_file, _ := readFile(input__file)

		tab_string := splitIntoWords(content_file)

		for i, world := range tab_string {
			if containsString(world, "(low,") || world == "(low)" {
				tr++
				trim = strings.Trim(tab_string[i+1], ")")
				int_trim := stringToInt(trim)
				for nb := int_trim; nb > 0; nb-- {
					k := 1
					tab_string[i-nb] = makeLowercase(tab_string[i-nb])
					k++
					int_trim--
				}
				deleteElement(tab_string, i)
				deleteElement(tab_string, i)

			} else if containsString(world, "(cap,") || world == "(cap)" {
				tr++
				trim = strings.Trim(tab_string[i+1], ")")
				int_trim := 0
				if stringToInt(trim) > 0 {
					int_trim = stringToInt(trim)

				} else {
					int_trim = 1
				}

				for nb := int_trim; nb > 0; nb-- {
					k := 1
					tab_string[i-nb] = capitalize(tab_string[i-nb])
					k++
					int_trim--
				}
				deleteElement(tab_string, i)
				deleteElement(tab_string, i)

			} else if containsString(world, "(up,") || world == "(up)" {
				tr++
				trim = strings.Trim(tab_string[i+1], ")")
				int_trim := stringToInt(trim)
				for nb := int_trim; nb > 0; nb-- {
					k := 1
					tab_string[i-nb] = makeUppercase(tab_string[i-nb])
					k++
					int_trim--
				}
				deleteElement(tab_string, i)
				deleteElement(tab_string, i)

			} else if world == "(bin)" {
				tr++
				tab_string[i-1] = binaryToDecimal(tab_string[i-1])
				deleteElement(tab_string, i)

			} else if world == "(hex)" {
				tr++
				tab_string[i-1], _ = hexToDec(tab_string[i-1])
				deleteElement(tab_string, i)

			}

		}

		c := len(tab_string) - (2 * tr) + 1

		for j := 0; j <= c; j++ {
			new_string = append(new_string, tab_string[j])

		}

		fmt.Println(tab_string)
		stringg := tabTostring(new_string)
		fmt.Println(stringg)

		writeToFile(output__file, stringg)

	}

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

func removeNonWords(text string) string {
	re := regexp.MustCompile(`(\([^()]*\))|([^a-zA-Z0123456789\s]+)`)
	words := re.ReplaceAllString(text, " ")
	return strings.TrimSpace(words)
}
func tabTostring(words []string) string {
	txt := ""
	for _, str := range words {
		txt = txt + str + " "

	}
	return txt
}
