package ranks

import (
	"gonum.org/v1/gonum/stat/combin"
)

type MysteryCard interface {
	Match(int) bool
}

type PlacementCard struct {
	value int
}

type MachineCard struct {
	value int
}

type WildCard struct {
	value int
}

func (placement PlacementCard) Match(yourRank int) bool {
	return placement.value == yourRank
}

func (machine MachineCard) Match(yourRank int) bool {
	// in the game, you can select any machine and the
	// cards are revealed before you finish your placements.
	// it's hard not to match these cards if you're trying for matches
	return true
}

func (wild WildCard) Match(yourRank int) bool {
	// rank values range from 0 to 98, so we add 1
	// wild values already range from 1 to 9
	realRankValue := yourRank + 1
	return realRankValue%10 == wild.value || realRankValue/10 == wild.value
}

func getAllMysteryRanks() []MysteryCard {
	var allCards []MysteryCard

	for i := 0; i < totalPlacements; i++ {
		allCards = append(allCards, PlacementCard{i})
	}

	for i := 0; i < TotalMachines; i++ {
		allCards = append(allCards, MachineCard{i})
	}

	for i := 0; i < totalWildDigits; i++ {
		// ensure wild values range from 1 to 9 for convenience
		allCards = append(allCards, WildCard{i + 1})
	}

	return allCards
}

func IndexToMysteryRank(dst []MysteryCard, idx int) {
	combo := combin.IndexToCombination(nil, idx, len(allMysteryCards), TotalMysteryCards)
	for i := 0; i < len(dst); i++ {
		dst[i] = allMysteryCards[combo[i]]
	}
}

var allMysteryCards []MysteryCard = getAllMysteryRanks()
var TotalMysteryRankCombinations int = combin.Binomial(len(allMysteryCards), TotalMysteryCards)
