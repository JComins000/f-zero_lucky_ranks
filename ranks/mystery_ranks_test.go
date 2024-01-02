package ranks

import (
	"testing"
	"reflect"
)

func TestMatch(test *testing.T) {
	if !(PlacementCard{0}.Match(0)) {
		test.Errorf("expected placement and rank of the same value to match")
	}
	if !(PlacementCard{23}.Match(23)) {
		test.Errorf("expected placement and rank of the same value to match")
	}
	if (PlacementCard{1}.Match(2)) {
		test.Errorf("expected placement and rank of the differ value not to match")
	}
	
	if !(MachineCard{1}.Match(2)) {
		test.Errorf("expected machine match to always return true")
	}
	if !(MachineCard{1}.Match(1)) {
		test.Errorf("expected machine match to always return true")
	}
	
	// remember, wildcards range from 1-9 and placements are off by one
	if !(WildCard{1}.Match(0)) {
		test.Errorf("expect wilcard 1 to match first place")
	}
	if !(WildCard{1}.Match(9)) {
		test.Errorf("expect wilcard 1 to match tenth place")
	}
	if !(WildCard{9}.Match(18)) {
		test.Errorf("expect wilcard 9 to match 19th place")
	}
	if (WildCard{2}.Match(18)) {
		test.Errorf("expect wilcard 2 not to match 19th place")
	}
	
}

func TestMysteryCardList(test *testing.T) {
	if len(allMysteryCards) != totalPlacements + totalMachines + totalWildDigits {
		test.Errorf("expected the list to have all cards for each type")
	}
}

func TestIndexToMysteryRank(test *testing.T) {
	mysteryRankDst := make([]MysteryCard, TotalMysteryCards)
	IndexToMysteryRank(mysteryRankDst, 0)
	
	if !reflect.DeepEqual(mysteryRankDst, allMysteryCards[:TotalMysteryCards]) {
		test.Errorf("expecting first combo to be the beginning of the list, got %v, %v", mysteryRankDst, allMysteryCards[:TotalMysteryCards])
	}

	IndexToMysteryRank(mysteryRankDst, TotalMysteryRankCombinations - 1)
	
	if !reflect.DeepEqual(mysteryRankDst, allMysteryCards[len(allMysteryCards) - TotalMysteryCards:]) {
		test.Errorf("expecting last combo to be the end of the list, got %v, %v", mysteryRankDst, allMysteryCards[:TotalMysteryCards])
	}
}

func TestIndexToMysteryRankBound(test *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            test.Errorf("IndexToMysteryRank did not panic when %v should be too large for it", TotalMysteryRankCombinations)
        }
    }()
	mysteryRankDst := make([]MysteryCard, TotalMysteryCards)
	IndexToMysteryRank(mysteryRankDst, TotalMysteryRankCombinations)
}