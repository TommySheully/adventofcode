package day_5_star_2

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func FormatInput() []Range {
	input := ReadInput()

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
	}

	return ranges
}

func Day5Star2() int {
	ranges := FormatInput()

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].start < ranges[j].start
	})

	res := []Range{}

	current := ranges[0]

	for i := 1; i < len(ranges); i++ {
		next := ranges[i]

		if next.start <= current.end + 1 {
			if next.end > current.end {
				current.end = next.end
			}
		} else {
			res = append(res, current)
			current = next
		}
	}

	res = append(res, current)

	total := 0
	for _, r := range res {
		total += r.end - r.start + 1
	}

	fmt.Println("merged ranges:", res)
	fmt.Println("total count:", total)

	return total
}
