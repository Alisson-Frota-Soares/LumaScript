package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type variable struct {
	Name     string
	Value    string
	type_var string
}

var global_vars []variable

type token struct {
	Value     string
	TokenType string
}

func main() {

	// Inicialize a variável global com um slice vazio no início do programa
	global_vars = make([]variable, 0)
	//global_vars = append(global_vars, variable{Name: "Var1", Value: "Value1", type_var: "Type1"})

	file, err := os.ReadFile("LumaScript.lum")
	if err != nil {
		log.Fatal(err)
	}

	// defer file.Close()

	tokens := tokenizer(string(file))

	fmt.Println(tokens)

	// read the file line by line using scanner
	// scanner := bufio.NewScanner(file)

	// for scanner.Scan() {

	// 	tokens := tokenizer(scanner.Text())

	// 	fmt.Println(tokens)

	// 	executor(tokens)
	// }

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }

	//fmt.Println(global_vars)

}

func executor(tokens []string) {

	if len(tokens) > 0 {

		switch tokens[0] {
		case "print":
			print(tokens)
		case "var":
			createVar(tokens)
		default:
			fmt.Println("unexpected sintax")
		}

	}
}

func get_var_value(name string) string {
	for _, v := range global_vars {
		if v.Name == name {
			return v.Value
		}
	}

	return "no var found"
}

func createVar(line_splitted []string) {
	var_name := line_splitted[2]
	var_value := line_splitted[4]
	var_type := line_splitted[1]
	if !hasMultipleWords(line_splitted[1]) {

		if line_splitted[3] == "=" {
			global_vars = append(global_vars, variable{Name: var_name, Value: var_value, type_var: var_type})

		}

	}

}

func print(line_splitted []string) {
	if hasMultipleWords(line_splitted[1]) {
		fmt.Println(line_splitted[1])
	} else {
		fmt.Println(get_var_value(line_splitted[1]))
	}
}

func hasMultipleWords(input string) bool {
	// Use strings.Fields para dividir a string em palavras
	words := strings.Fields(input)

	// Verifique o número de palavras
	return len(words) > 1
}

func tokenizer(input string) []token {
	re := regexp.MustCompile(`"([^"\\]*(\\.[^"\\]*)*)"`)
	tempMarker := "_temp_marker_"
	input = re.ReplaceAllStringFunc(input, func(s string) string {
		//s = strings.Trim(s, `"`)
		s = strings.Replace(s, `\"`, `"`, -1)
		return strings.Replace(s, " ", tempMarker, -1)
	})

	lines := strings.Split(input, "\n")
	var tokens []token

	for _, line := range lines {
		if strings.HasPrefix(line, "//") {
			continue
		}

		words := strings.Fields(line)

		for i, word := range words {
			words[i] = strings.Replace(word, tempMarker, " ", -1)
		}

		for _, word := range words {
			var tokenType string

			switch {
			case word == "print":
				tokenType = "print"
			case word == "string" || word == "int" || word == "var":
				tokenType = "type"
			case word == "=":
				tokenType = "assignment"
			case strings.HasPrefix(word, `"`) && strings.HasSuffix(word, `"`):
				tokenType = "string"
			case isInteger(word):
				tokenType = "integer"
			case isFloat(word):
				tokenType = "float"
			case isIdentifier(word):
				tokenType = "identifier"
			case isMathExpression(word):
				tokenType = "math_expression"
			case word == "(" || word == ")":
				tokenType = "parenthesis"
			default:
				tokenType = "undefined"
			}

			tokens = append(tokens, token{Value: word, TokenType: tokenType})
		}
	}

	return tokens
}

func isInteger(s string) bool {
	_, err := fmt.Sscanf(s, "%d")
	return err == nil
}

func isFloat(s string) bool {
	_, err := fmt.Sscanf(s, "%f")
	return err == nil
}

func isIdentifier(s string) bool {
	return regexp.MustCompile(`^[a-zA-Z_]\w*$`).MatchString(s)
}

func isMathExpression(s string) bool {
	return regexp.MustCompile(`^[+\-*/=<>]+$`).MatchString(s)
}
