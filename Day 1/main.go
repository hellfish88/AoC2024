package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

var lhs []int
var rhs []int

var returnsum int = 0

func main() {

	// fmt.Println(returnFile("./input.txt")[0])

	for _, val := range returnFile("./input.txt") {
		retval := strings.Fields(val)
		lhsval, _ := strconv.Atoi(retval[0])
		rhsval, _ := strconv.Atoi(retval[1])
		lhs = append(lhs, lhsval)
		rhs = append(rhs, rhsval)

	}
	sort.Ints(lhs)
	sort.Ints(rhs)

	for i := range len(lhs) {
		returnsum += diff(lhs[i], rhs[i])
	}

	fmt.Println(returnsum)
	fmt.Println(toccurances(lhs, rhs))
}

func returnFile(p string) []string {
	file, err := os.ReadFile(p)

	if err != nil {
		log.Fatal(err)
	}

	return []string(strings.Split(string(file), "\n"))
}

func diff(lhs, rhs int) int {
	if lhs < rhs {
		return rhs - lhs
	} else {
		return lhs - rhs
	}
}

func toccurances(lhs, rhs []int) int {
	var similarity_score int = 0

	for _, val := range lhs {
		similarity_score += val * occurances(val, rhs)
	}
	return similarity_score
}

func occurances(i int, arr []int) int {
	var occurances int = 0

	for _, val := range arr {
		if i == val {
			occurances++
		}
	}
	return occurances
}
