package main

import (
	"fmt"
	"fzero/luckyranks/ranks"
	"math/rand"
	"strings"
)

func iterate(resultCounter *LuckyRankResults) {
	yourRankDst := ranks.YourRankHand{}
	mysteryCardDst := ranks.MysteryCardHand{}
	// for each your rank
	for yourRankIndex := 0; yourRankIndex < ranks.TotalYourRankCombinations; yourRankIndex++ {
		ranks.IndexToYourRank(&yourRankDst, yourRankIndex)
		// track number of digits
		uniqueDigits := ranks.UniqueDigits(yourRankDst[:])
		// for each mystery rank
		for mysteryRankIndex := 0; mysteryRankIndex < ranks.TotalMysteryRankCombinations; mysteryRankIndex++ {
			if mysteryRankIndex%1000000 == 0 {
				fmt.Printf("your rank progress %v/%v: mystery rank progress %v/%v\n", yourRankIndex, ranks.TotalYourRankCombinations, mysteryRankIndex, ranks.TotalMysteryRankCombinations)
			}
			ranks.IndexToMysteryRank(&mysteryCardDst, mysteryRankIndex)
			// track number of machinecards and get matches
			machineCards, match := computeMatches(mysteryCardDst, yourRankDst)
			countResults(resultCounter, uniqueDigits, machineCards, match)
		}
	}
}

func random(resultCounter *LuckyRankResults, n int) {
	yourRankDst := ranks.YourRankHand{}
	// if we want perfect placements, use the line below-- remember placements are off by one
	// yourRankDst := ranks.YourRankHand{0,22,44,66,88}
	mysteryCardDst := ranks.MysteryCardHand{}
	var progress int

	fmt.Println("--- jank progress bar ---")
	fmt.Println(strings.Repeat(".", 99))
	for i := 0; i < n; i++ {
		if i*100/n > progress {
			progress += 1
			fmt.Print("#")
		}
		
		yourRankIndex := rand.Intn(ranks.TotalYourRankCombinations)
		ranks.IndexToYourRank(&yourRankDst, yourRankIndex)

		mysteryRankIndex := rand.Intn(ranks.TotalMysteryRankCombinations)
		ranks.IndexToMysteryRank(&mysteryCardDst, mysteryRankIndex)
		// track number of digits
		uniqueDigits := ranks.UniqueDigits(yourRankDst[:])
		// track number of machinecards and get matches
		machineCards, match := computeMatches(mysteryCardDst, yourRankDst)
		countResults(resultCounter, uniqueDigits, machineCards, match)
	}
	fmt.Println()
}

func main() {
	resultCounter := LuckyRankResults{}

	// TODO use os args to decide between iterate and random
	// iterate(resultCounter, yourRankDst, mysteryCardDst)
	random(&resultCounter, 1000*1000*50)
	fmt.Println()
	readResults(resultCounter)
}
