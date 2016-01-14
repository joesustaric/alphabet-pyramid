package main

import (
	"flag"
	"fmt"
	"regexp"
	"strings"
)

var letter *string
var alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

const (
	inputStringSize = 1
	diamond         = `
{{.First}}
{{.Second}} {{.Second}}
{{.Third}} {{.Fourth}}
{{.Fifth}} {{.Fifth}}
{{.Sixth}} {{.Sixth}}
{{.Fifth}} {{.Fifth}}
{{.Third}} {{.Fourth}}
{{.Second}} {{.Second}}
{{.First}}
`
)

type Letters struct {
	First, Second, Third, Fourth, Fifth string
}

func NewLetters(first, second, third, fourth, fifth string) Letters {
	return Letters{First: first, Second: second, Third: third, Fourth: fourth, Fifth: fifth}
}

func init() {
	letter = flag.String("letter", "", "The letter which is the middle of the diamond.")
	flag.Parse()
}

func main() {

	if input, err := Parse(*letter); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Printf("letter = %s \n", input)
	}

}

func Parse(input string) (string, error) {
	r := strings.TrimSpace(input)

	if notCorrectLength(r) {
		return "", fmt.Errorf("input incorrect %s", r)
	} else if notAlphabetCharacter(r) {
		return "", fmt.Errorf("not a-z or A-Z character")
	}

	return strings.ToUpper(r), nil
}

func GetDiamondLetters(seed string) Letters {

	r := strings.Split(alphabet, seed)

	var letters []string
	count := 0
	letters = append(letters, seed)

	for i := len(r[0]) - 1; i >= 0; i-- {
		letters = append(letters, string(r[0][i]))
		count++
		if count == 4 {
			break
		}
	}

	if len(letters) != 4 {
		for i := len(r[1]) - 1; i > 0; i-- {
			letters = append(letters, string(r[1][i]))
			count++
			if count == 4 {
				break
			}

		}
	}

	return NewLetters(letters[4], letters[3], letters[2], letters[1], letters[0])

}

func notCorrectLength(input string) bool {
	if len(input) != inputStringSize {
		return true
	}
	return false
}

func notAlphabetCharacter(input string) bool {
	matched, err := regexp.MatchString("[a-z|A-Z]", input)
	if !matched || err != nil {
		return true
	}
	return false
}
