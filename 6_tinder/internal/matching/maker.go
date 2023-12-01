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
