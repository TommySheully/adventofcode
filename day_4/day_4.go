package day_4

import (
	"bufio"
	"fmt"
	"os"
)

type command struct {
	line  string
}

func ReadInput() []command {
	// file, _ := os.Open("day_4/test.txt")
	file, _ := os.Open("day_4/input.txt")
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

func Day4Star1() int {
	day4Input := ReadInput()

	res := 0
	rows := len(day4Input)

	for i := 0; i < rows; i++ {
		line := day4Input[i].line
			lineLength := len(line)

			fmt.Println("line", line)

			for j := 0; j < lineLength; j++ {
				current := line[j]

				if current != '@' {
					continue
				}

				count := 0

				for rows_index := -1; rows_index <= 1; rows_index++ {

					for cols_index := -1; cols_index <= 1; cols_index++ {
						if rows_index == 0 && cols_index == 0 {
							continue
						}

						newRowIndex := i + rows_index 
						newColIndex := j + cols_index 

						if newRowIndex >= 0 && newRowIndex < rows && newColIndex >= 0 && newColIndex < lineLength {
							if day4Input[newRowIndex].line[newColIndex] == '@' {
								count++
							}
						}
					}
				}

				if count < 4 {
					res++
				}
			}
	}

	fmt.Println("day4Input", res)

	return res
}
