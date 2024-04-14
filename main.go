package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func hex(hex string) string { //Conversion function in base 16 (16lık tabandaki dönüşüm fonksiyonu)
	decimal := 0
	for index, digit := range hex {
		number := 0
		if digit >= '0' && digit <= '9' {
			number = int(digit - '0')
		} else if digit >= 'a' && digit <= 'f' {
			number = int(digit - 'a' + 10)
		} else if digit >= 'A' && digit <= 'F' {
			number = int(digit - 'A' + 10)
		} else {
			return "Hata"
		}

		power := len(hex) - index - 1

		for i := 0; i < power; i++ {
			number *= 16
		}

		decimal += number
	}
	return strconv.Itoa(int(decimal))
}

func bin(bin string) string { // Conversion function in base 2 (2 lik tabandaki dönüşüm fonksiyonu)
	decimal := 0

	for index, digit := range bin {

		number := 0

		if digit == '0' || digit == '1' {
			number = int(digit - '0')
		} else {
			return "ERROR"
		}

		power := len(bin) - index - 1

		for i := 0; i < power; i++ {
			number *= 2
		}

		decimal += number

	}
	return strconv.Itoa(int(decimal))
}

func up(str string) string { // Uppercase function (Büyük harfe çevirme fonksiyonu)
	output := ""
	for _, i := range str {

		if (i >= 'a' && i <= 'z') || i >= 'A' && i <= 'Z' {

			if i >= 'a' && i <= 'z' {
				i = i - 'a' + 'A'
				output += string(i)
			} else {
				output += string(i)
			}

		} else {
			output += string(i)
		}

	}
	return output
}

func low(str string) string { // Lowercase conversion function(Küçük harfe çevirme fonksiyonu)
	output := ""
	for _, i := range str {

		if (i >= 'a' && i <= 'z') || i >= 'A' && i <= 'Z' {

			if i >= 'A' && i <= 'Z' {
				i = i - 'A' + 'a'
				output += string(i)
			} else {
				output += string(i)
			}

		} else {
			output += string(i)
		}

	}
	return output
}

func cap(str string) string { // Capitalize the first letter of a word (Kelime ilk harfini büyük harfe çevirme)
	output := ""
	for index, i := range str {
		if index == 0 && (i >= 'a' && i <= 'z') {
			i = i - 'a' + 'A'
			output += string(i)
		} else {
			output += string(i)
		}
	}
	return output
}

func rune_control(str string, char rune) bool {
	for _, c := range str {
		if c == char {
			return true
		}
	}
	return false
}

func space_control(statement string) string {
	punctuation := ",.!?;:'\"()-"
	result := ""

	for i := 0; i < len(statement); i++ {
		ispunctuation := false
		for j := 0; j < len(punctuation); j++ {
			if statement[i] == punctuation[j] {
				ispunctuation = true
				break
			}
		}

		if ispunctuation == true {
			result += string(statement[i])
			if i < len(statement)-1 && statement[i+1] == ' ' {
				i++
			}
		} else if statement[i] == ' ' {
			if i < len(statement)-1 && rune_control(punctuation, rune(statement[i+1])) {
				continue
			} else {
				result += string(statement[i])
			}
		} else {
			result += string(statement[i])
		}
	}

	return result
}

func punctuation_control(ch rune) bool {

	punctuation := []string{",", ".", "!", "?", ";", ":", "'", "\"", "(", ")", "-"}

	for _, i := range punctuation {

		if string(ch) == i {
			return true
		}
	}
	return false
}

func letter_control(ch rune) bool {

	if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
		return true
	}
	return false
}
func punctuation_edit(statement string) string {
	statement = space_control(statement)
	char_array := []rune(statement)
	output := []rune{}
	for i := 0; i < len(char_array)-1; i++ {
		if punctuation_control(char_array[i]) && letter_control(char_array[i+1]) {
			if char_array[i] == '\'' && letter_control(char_array[i+1]) || char_array[i] == '-' && letter_control(char_array[i+1]) || char_array[i] == '"' {
				output = append(output, char_array[i])
			} else {
				output = append(output, char_array[i], ' ')
			}
		} else if char_array[i] == ':' || char_array[i] == ';' && (char_array[i+1] == '\'' || char_array[i+1] == '"') {
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
