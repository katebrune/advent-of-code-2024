package day01

import(
	"fmt"
	"strings"
	"strconv"
	"sort"
	"advent-of-code-2024/util"
)

func ProcessDifferences(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	twoLists := parseStringToTwoLists(data)
	orderedLists := orderTwoLists(twoLists)
	differences := getDifferencesOfTwoLists(orderedLists)
	sum := sumList(differences)
	fmt.Printf("differences: %d\n", sum)
}

func ProcessSimilarities(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
	}
	twoLists := parseStringToTwoLists(data)
	similarities := getSimilaritiesOfTwoLists(twoLists)
	sum := sumList(similarities)
	fmt.Printf("similarities: %d\n", sum)
}

type TwoLists struct {
	ListA []int
	ListB []int
}

func parseStringToTwoLists(str string) TwoLists {
	listA := []int{}
	listB := []int{}

	lines := strings.Split(str, "\n")
	
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}

		numbers := strings.Split(line, "  ")
		for i, part := range numbers {
			numbers[i] = strings.TrimSpace(part)
		}
	
		numberA, _ := strconv.Atoi(numbers[0])
		numberB, _ := strconv.Atoi(numbers[1])

		listA = append(listA, numberA)
		listB = append(listB, numberB)
	}
	return TwoLists {
		ListA: listA,
		ListB: listB,
	}
}

func orderTwoLists(lists TwoLists) TwoLists {
	sort.Ints(lists.ListA)
	sort.Ints(lists.ListB)
	return lists
}

func getDifferencesOfTwoLists(lists TwoLists) []int {
	differences := []int{}
	for i, numberA := range lists.ListA {
		numberB := lists.ListB[i]

		difference := 0
		if numberA > numberB {
			difference = numberA - numberB
		} else {
			difference = numberB - numberA
		}

		differences = append(differences, difference)
	}
	return differences
}

func getSimilaritiesOfTwoLists(lists TwoLists) []int {
	similarities := []int{}

	for _, numberA := range lists.ListA {
		count := 0
		for _, numberB := range lists.ListB {
			if numberA == numberB {
				count += 1
			}
		}
		similarities = append(similarities, count * numberA)
	}
	return similarities
}

func sumList(list []int) int {
	sum := 0

	for _, num := range list {
		sum += num
	}

	return sum
}
