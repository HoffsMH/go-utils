package kata

import (
	"sort"
	"testing"
)
func arraysEqual(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }

    for i := range a {
        if a[i] != b[i] {
            return false
        }
    }

    return true
}

func TestBasicQuickSort(t *testing.T) {
  input := []int{45, 38, 26, 12, 89, 53, 74, 91, 21, 34}
  result := make([]int, len(input))
  copy(result, input)
  expected := []int{45, 38, 26, 12, 89, 53, 74, 91, 21, 34}
  sort.Ints(expected)

	QuickSort(result)
	if !arraysEqual(result, expected) {
		t.Errorf("QuickSort(%d) = %d; expected %d", input, result, expected)
	}
}
