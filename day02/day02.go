package day02

import (
	"fmt"
	"strings"
	"strconv"
	"advent-of-code-2024/util"
)

func FindSafeReports(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	reports := parseStringTo2dIntArray(data)
	validReports := countValidReports(reports)
	fmt.Printf("Safe Reports: %d\n", validReports)
}

func FindSafeReportsWithDampener(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	reports := parseStringTo2dIntArray(data)
	validReports := countValidReportsWithDampener(reports)
	fmt.Printf("Safe Reports With Dampener: %d\n", validReports)
}

func parseStringTo2dIntArray(str string) [][]int {
	list := [][]int{}

	lines := strings.Split(str, "\n")

	for _, line := range lines {
		asChars := strings.Split(line, " ")
		asInts := []int{}

		for _, c := range asChars {
			n, _ := strconv.Atoi(c)
			asInts = append(asInts, n)
		}
		
		list = append(list, asInts)
	}
	return list
}

func countValidReports(reports [][]int) int {
	valid := 0

	for _, report := range reports {
		isValidIncreasingReport := checkAllIncreasing(report)
		isValidDecreasingReport := checkAllDecreasing(report)

		if isValidIncreasingReport || isValidDecreasingReport {
			valid += 1
		}
	}
	
	return valid
}

func countValidReportsWithDampener(reports [][]int) int {
	valid := 0

	for _, report := range reports {
		isValidIncreasingReport := dampener(checkAllIncreasing, report)
		isValidDecreasingReport := dampener(checkAllDecreasing, report)

		if isValidIncreasingReport || isValidDecreasingReport {
			valid += 1
		}
	}

	return valid
}

func checkAllIncreasing(list []int) bool {
	isOkay := true

	for i, _ := range list {
		if i >= len(list) -1 {
			break
		}
		if list[i+1] <= list[i] {
			isOkay = false
		} else if list[i+1] - list[i] > 3 {
			isOkay = false
		}
	}

	return isOkay
}


func checkAllDecreasing(list []int) bool {
	isOkay := true

	for i, _ := range list {
		if i >= len(list) - 1 {
			break
		}
		if list[i+1] >= list[i] {
			isOkay = false
		} else if list[i] - list[i+1] > 3 {
			isOkay = false
		}
	}

	return isOkay
}

func dampener(checkFn func(list []int) bool, list []int) bool {
	okay := checkFn(list)

	if !okay {
		for i, _ := range list {
			if i > len(list) - 1 {
				break
			}
			isValidInterval := checkFn([]int{list[i], list[i+1]})
			if !isValidInterval {
				listCopy := make([]int, len(list))
				copy(listCopy, list)
				isValidWithoutCurrentElement := checkFn(append(listCopy[:i], listCopy[i+1:]...))
				if isValidWithoutCurrentElement {
					okay = true
				} else {
					listCopy2 := make([]int, len(list))
					copy(listCopy2, list)
					isValidWithoutNextElement := checkFn(append(listCopy2[:i+1], listCopy2[i+2:]...))
					if isValidWithoutNextElement {
						okay = true
					}
				}
				break
			}
		}
	}

	return okay
}
