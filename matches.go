package main

import (
	"fmt"
	"fzero/luckyranks/ranks"
)

const possibleDigits = 8
const leastPossibleDigits = 3
type LuckyRankResults [possibleDigits][5][2]int

func countResults(results *LuckyRankResults, digits int, machines int, match bool) {
	// indexed by [unique digits - 3][machine cards present][1 or 0 representing 5 matches or not]
	// there will always be at least 2 unique digits for any 5 placements.
	// for example, (11, 12, 21, 22, x) has three unique digits
	var gotFiveMatches int8
	if match {
		gotFiveMatches = 1
	}
	// convert to indices and store value
	results[digits-3][machines][gotFiveMatches] += 1
}

func readResults(results LuckyRankResults) {
	var totalMatch, totalNoMatch int
	var machineMatch, machineNoMatch [ranks.TotalMachines + 1]int
	var digitMatch, digitNoMatch [possibleDigits]int

	var printPercentage = func(a int, b int) {
		fmt.Printf("%10d matches, %10d misses. %10d sum,%.2f%%\n", a, b, a+b, float64(a)*100/float64(a+b))
	}

	for digitCount, digitTable := range results {
		for machineCount, machineData := range digitTable {
			machineNoMatch[machineCount] += machineData[0]
			digitNoMatch[digitCount] += machineData[0]

			machineMatch[machineCount] += machineData[1]
			digitMatch[digitCount] += machineData[1]
			fmt.Printf("%v unique digits, %v machine cards: ", digitCount+leastPossibleDigits, machineCount)
			printPercentage(machineData[1], machineData[0])
		}
		fmt.Println()
	}
	fmt.Println()

	for digitCount := 0; digitCount < possibleDigits; digitCount++ {
		fmt.Printf("%v unique digits: ", digitCount+leastPossibleDigits)
		printPercentage(digitMatch[digitCount], digitNoMatch[digitCount])
	}
	fmt.Println()

	for machineCount := 0; machineCount < ranks.TotalMachines+1; machineCount++ {
		totalMatch += machineMatch[machineCount]
		totalNoMatch += machineNoMatch[machineCount]

		fmt.Printf("%v machine cards: ", machineCount)
		printPercentage(machineMatch[machineCount], machineNoMatch[machineCount])
	}
	fmt.Println()

	fmt.Print("In total: ")
	printPercentage(totalMatch, totalNoMatch)
}

func computeMatches(mysteryCards ranks.MysteryCardHand, yourPlacements ranks.YourRankHand) (int, bool) {
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
	return machines, matches > 5
}
