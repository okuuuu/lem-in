package antfarm

import (
	"bufio"
)

func Solve(scanner *bufio.Scanner) error {
	farm := PopulateFarm()
	var err error
	for scanner.Scan() {
		err = farm.ReadLine(scanner.Text())
		if err != nil {
			panic(err)
		}
	}
	err = farm.FindResult()
	if err != nil {
		panic(err)
	}
	return nil
}
