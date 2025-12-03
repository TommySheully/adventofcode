package day_3

import (
	"fmt"
	"strconv"
)

func Day3Star2() int {
	day3Input := ReadInput()

	res := 0

	for _, command := range day3Input {
		line := command.line

		needDelete := len(line) - 12
		var stack []int

		for i := 0; i < len(line); i++ {
			current := int(line[i] - '0')

			for needDelete > 0 && len(stack) > 0 && stack[len(stack)-1] < current {
				stack = stack[:len(stack)-1]
				needDelete--
			}

			stack = append(stack, current)
		}

		for needDelete > 0 {
			stack = stack[:len(stack)-1]
			needDelete--
		}


		result := ""

		for _, val := range stack {
			result += strconv.Itoa(val)
		}
		intValue, _ := strconv.Atoi(result)
		res += intValue
	}

	fmt.Println("day3Input", res)

	return res
}
