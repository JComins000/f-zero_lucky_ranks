package ranks

import (
	"gonum.org/v1/gonum/stat/combin"
)

func IndexToYourRank(dst []int, idx int) {
	combin.IndexToCombination(dst, idx, totalPlacements, TotalYourRanks)
}

// Accepts placements (0 .. 98), converts them to real placements (1 .. 99), and counts the digits
// also, 0 doesn't count as a digit since there is no "0" wild mystery card
func UniqueDigits(yourRanks []int) int {
	digitMatches := make(map[int]bool)
	for _, rank := range yourRanks {
		// I thought about using variables for 1s digit and 10s digit, but I thought this would be more fun.
		// I don't even need a special case for leading 0s since they don't count as digits.
		for digitCompare := rank + 1; digitCompare != 0; digitCompare /= 10 {
			digit := digitCompare % 10
			if digit != 0 {
				digitMatches[digitCompare % 10] = true
			}
		}
	}

	totalMatches := 0
	for _, digitMatch := range digitMatches {
		if digitMatch {
			totalMatches += 1
		}
	}
	return totalMatches
}

var TotalYourRankCombinations int = combin.Binomial(totalPlacements, TotalYourRanks)
