package day_2

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type command struct {
	start  int
	end    int
}

func ReadInput() []command {
	// file, _ := os.Open("day_2/test.txt")
	file, _ := os.Open("day_2/input.txt")
	defer file.Close()

	fmt.Println("file", file)

	var commands []command
	scanner := bufio.NewScanner(file)

	if !scanner.Scan() {
		fmt.Println("Файл пуст или ошибка чтения")
		return commands
	}

	splitLine := strings.Split(scanner.Text(), ",")

	for _, value := range splitLine {
		value = strings.TrimSpace(value)
		if value == "" {
			continue
		}

		splitValue := strings.Split(value, "-")
		if len(splitValue) < 2 {
			fmt.Printf("Предупреждение: некорректный формат '%s', пропускаем\n", value)
			continue
		}

		start, _ := strconv.Atoi(splitValue[0])
		end, _ := strconv.Atoi(splitValue[1])

		commands = append(commands, command{
			start: start,
			end:   end,
		})
	}

	return commands
}

func Day2Star1() int {
	day2Input := ReadInput()

	res := 0

	for _, command := range day2Input {
		fmt.Println(command)

		for i := command.start; i <= command.end; i++ {
			length := len(strconv.Itoa(i))

			left := strconv.Itoa(i)[:length/2]
			right := strconv.Itoa(i)[length/2:]

			if length % 2 != 0 {
				continue
			}

			if left == right {
				res += i
			}
		}
	}

	fmt.Println("res",	res)

	return res
}
