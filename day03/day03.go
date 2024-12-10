package day03

import (
	"fmt"
	"regexp"
	"strconv"
	"advent-of-code-2024/util"
)

type Arithmetic struct {
	numbers []int
	command string
}

func ComputeMultiplications(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	arithmeticList, _ := parseJumbledStringToArithmeticList(data)
	values := computeArthmeticList(arithmeticList)
	sum := sumList(values)
	fmt.Printf("Sum: %d\n", sum)
}

func ComputeMultiplicationsWithDoDont(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	list, _ := parseJumbledStringToArithmeticListWithDoDont(data)
	values := computeArthmeticList(list)
	sum := sumList(values)
	fmt.Printf("Sum with Do/Don't: %d\n", sum)
}

func parseJumbledStringToArithmeticList(str string) ([]Arithmetic, error) {
	pattern := "mul\\((\\d+),(\\d+)\\)"
	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return []Arithmetic{}, err
	}
	matches := re.FindAllString(str, -1)

	list := []Arithmetic{}

	for _, match := range matches {
		submatches := re.FindStringSubmatch(match)
		n1, _ := strconv.Atoi(submatches[1])
		n2, _ := strconv.Atoi(submatches[2])
		arithmetic := Arithmetic{
			numbers: []int{n1, n2},
			command: "mul",
		}
		list = append(list, arithmetic)
	}
	return list, nil
}

func parseJumbledStringToArithmeticListWithDoDont(str string) ([]Arithmetic, error) {
	result := []Arithmetic{}

	dontPattern := regexp.MustCompile("don't\\(\\)")
	doPattern := regexp.MustCompile("do\\(\\)")
	mulPattern := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")

	sections := dontPattern.Split(str, -1)
	
	for i, section := range sections {
		var matches []string
		if i == 0 {
			matches = mulPattern.FindAllString(section, -1)
		} else {
			dos := doPattern.Split(section, -1)
			matches = []string{}
			if len(dos) > 1 {
				for i, d := range dos {
					if (i == 0) {
						continue
					}
					submatches := mulPattern.FindAllString(d, -1)
					matches = append(matches, submatches...)
				}
			} 
		}
		for _, match := range matches {
			submatches := mulPattern.FindStringSubmatch(match)
			n1, _ := strconv.Atoi(submatches[1])
			n2, _ := strconv.Atoi(submatches[2])
			arithmetic := Arithmetic{
				numbers: []int{n1,n2},
				command: "mul",
			}
			result = append(result, arithmetic)
		}
	}
	
	return result, nil
}

func computeArthmeticList(list []Arithmetic) []int {
	result := []int{}

	for _, arithmetic := range list {
		computation := arithmetic.numbers[0] * arithmetic.numbers[1]
		result = append(result, computation)
	}

	return result
}

func sumList(list []int) int {
	sum := 0

	for _, num := range list {
		sum += num
	}

	return sum
}