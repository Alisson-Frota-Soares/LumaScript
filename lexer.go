package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Token struct {
	Value     string
	TokenType string
}

func main() {
	file := getFile("LumaScript.lum")
	tokens := lexer(string(file))

	fmt.Println(tokens)
}

func lexer(input string) []Token {

	//input = strings.ReplaceAll(input, " ", "")
	input = removeCommentedLines(input)

	char := []rune(input)

	tokens := make([]Token, 0)

	for _, c := range char {
		c_string := string(c)
		var tokenType string
		switch {
		case c_string == "(" || c_string == ")":
			tokenType = ""
		}

		tokens = append(tokens, Token{Value: c_string, TokenType: tokenType})
	}

	return tokens
}

func removeCommentedLines(input string) string {
	lines := strings.Split(input, "\n")
	var resultLines []string

	for _, line := range lines {
		if !strings.HasPrefix(line, "//") {
			resultLines = append(resultLines, line)
		}
	}

	result := strings.Join(resultLines, "\n")
	return result
}

func getFile(path string) []byte {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	return file
}
