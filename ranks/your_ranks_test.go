package ranks

import (
	"reflect"
	"testing"
)

func TestIndexToYourRank(test *testing.T) {
	// actual
	yourRankDst := YourRankHand{}
	IndexToYourRank(&yourRankDst, 0)
	
	// expected
	var expectedRanks = make([]int, TotalYourRanks)
	for i := 0; i < TotalYourRanks; i++ {
		expectedRanks[i] = i
	}
	
	// condition
	if !reflect.DeepEqual(yourRankDst[:], expectedRanks) {
		test.Errorf("expected counting numbers as first combo, got %v, %v", yourRankDst, expectedRanks)
	}

	// actual
	IndexToYourRank(&yourRankDst, TotalYourRankCombinations - 1)

	// expected
	for i := 0; i < TotalYourRanks; i++ {
		expectedRanks[TotalYourRanks - i - 1] = totalPlacements - i - 1
	}
	
	// condition
	if !reflect.DeepEqual(yourRankDst[:], expectedRanks) {
		test.Errorf("expected reverse counting numbers as last combo, got %v, %v", yourRankDst, expectedRanks)
	}
}

func TestIndexToYourRankBound(test *testing.T) {
    defer func() {
        if r := recover(); r == nil {
            test.Errorf("IndexToYourRank did not panic when %v should be too large for it", TotalYourRankCombinations)
        }
    }()
	yourRankDst := YourRankHand{}
	IndexToYourRank(&yourRankDst, TotalYourRankCombinations)
}

func TestUniqueDigits(test *testing.T) {
	var testDigits = func(digits []int, expectedValue int) {
		// convert digits to placements since the function uses placements (which are off by one)
		actualPlacements := make([]int, len(digits))

		for index, placement := range digits {
			actualPlacements[index] = placement - 1
		}
		actualValue := UniqueDigits(actualPlacements)
		if UniqueDigits(actualPlacements) != expectedValue {
			test.Errorf("expected placements %v to have %v unique digits, but we got %v", digits, expectedValue, actualValue)
		}
	}
	testDigits([]int {1, 2, 3}, 4)
	testDigits([]int {1}, 2)
	testDigits([]int {1234567890}, 10)
	testDigits([]int {123456789123456789}, 9)
	testDigits([]int {9}, 2)
	testDigits([]int {10, 1}, 2)
	testDigits([]int {40, 11}, 3)
}