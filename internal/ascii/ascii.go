package ascii

import (
	"fmt"
	"os"
)

func Create(content []string) (map[rune][]string, error) {
	if len(content) != 855 {
		return nil, fmt.Errorf("non-valid number of lines")
	}
	maps := make(map[rune][]string)
	element := rune(32)
	previous, current := 1, 9
	for current <= len(content) {
		maps[element] = content[previous:current]
		element++
		previous = current + 1
		current += 9
	}
	return maps, nil
}

func Print(maps map[rune][]string, words []string, outputWriter *os.File) error {
	for _, word := range words {
		if word == "" {
			fmt.Fprintln(outputWriter)
			break
		}
		for i := 0; i < 8; i++ {
			for _, key := range word {
				fmt.Fprint(outputWriter, maps[key][i])
			}
			if i != 8 {
				fmt.Fprintln(outputWriter)
			}
		}
	}
	return nil
}
