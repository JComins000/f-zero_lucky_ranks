package main

import (
	"fmt"
	"fzero/luckyranks/ranks"
)

func compareCompareMysteryWithYourRank(mysteryCards []ranks.MysteryCard, yourPlacements []int) int {
	matches := 0
	for _, card := range mysteryCards {
		for _, placement := range yourPlacements {
			if card.Match(placement) {
				matches += 1
				break
			}
		}
	}
	return matches
}

func main() {
	// not done yet
	// TODO, test mystery ranks, continue iterating over everything
	fmt.Println(ranks.TotalYourRankCombinations)
	yourRankDst := make([]int, ranks.TotalYourRanks)
	ranks.IndexToYourRank(yourRankDst, 12)
	fmt.Println(yourRankDst)

	mysteryCardDst := make([]ranks.MysteryCard, ranks.TotalMysteryCards)
	ranks.IndexToMysteryRank(mysteryCardDst, 14)
	fmt.Println(mysteryCardDst)
	fmt.Println(compareCompareMysteryWithYourRank(mysteryCardDst, yourRankDst))

	fmt.Println(ranks.UniqueDigits(yourRankDst))
	fmt.Println(ranks.UniqueDigits([]int{1, 2, 3, 4, 5}))
	fmt.Println(ranks.UniqueDigits([]int{11, 27, 38, 49, 5}))
}
