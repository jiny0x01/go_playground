package lotto_util

import "testing"

func TestContain(t *testing.T) {
	result := Contain([]int{1, 2, 3, 4, 5}, 5)
	expected := true
	if result != expected {
		t.Fail()
	}
	result2 := Contain([]float64{1.0, 2.2, 3.3}, 3.14)
	expected2 := false
	if result2 != expected2 {
		t.Fail()
	}
	result3 := Contain([]string{"alice", "bob", "jiny"}, "jiny")
	expected3 := true
	if result3 != expected3 {
		t.Fail()
	}
}
