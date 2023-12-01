package matching

import (
	"github.com/elliotchance/orderedmap/v2"
)

var (
	Boys  *orderedmap.OrderedMap[int, []Single]
	Girls *orderedmap.OrderedMap[int, []Single]
)

func Initialize() {
	Boys = orderedmap.NewOrderedMap[int, []Single]()
	Girls = orderedmap.NewOrderedMap[int, []Single]()
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
	}
	if single.IsGirl() {
		list := Girls.GetOrDefault(single.Height, []Single{})
		list = append(list, single)
		Girls.Set(single.Height, list)
	}
}

func match(single Single) []Single {
	if single.IsBoy() && Girls.Len() > 0 {
		return Girls.Front().Value
	}
	if single.IsGirl() && Boys.Len() > 0 {
		return Boys.Front().Value
	}
	return []Single{}
}
