package main

import "testing"

func TestMain(t *testing.T) {
	expected := "ab"
	result := Solution("acb")
	if result != expected {
		t.Errorf("Incorrect, Result is %s but expected is: %s", result, expected)
	}

	expected = "ho"
	result = Solution("hot")
	if result != expected {
		t.Errorf("Incorrect, Result is %s but expected is: %s", result, expected)
	}

	expected = "cdility"
	result = Solution("codility")
	if result != expected {
		t.Errorf("Incorrect, Result is %s but expected is: %s", result, expected)
	}

	expected = "aaa"
	result = Solution("aaaa")
	if result != expected {
		t.Errorf("Incorrect, Result is %s but expected is: %s", result, expected)
	}
}
