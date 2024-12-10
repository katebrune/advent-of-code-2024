package day04

import (
	"fmt"
	"regexp"
	"strings"
	"advent-of-code-2024/util"
)

func ProcessWordSearch(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	all := []string{}

	horizontalL2R := parseHorizontalL2R(data)
	all = append(all, horizontalL2R...)

	horizontalR2L := parseHorizontalR2L(data)
	all = append(all, horizontalR2L...)

	verticalT2B := parseVerticalT2B(data)
	all = append(all, verticalT2B...)

	verticalB2T := parseVerticalB2T(data)
	all = append(all, verticalB2T...)

	diagonalTL2BR := parseDiagonalTL2BR(data)
	all = append(all, diagonalTL2BR...)

	diagonalTR2BL := parseDiagonalTR2BL(data)
	all = append(all, diagonalTR2BL...)

	diagonalBL2TR := parseDiagonalBL2TR(data)
	all = append(all, diagonalBL2TR...)

	diagonalBR2TL := parseDiagonalBR2TL(data)
	all = append(all, diagonalBR2TL...)

	occurances := 0

	for _, line := range all {
		count := countXmasOccurances(line)

		occurances += count
	}

	fmt.Printf("XMAS Occurances: %d\n", occurances)
}

func ProcessXWordSearch(fileName string) {
	data, err := util.ReadFileToString(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	list := parseStringToLetterAs(data)

	sum := 0

	for _, item := range list {
		count := countMasInLetterA(item)
		if count == 2 {
			sum += 1
		}
	}

	fmt.Printf("X Occurances: %d\n", sum)
}

func countXmasOccurances(list string) int {
	xmas := regexp.MustCompile("XMAS")
	matches := xmas.FindAllString(list, -1)

	return len(matches)
}

func parseHorizontalL2R(str string) []string {
	lines := strings.Split(str, "\n")

	return lines
}

func parseHorizontalR2L(str string) []string {
	lines := parseHorizontalL2R(str)

	for i, line := range lines {
		lines[i] = reverseString(line)
	}

	return lines
}

func parseVerticalT2B(str string) []string {
	lines := parseHorizontalL2R(str)
	columns := []string{}

	for i, line := range lines {
		if i == 0 {
			for _, r := range line {
				columns = append(columns, string(r))
			}
		} else {
			for j, r := range line {
				columns[j] = columns[j] + string(r)
			}
		}
	}

	return columns
}

func parseVerticalB2T(str string) []string{
	columns := parseVerticalT2B(str)

	for i, column := range columns {
		columns[i] = reverseString(column)
	}

	return columns
}

func parseDiagonalTL2BR(str string) []string {
	lines := parseHorizontalL2R(str)
	diagonals := make([]string, len(lines)+len(lines[0]) - 1)

	for i, line := range lines {
		for j, r := range line {
			if i==0 {
				diagonals[j] = diagonals[j] + string(r)
			} else if j==0 {
				diagonals[len(lines[0]) - 1 + i] = diagonals[len(lines[0]) - 1 + i] + string(r) 
			} else {
				if i - j == 0 {
					diagonals[0] = diagonals[0] + string(r)
				} else if i-j < 0 {
					index := j-i
					diagonals[index] = diagonals[index] + string(r)
				} else {
					index := len(lines[0]) - 1 + i - j
					diagonals[index] = diagonals[index] + string(r)
				}
			}
		}
	}

	return diagonals
}

func parseDiagonalBR2TL(str string) []string {
	diagonals := parseDiagonalTL2BR(str)

	for i, line := range diagonals {
		diagonals[i] = reverseString(line)
	}

	return diagonals
}

func parseDiagonalTR2BL(str string) []string {
	lines := parseHorizontalR2L(str)
	joined := strings.Join(lines, "\n")
	diagonals := parseDiagonalTL2BR(joined)

	return diagonals
}

func parseDiagonalBL2TR(str string) []string {
	diagonals := parseDiagonalTR2BL(str)

	for i, line := range diagonals {
		diagonals[i] = reverseString(line)
	}

	return diagonals
}

func reverseString(str string) string {
	runes := []rune(str)

    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }

    return string(runes)
}

type LetterA struct {
	tl rune
	tr rune
	bl rune
	br rune
}

func parseStringToLetterAs(str string) []LetterA {
	lines := strings.Split(str, "\n")
	runes := [][]rune{}

	for _, line := range lines {
		runes = append(runes, []rune(line))
	}

	list := []LetterA{}

	for i, line := range runes {
		if i == 0 || i == len(runes) - 1 {
			continue
		}
		for j, r := range line {
			if j == 0 || j == len(line) - 1 {
				continue
			}
			if r == 'A' {
				item := LetterA{
					tl: runes[i-1][j-1],
					tr: runes[i-1][j+1],
					bl: runes[i+1][j-1],
					br: runes[i+1][j+1],
				}
				list = append(list, item)
			}
		}
	}

	return list
}

func countMasInLetterA(a LetterA) int {
	count := 0

	if a.tl == 'M' && a.br == 'S' {
		count += 1
	}

	if a.tl == 'S' && a.br == 'M' {
		count += 1
	}

	if a.tr == 'M' && a.bl == 'S' {
		count += 1
	}

	if a.tr == 'S' && a.bl == 'M' {
		count += 1
	}

	return count
}