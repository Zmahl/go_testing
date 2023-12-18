package arrays

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("Slice using size 3", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	t.Run("Using test all on 2 arrays", func(t *testing.T) {
		numbers1 := []int{1, 2}
		numbers2 := []int{0, 9}

		got := SumAll(numbers1, numbers2)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v wanted %v", got, want)
		}
	}
	t.Run("test some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("test nil slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})
}
