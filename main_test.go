package main

import (
	"reflect"
	"testing"
)

func TestPowerSet(t *testing.T) {
	got := PowerSet([]int{1, 2, 3})
	want := [][]int{
		{}, {1}, {2}, {1, 2}, {3}, {1, 3}, {2, 3}, {1, 2, 3},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("PowerSet = %v, want %v", got, want)
	}

	if got := PowerSet(nil); !reflect.DeepEqual(got, [][]int{{}}) {
		t.Errorf("PowerSet(nil) = %v, want [[]]", got)
	}
	if n := len(PowerSet([]int{1, 2, 3, 4, 5})); n != 32 {
		t.Errorf("size = %d, want 32", n)
	}
}
