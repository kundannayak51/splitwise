package splitvalidation

import "splitwise/entity"

func GetSplitObjectFromFactory(splitType entity.SPLITTYPE) ExpenseSplit {
	switch splitType {
	case entity.EQUAL:
		equalSplitObject := NewEqualSplit()
		return equalSplitObject
	}

	return nil
}
