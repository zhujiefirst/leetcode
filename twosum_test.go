package main

import "testing"

func TestTwoSum(t *testing.T)  {
	nums := []int{10, 2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	if len(result) != 2 || result[0] != 1 || result[1] != 2 {
		t.Error("FAILED")
	} else {
		t.Log("SUCC")
	}
}

func TestTwoSumTwoHash(t *testing.T)  {
	nums := []int{3, 2, 4}
	target := 6
	result := twoSumTwoHash(nums, target)
	if len(result) != 2 || result[0] != 1 || result[1] != 2 {
		t.Errorf("FAILED: %v", result)
	} else {
		t.Log("SUCC")
	}
}

func TestTwoSumOneHash(t *testing.T)  {
	nums := []int{3, 3}
	target := 6
	result := twoSumOneHash(nums, target)
	if len(result) != 2 || result[0] != 1 || result[1] != 0 {
		t.Errorf("FAILED: %v", result)
	} else {
		t.Log("SUCC")
	}
}
