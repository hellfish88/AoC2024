package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type levelreport struct {
	levelsAmount int
	levels       []int
	safe         bool
}

func main() {
	var levelreports = []levelreport{}
	for _, report := range loadFile("./input.txt") {
		levelreports = append(levelreports, createLevelReport(report))
	}

	var total_safe int = 0

	for _, report := range levelreports {
		safeReport(&report)
		if report.safe {
			total_safe++
		}
	}
	fmt.Println(total_safe)
}

func loadFile(p string) []string {
	file, err := os.ReadFile(p)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return []string(strings.Split(string(file), "\n"))
}

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return i
}

func createLevelReport(s string) levelreport {
	returnReport := levelreport{}
	fields := strings.Fields(s)
	for _, v := range fields {
		returnReport.levels = append(returnReport.levels, atoi(v))
	}
	returnReport.levelsAmount = len(fields)

	returnReport.safe = false

	return returnReport
}

func safeReport(lr *levelreport) {

	if lr.allDecrement() || lr.allIncrement() {
		lr.safe = true
	} else {
		lr.safe = false
	}

}
func (lr *levelreport) allIncrement() bool {
	for i := 0; i < lr.levelsAmount-1; i++ {
		level := lr.levels[i]
		nextLevel := lr.levels[i+1]

		diff := nextLevel - level
		if diff < 1 || diff > 3 {
			return false
		}

		if nextLevel <= level {
			return false
		}
	}
	return true
}

func (lr *levelreport) allDecrement() bool {
	for i := 0; i < lr.levelsAmount-1; i++ {
		level := lr.levels[i]
		nextLevel := lr.levels[i+1]

		diff := level - nextLevel
		if diff < 1 || diff > 3 {
			return false
		}

		if nextLevel >= level {
			return false
		}
	}
	return true
}
