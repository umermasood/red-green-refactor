package arraysandslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any size", func(t *testing.T) {
		nums := []int{1, 2, 3}

		got := Sum(nums)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, nums)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("sum some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{0, 1}, []int{9, 9})
		want := []int{5, 1, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{1, 2, 3})
		want := []int{0, 5}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
