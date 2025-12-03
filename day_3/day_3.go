package day_3

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	line  string
}

func ReadInput() []command {
	// file, _ := os.Open("day_3/test.txt")
	file, _ := os.Open("day_3/input.txt")
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

func Day3Star1() int {
	day3Input := ReadInput()

	res := 0

	for _, command := range day3Input {
		line := command.line
		max := 0
	
		for i := 0; i < len(line); i++ {
			for j := 0; j < len(line); j++ {

				if i < j {
					current := int(line[i] - '0') * 10 + int(line[j] - '0')

					if current > max {
						max = current
					}
				}
			}
		}
		res += max
	}

	fmt.Println("day3Input", res)

	return res
}
