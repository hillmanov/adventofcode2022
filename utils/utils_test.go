package utils

import (
	"testing"
)

func TestMin(t *testing.T) {
	if Min(1, 2) != 1 {
		t.Error("Min(1, 2) != 1")
	}
	if Min(2, 1) != 1 {
		t.Error("Min(2, 1) != 1")
	}

	if Min(1.0, 2.0) != 1.0 {
		t.Error("Min(1.0, 2.0) != 1.0")
	}
	if Min(2.0, 1.0) != 1.0 {
		t.Error("Min(2.0, 1.0) != 1.0")
	}
}

func TestMax(t *testing.T) {
	if Max(1, 2) != 2 {
		t.Error("Max(1, 2) != 2")
	}
	if Max(2, 1) != 2 {
		t.Error("Max(2, 1) != 2")
	}

	if Max(1.0, 2.0) != 2.0 {
		t.Error("Max(1.0, 2.0) != 2.0")
	}
	if Max(2.0, 1.0) != 2.0 {
		t.Error("Min(2.0, 1.0) != 2.0")
	}
}

func TestMaxOf(t *testing.T) {
	if MaxOf([]int{1, 2, 3}) != 3 {
		t.Error("MaxOf([]int{1, 2, 3}) != 3")
	}

	if MaxOf([]int{3, 2, 1}) != 3 {
		t.Error("MaxOf([]int{3, 2, 1}) != 3")
	}
}

func TestMinOf(t *testing.T) {
	if MinOf([]int{1, 2, 3}) != 1 {
		t.Error("MinOf([]int{1, 2, 3}) != 1")
	}

	if MinOf([]int{3, 2, 1}) != 1 {
		t.Error("MinOf([]int{3, 2, 1}) != 1")
	}
}

func TestMinMax(t *testing.T) {
	min, max := MinMax([]int{1, 2, 3})
	if min != 1 || max != 3 {
		t.Error("MinMax([]int{1, 2, 3}) != (1, 3)")
	}

	min, max = MinMax([]int{3, 2, 1})
	if min != 1 || max != 3 {
		t.Error("MinMax([]int{3, 2, 1}) != (1, 3)")
	}
}

func TestSumOf(t *testing.T) {
	if SumOf([]int{1, 2, 3}) != 6 {
		t.Error("SumOf([]int{1, 2, 3}) != 6")
	}
}

func TestAbs(t *testing.T) {
	if Abs(-1) != 1 {
		t.Error("Abs(-1) != 1")
	}
	if Abs(1) != 1 {
		t.Error("Abs(1) != 1")
	}

	if Abs(-1.5) != 1.5 {
		t.Error("Abs(-1.5) != 1.5")
	}
	if Abs(1.75) != 1.75 {
		t.Error("Abs(1.75) != 1.75")
	}
}

func TestUniqueOf(t *testing.T) {
	if len(UniqueOf([]int{1, 2, 3, 2, 1})) != 3 {
		t.Error("len(UniqueOf([]int{1, 2, 3, 2, 1})) != 3")
	}

	if len(UniqueOf([]float64{1.5, 2.5, 3.5, 2.5, 1.5})) != 3 {
		t.Error("len(UniqueOf([]int{1.5, 2.5, 3.5, 2.5, 1.5})) != 3")
	}
}

func TestIndexOf(t *testing.T) {
	if IndexOf([]int{1, 2, 3}, 2) != 1 {
		t.Error("IndexOf([]int{1, 2, 3}, 2) != 1")
	}

	if IndexOf([]int{1, 2, 3}, 4) != -1 {
		t.Error("IndexOf([]int{1, 2, 3}, 4) != -1")
	}
}

func TestReverse(t *testing.T) {
	if Reverse([]int{1, 2, 3, 4, 5, 6})[0] != 6 {
		t.Error("Reverse([]int{1, 2, 3, 4, 5, 6})[0] != 6")
	}
}

func TestPop(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5, 6}
	value, collection := Pop(collection)
	if value != 6 && collection[0] != 2 {
		t.Error("Pop([]int{1, 2, 3, 4, 5, 6}) != (6, []int{2, 3, 4, 5, 6})")
	}
}

func TestRemoveAtIndex(t *testing.T) {
	collection := []int{1, 2, 3, 4, 5, 6}
	collection = RemoveAtIndex(collection, 1)
	if len(collection) != 5 {
		t.Error("len(RemoveAtIndex([]int{1, 2, 3, 4, 5, 6}, 2)) != 5")
	}
}

func TestPermutations(t *testing.T) {
	if len(Permutations([]int{1, 2, 3})) != 6 {
		t.Error("len(Permutations([]int{1, 2, 3})) != 6")
	}
}
