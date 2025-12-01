package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Command struct {
	position string // "L" or "R"
	steps    int
}

func readInput() []Command {
	// file, err := os.Open("day_1/test.txt")
	file, err := os.Open("day_1/input.txt")
	if err != nil {
		fmt.Printf("Ошибка при открытии файла: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	var commands []Command
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		line := scanner.Text()

		if line != "" {
			steps, err := strconv.Atoi(line[1:])
			if err != nil {
				fmt.Printf("Ошибка при преобразовании строки в число: %v\n", err)
				os.Exit(1)
			}

			commands = append(commands, Command{
				position: string(line[0]),
				steps:    steps,
			})
		}
	}
	
	if err := scanner.Err(); err != nil {
		fmt.Printf("Ошибка при чтении файла: %v\n", err)
		os.Exit(1)
	}
	
	return commands
}

func main() {
	start := 50
	commands := readInput()

	res := 0
	
	if len(commands) > 0 {
		for _, command := range commands {
			fmt.Println(command.position, command.steps)

			if command.position == "L" {
				start -= command.steps % 100

				if start <= -100 {
					start = 100 + start
				}

				if start == 0 {
					res += 1
				}
			}

			if command.position == "R" {
				start += command.steps  % 100

				if start >= 100 {
					start = start - 100
				}

				if start == 0 {
					res += 1
				}
			}
			fmt.Println("start", start)
   	}

	}

	fmt.Println("res",res)
}