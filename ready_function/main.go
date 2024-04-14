package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func hex(str string) string { //Conversion function in base 16 (16lık tabandaki dönüşüm fonksiyonu)
	values, _ := strconv.ParseInt(str, 16, 64)
	return strconv.Itoa(int(values))
}

func bin(str string) string { // Conversion function in base 2 (2 lik tabandaki dönüşüm fonksiyonu)
	values, _ := strconv.ParseInt(str, 2, 64)
	return strconv.Itoa(int(values))
}

func up(str string) string { // Uppercase function (Büyük harfe çevirme fonksiyonu)
	return strings.ToUpper(str)
}

func low(str string) string { // Lowercase conversion function(Küçük harfe çevirme fonksiyonu)
	return strings.ToLower(str)
}

func cap(str string) string { // Capitalize the first letter of a word (Kelime ilk harfini büyük harfe çevirme)
	return strings.Title(str)
}

func punctuation_edit(statement string) string {
	punctuations := []string{",", ".", "!", "?", ";", ":", "'", "\"", "(", ")", "-"}

	for _, punctuation := range punctuations {
		statement = strings.Replace(statement, " "+punctuation, punctuation, -1)
		statement = strings.Replace(statement, punctuation+" ", punctuation, -1)
	}

	char_array := []rune(statement)
	output := []rune{}
	for i := 0; i < len(char_array)-1; i++ {
		if unicode.IsPunct(char_array[i]) && unicode.IsLetter(char_array[i+1]) {
			if char_array[i] == '\'' && unicode.IsLetter(char_array[i+1]) || char_array[i] == '-' && unicode.IsLetter(char_array[i+1]) || char_array[i] == '"' {
				output = append(output, char_array[i])
			} else {
				output = append(output, char_array[i], ' ')
			}
		} else if char_array[i] == ':' || char_array[i] == ';' && (char_array[i+1] == '\'' || (char_array[i+1]) == '"') {
			output = append(output, char_array[i], ' ')
		} else {
			output = append(output, char_array[i])
		}
	}
	output = append(output, char_array[len(char_array)-1])

	return string(output)
}

func ChangeA(s []string) []string {
	vowels := []string{"a", "e", "i", "o", "u", "h", "A", "E", "I", "O", "U", "H"}

	for i, word := range s {
		for _, letter := range vowels {
			if word == "a" && string(s[i+1][0]) == letter {
				s[i] = "an"
			} else if word == "A" && string(s[i+1][0]) == letter {
				s[i] = "An"
			}
		}
	}
	return s
}

func main() {
	samplefile, err := os.Open("sample.txt")
	if err != nil {
		fmt.Println("No file opened", err)
		return
	}
	defer samplefile.Close()

	line := bufio.NewScanner(samplefile) // reads file data line by line (dosya verilerini satır satır okur)

	resultFile, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("No file opened", err)
		return
	}
	defer resultFile.Close()

	for line.Scan() {
		line := line.Text() // allows us to read lines (satırları okumamızı sağlar)
		words := strings.Fields(line)
		words = ChangeA(words)

		output := ""

		for i := len(words) - 1; i >= 0; i-- {
			if words[i] == "(up)" {
				words[i-1] = up(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(up," {
				value := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(value)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = up(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(low)" {
				words[i-1] = low(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(low," {
				value := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(value)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = low(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(cap)" {
				words[i-1] = cap(words[i-1])
				i--

			} else if strings.Contains(words[i], ")") && words[i-1] == "(cap," {
				value := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(value)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = cap(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(hex)" {
				words[i-1] = hex(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(hex," {
				value := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(value)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = hex(words[i-n+a-1])
				}
				i--
				continue
			} else if words[i] == "(bin)" {
				words[i-1] = bin(words[i-1])
				i--
			} else if strings.Contains(words[i], ")") && words[i-1] == "(bin," {
				value := strings.Trim(words[i], ")")
				n, _ := strconv.Atoi(value)
				for a := 0; a < n; a++ {
					words[i-n+a-1] = bin(words[i-n+a-1])
				}
				i--
				continue
			}
			output = words[i] + " " + output
		}

		line = strings.Join(words, " ")
		line = punctuation_edit(output)

		// Result.txt write to the file
		resultFile, err := os.OpenFile("result.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
		if err != nil {
			fmt.Println("file write error", err)
			return
		}
		defer resultFile.Close()

		_, err = resultFile.WriteString(line + "\n")
		if err != nil {
			fmt.Println("file write error", err)
			return
		}
	}

	if err := line.Err(); err != nil {
		fmt.Println("file read error", err)
		return
	}
}
