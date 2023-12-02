package matching

import (
	"github.com/elliotchance/orderedmap/v2"
	"math"
	"sort"
)

var (
	Boys  *orderedmap.OrderedMap[int, []Single]
	Girls *orderedmap.OrderedMap[int, []Single]

	SortedBoys  []int
	SortedGirls []int
)

func Initialize() {
	Boys = orderedmap.NewOrderedMap[int, []Single]()
	Girls = orderedmap.NewOrderedMap[int, []Single]()

	SortedBoys = []int{}
	SortedGirls = []int{}
}

func GetPossibleMatches(count int) []Single {
	if Boys.Len() == 0 && Girls.Len() == 0 {
		return []Single{}
	}

	if Boys.Len() == 0 {
		return searchPossibleMatches(count, SortedGirls, Girls)
	}
	if Girls.Len() == 0 {
		return searchPossibleMatches(count, SortedBoys, Boys)
	}

	half := count / 2
	take := half

	if Boys.Len() < half {
		take = Boys.Len()
		possibleGirls := searchPossibleMatches(count-take, SortedGirls, Girls)
		possibleBoys := searchPossibleMatches(take, SortedBoys, Boys)
		return append(possibleBoys, possibleGirls...)
	}

	possibleGirls := searchPossibleMatches(take, SortedGirls, Girls)
	possibleBoys := searchPossibleMatches(count-take, SortedBoys, Boys)
	return append(possibleGirls, possibleBoys...)
}

func searchPossibleMatches(
	count int,
	sortedHeights []int,
	lookupMap *orderedmap.OrderedMap[int, []Single],
) []Single {
	var possibleMatches []Single

	for _, height := range sortedHeights {
		singles, _ := lookupMap.Get(height)
		take := math.Min(float64(len(singles)), float64(count))
		possibleMatches = append(possibleMatches, singles[0:int(take)]...)
		count -= int(take)
		if count == 0 {
			break
		}
	}
	return possibleMatches
}

func AddAndMatch(single Single) []Single {
	add(single)
	return match(single)
}

func add(single Single) {
	if single.IsBoy() {
		list := Boys.GetOrDefault(single.Height, []Single{})
		list = append(list, single)
		Boys.Set(single.Height, list)

		if len(list) == 1 {
			SortedBoys = Boys.Keys()
			for i, j := 0, len(SortedBoys)-1; i < j; i, j = i+1, j-1 {
				SortedBoys[i], SortedBoys[j] = SortedBoys[j], SortedBoys[i]
			}
		}
	}
	if single.IsGirl() {
		list := Girls.GetOrDefault(single.Height, []Single{})
		list = append(list, single)
		Girls.Set(single.Height, list)

		if len(list) == 1 {
			SortedGirls = Girls.Keys()
			sort.Ints(SortedGirls)
		}
	}
}

func match(single Single) []Single {
	if single.IsBoy() && Girls.Len() > 0 {
		shortest := Girls.Front()
		if single.Height < shortest.Key {
			return []Single{}
		}
		return shortest.Value
	}

	if single.IsGirl() && Boys.Len() > 0 {
		highest := Boys.Front()
		if single.Height > highest.Key {
			return []Single{}
		}
		return highest.Value
	}

	return []Single{}
}
