package main

func diff(old []string, new []string) (add []string, remove []string) {
	left := make([]interface{}, len(old))
	for i, v := range old {
		left[i] = v
	}

	right := make([]interface{}, len(new))
	for i, v := range new {
		right[i] = v
	}

	l := lcs.New(left, right)

	leftidx := 0
	rightidx := 0

	for _, pair := range l.IndexPairs() {
		for leftidx < len(left) && leftidx <= pair.Left {
			if leftidx < pair.Left {
				remove = append(remove, old[leftidx])
			}
			leftidx++
		}

		for rightidx < len(right) && rightidx <= pair.Right {
			if rightidx <= pair.Right {
				add = append(add, new[rightidx])
			}
			rightidx++
		}
	}

	for leftidx < len(left) {
		remove = append(remove, old[leftidx])
		leftidx++
	}
	for rightidx < len(right) {
		add = append(add, new[rightidx])
		rightidx++
	}
	return add, remove
}
