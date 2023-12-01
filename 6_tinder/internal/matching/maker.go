package matching

import (
	"github.com/elliotchance/orderedmap/v2"
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
		value, _ := Girls.Get(SortedGirls[0])
		return value
	}
	if Girls.Len() == 0 {
		return Boys.Back().Value
	}

	return []Single{}
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
