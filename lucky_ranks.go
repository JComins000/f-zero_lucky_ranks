package main

import (
	// "bufio"
	"fmt"
	"fzero/luckyranks/ranks"
	// "os"
)

func computeMatches(mysteryCards []ranks.MysteryCard, yourPlacements []int) (int, int) {
	matches := 0
	machines := 0
	// order matters here. we care how many cards are matched
	for _, card := range mysteryCards {
		switch card.(type) {
		case ranks.MachineCard:
			machines += 1
		}
		for _, placement := range yourPlacements {
			if card.Match(placement) {
				// when a placement matches a card, we count the card as matched and break to the next card
				matches += 1
				break
			}
		}
	}
	return machines, matches
}

func main() {
	// TODO: cut down on total combos using math. Maybe compute one section of the chart at a time

	// indexed by [unique digits - 2][machine cards present][5 or greater matches]
	// there will always be at least 2 unique digits for any 5 placements.
	// for example, (01, 02, 12, 21, 11) has two unique digits (0s dont count)
	resultCounter := [8][ranks.TotalMachines + 1][2]int{}

	yourRankDst := make([]int, ranks.TotalYourRanks)
	mysteryCardDst := make([]ranks.MysteryCard, ranks.TotalMysteryCards)

	// for each your rank
	for yourRankIndex := 0; yourRankIndex < ranks.TotalYourRankCombinations; yourRankIndex++ {
		ranks.IndexToYourRank(yourRankDst, yourRankIndex)
		// track number of digits
		uniqueDigits := ranks.UniqueDigits(yourRankDst)
		// for each mystery rank
		for mysteryRankIndex := 0; mysteryRankIndex < ranks.TotalMysteryRankCombinations; mysteryRankIndex++ {
			if mysteryRankIndex % 1000000 == 0 {
				fmt.Printf("your rank progress %v/%v: mystery rank progress %v/%v\n", yourRankIndex, ranks.TotalYourRankCombinations, mysteryRankIndex, ranks.TotalMysteryRankCombinations)
			}
			ranks.IndexToMysteryRank(mysteryCardDst, mysteryRankIndex)
			// track number of machinecards and get matches
			machineCards, matches := computeMatches(mysteryCardDst, yourRankDst)
			var gotFiveMatches int8
			if matches >= 5 {
				gotFiveMatches = 1
			}
			// convert to indices and store value
			resultCounter[uniqueDigits-2][machineCards][gotFiveMatches] += 1
		}
	}
}
