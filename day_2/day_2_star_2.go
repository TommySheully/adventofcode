package day_2

import (
	"fmt"
	"strconv"
)

func Day2Star2() int {
	day2Input := ReadInput()

	res := 0

	for _, command := range day2Input {
		fmt.Println(command)

		for i := command.start; i <= command.end; i++ {
			iStr := strconv.Itoa(i)
			length := len(iStr)

			for groupSize := 1; groupSize <= length/2; groupSize++ {

				if length % groupSize == 0 {
					group := iStr[:groupSize]

					g := groupSize 

					for  g < length {
						if iStr[g:g+groupSize] != group {
							break
						}

						g += groupSize	
					}

					if g >= length {
						res += i
						fmt.Println("+++i", i)
						fmt.Println("res", res)
						break 
					}	
				}
			}	
		}
	}

	fmt.Println("res",	res)

	return res
}
