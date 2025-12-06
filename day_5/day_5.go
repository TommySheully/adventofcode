package day_5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	line  string
}

func ReadInput() []command {
	// file, _ := os.Open("day_5/test.txt")
	file, _ := os.Open("day_5/input.txt")
	defer file.Close()

	var commands []command
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			commands = append(commands, command{
				line: line,
			})
		}
	}

	return commands
}

type Range struct {
	start int
	end   int
}

type formatInput struct {
	ingredients []int
	ranges      []Range
}

func (f formatInput) isInRange(num int) bool {
	for _, r := range f.ranges {
		if num >= r.start && num <= r.end {
			return true
		}
	}
	return false
}

func FormatInput() formatInput {
	input := ReadInput()

	ingredients := []int{}
	ranges := []Range{}
	
	for _, command := range input {
		if strings.Contains(command.line, "-") {
			value := strings.Split(command.line, "-")

			start, _ := strconv.Atoi(value[0])
			end, _ := strconv.Atoi(value[1])

			ranges = append(ranges, Range{
				start: start,
				end:   end,
			})
		}

		if !strings.Contains(command.line, "-") {
			value, _ := strconv.Atoi(command.line)
			ingredients = append(ingredients, value)
		}
	}

	return formatInput{
		ingredients: ingredients,
		ranges:      ranges,
	}
}

func Day5Star1() int {
	formattedInput := FormatInput()

	res := 0

	for _, ingredient := range formattedInput.ingredients {
		if formattedInput.isInRange(ingredient) {
			res++
		}
	}

	fmt.Println("res", res)

	return res
}
