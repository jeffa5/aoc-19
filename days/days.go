package days

import (
	"fmt"

	"github.com/jeffa5/aoc-19/days/day1"
	"github.com/jeffa5/aoc-19/days/day2"
)

var dayMap = map[int](func(int, string) string){
	1: day1.Day1,
	2: day2.Day2,
}

func Solve(d, p int) {
	base := "input/day"
	if d == 0 {
		for d, f := range dayMap {
			if p == 0 {
				fmt.Printf("Day %d Part 1: %s\n", d, f(1, base))
				fmt.Printf("Day %d Part 2: %s\n", d, f(2, base))
			} else {
				fmt.Printf("Day %d Part %d: %s\n", d, p, f(p, base))
			}
		}
		return
	}

	f, ok := dayMap[d]
	if !ok {
		fmt.Printf("Invalid day selection: %d %d\n", d, p)
		return
	}

	s := f(p, base)

	fmt.Println(s)
}
