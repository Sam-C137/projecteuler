package names_scores

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func Run() int {
	file, err := os.Open("0022_names.txt")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while opening file: %s", err)
		os.Exit(-1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var names []string

	if scanner.Scan() {
		for _, name := range strings.Split(scanner.Text(), ",") {
			name = strings.Trim(name, "\"")
			if name != "" {
				names = append(names, name)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while scanning file: %s", err)
		os.Exit(-1)
	}

	sort.Strings(names)
	total := 0

	for i, name := range names {
		score := 0
		for _, c := range name {
			score += int(c) - 64
		}
		total += score * (i + 1)
	}

	return total
}
