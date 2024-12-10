package main

import (
	"fmt"
	"advent-of-code-2024/day01"
	"advent-of-code-2024/day02"
	"advent-of-code-2024/day03"
	"advent-of-code-2024/day04"
)

func main() {
	runDay04()
}

func runDay01(){
	fmt.Println("--- Day 01 ---")
	day01.ProcessDifferences("../day01/data.txt")
	day01.ProcessSimilarities("../day01/data.txt")
}

func runDay02(){
	fmt.Println("--- Day02 ---")
	day02.FindSafeReports("../day02/data.txt")
	day02.FindSafeReportsWithDampener("../day02/data.txt")
}

func runDay03(){
	fmt.Println("--- Day03 ---")
	day03.ComputeMultiplications("../day03/data.txt")
	day03.ComputeMultiplicationsWithDoDont("../day03/data.txt")
}

func runDay04(){
	fmt.Println("--- Day04 ---")
	day04.ProcessWordSearch("../day04/data.txt")
	day04.ProcessXWordSearch("../day04/data.txt")
}