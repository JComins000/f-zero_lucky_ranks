package ranks

import (
	"gonum.org/v1/gonum/stat/combin"
)

type YourRankHand [TotalYourRanks]int

func IndexToYourRank(dst *YourRankHand, idx int) {
	// use copy to "cast" int slice to YourRankDeck
	copy(dst[:], combin.IndexToCombination(nil, idx, totalPlacements, TotalYourRanks))
}

// Accepts placements (0 .. 98), converts them to real placements (1 .. 99), and counts the digits
func UniqueDigits(yourRanks []int) int {
	digitMatches := make(map[int]bool)
	for _, rank := range yourRanks {
		// Special case for leading Zeros
		if rank < 10 {
			digitMatches[0] = true
		}
		// I thought about using variables for 1s digit and 10s digit, but I thought this would be more fun.
		for digitCompare := rank + 1; digitCompare != 0; digitCompare /= 10 {
			digitMatches[digitCompare%10] = true
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
