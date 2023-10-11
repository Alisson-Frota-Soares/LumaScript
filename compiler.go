package main

import (
	"bufio"
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

func main() {

	// Inicialize a variável global com um slice vazio no início do programa
	global_vars = make([]variable, 0)
	//global_vars = append(global_vars, variable{Name: "Var1", Value: "Value1", type_var: "Type1"})

	file, err := os.Open("LumaScript.lum")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {

		line_splitted := splitter(scanner.Text())

		executor(line_splitted)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(global_vars)

}

func executor(line_splitted []string) {

	if len(line_splitted) > 0 {

		switch line_splitted[0] {
		case "print":
			print(line_splitted)
		case "var":
			createVar(line_splitted)
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

func splitter(input string) []string {
	// Use uma expressão regular para encontrar frases dentro de aspas
	re := regexp.MustCompile(`"[^"]+"`)

	// Substitua as frases dentro de aspas por um marcador temporário
	tempMarker := "_temp_marker_"
	input = re.ReplaceAllStringFunc(input, func(s string) string {
		s = strings.Trim(s, `"`) // Remove as aspas
		return strings.Replace(s, " ", tempMarker, -1)
	})

	// Divida a string em palavras usando espaços
	words := strings.Fields(input)

	// Restaure as frases dentro de aspas substituindo o marcador temporário por espaços
	for i, word := range words {
		words[i] = strings.Replace(word, tempMarker, " ", -1)
	}

	return words
}
