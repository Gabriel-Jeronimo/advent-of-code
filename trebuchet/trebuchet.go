package trebuchet

import (
	"bufio"
	"os"
	"strconv"
	"unicode"
)

func findDigits(line []rune) int {
	result := ""

	for i := 0; i < len(line); i++ {
		if unicode.IsDigit(line[i]) {
			result = result + string(line[i])
			break
		}
	}

	for i := len(line) - 1; i >= 0; i-- {
		if unicode.IsDigit(line[i]) {
			result = result + string(line[i])
			break
		}
	}

	value, err := strconv.Atoi(result)

	if err != nil {
		panic(err)
	}

	return value
}

// Why os.readfile don't work? What's the difference?

func GetSum(path string) (int, error) {
	file, err := os.Open(path)
	if err != nil {
		// Handle the error
		return -1, err
	}
	defer file.Close()

	var sum int = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sum += findDigits([]rune(scanner.Text()))
	}

	return sum, nil
}