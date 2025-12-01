package day_1

import (
	"fmt"
)

func Day1Star2() {
	start := 50
	commands := ReadInput()

	res := 0

	if len(commands) > 0 {
		for _, command := range commands {
			rotations := command.steps / 100
			steps := command.steps % 100

			res += rotations

			if command.position == "L" {
				// будет ли оно пересекать 0 сука
				hasRotation := steps > 0 && start > 0 && start <= steps

				newStart := (start - steps + 100) % 100

				newStartnoZero := newStart == 0

				if start == 0 {
					res += 1
				}

				if newStartnoZero && start != 0 && !hasRotation {
					res += 1
				}

				if hasRotation && !newStartnoZero {
					res += 1
				}

				start = newStart
			}

			if command.position == "R" {
				if start == 0 {
					res += 1
				}

				hasRotation := steps > 0 && start+steps >= 100

				newStart := (start + steps) % 100

				newStartnoZero := newStart == 0

				if newStartnoZero && start != 0 && !hasRotation {
					res += 1
				}

				if hasRotation && !newStartnoZero {
					res += 1
				}

				start = newStart
			}
		}
	}

	fmt.Println("res", res)
}
