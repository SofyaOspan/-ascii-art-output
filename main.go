package main

import (
	"ascii-art/internal/ascii"
	"ascii-art/pkg/file"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	banner := "standard"
	var input string
	var fileName string
	if len(os.Args) == 4 {
		input = os.Args[2]
		fileName = os.Args[1]
		banner = os.Args[len(os.Args)-1]
	} else if len(os.Args) == 3 {
		input = os.Args[2]
		fileName = os.Args[1]
	} else {
		log.Fatalf("ERROR: not enough arguments. Try using this template:  --output=<fileName.txt> 'something' standard")
	}

	input = strings.ReplaceAll(input, "\n", "\\n")
	for _, w := range input {
		if w < ' ' || w > '~' {
			log.Fatalf("ERROR: Non-ASCII character found: %q", w)
		}
	}

	fontFileName, err := getBanner(banner)
	if err != nil {
		fmt.Printf("Error: \"%s\"\n", err.Error())
		return
	}

	lines, err := file.ReadLine(fontFileName + ".txt")
	if err != nil {
		fmt.Printf("Error: \"%s\"\n", err.Error())
		return
	}

	maps, err := ascii.Create(lines)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	words := strings.Split(input, "\\n")

	var outputWriter *os.File
	if fileName != "" {
		fileName, err, ok := getFileName(fileName)
		if !ok {
			fmt.Printf("%s\n", err.Error())
			return
		}
		outputWriter, err = os.Create(fileName)
		if err != nil {
			fmt.Println(fileName)

			log.Fatalf("ERROR: Failed to create output file: %s", err)
		}
		defer outputWriter.Close()
	} else {
		outputWriter = os.Stdout
	}

	err = ascii.Print(maps, words, outputWriter)
	if err != nil {
		log.Fatalf("Error: %s\n", err.Error())
	}
}
//The getBanner function takes a string s as input and 
//returns a string representing the file path of the corresponding banner font, along with an error.
func getBanner(s string) (string, error) {
	fontsPath := "fonts/"
	switch s {
	case "shadow":
		return fontsPath + "shadow", nil
	case "standard":
		return fontsPath + "standard", nil
	case "thinkertoy":
		return fontsPath + "thinkertoy", nil
	}
	return "", errors.New("error: incorrect format of [BANNER], try using one of these: <standard>, <thinkertoy>, <shadow>")
}
//The getFileName function takes a string s as input and returns three values: a string representing the extracted file name, an error, and a boolean value.


func getFileName(s string) (string, error, bool) {
	if strings.HasPrefix(s, "--output=") && strings.HasSuffix(s, ".txt") {
		return s[9:], nil, true
	}
	return "", errors.New("error: incorrect format of [OPTION], try using this template: --output=<filename.txt>"), false
}
