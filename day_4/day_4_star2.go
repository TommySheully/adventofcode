package day_4

import (
	"fmt"
)

func Day4Star2() int {
	day4Input := ReadInput()

	res := 0
	rows := len(day4Input)

  for {
	  found := false

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

					day4Input[i].line = day4Input[i].line[:j] + "x" + day4Input[i].line[j+1:]

					found = true
				}
			}
	}

	if !found {
	  break
	}
  }

	fmt.Println("Day4Star2", res)

	return res
}
