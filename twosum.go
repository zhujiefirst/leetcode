package main

import "strconv"

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i] + nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{0, 0}
}

func twoSumTwoHash(nums []int, target int) []int {
	m := make(map[string]string) // 这里用k-v都是string的map的原因：简单使用int-int型map时，有可能target-v的过程中出现负数

	for k, v := range nums {
		m[strconv.Itoa(v)] = strconv.Itoa(k)
	}

	for k, v := range nums {
		if vv, ok := m[strconv.Itoa(target-v)]; ok {
			kk, _ := strconv.Atoi(vv)
			if k != kk {
				return []int{k,  kk}
			}
		}
	}

	return []int{0, 0}
}

func twoSumOneHash(nums []int, target int) []int {
	m := make(map[string]string) // 这里用k-v都是string的map的原因：简单使用int-int型map时，有可能target-v的过程中出现负数

	for k, v := range nums {
		// 这段逻辑在insert map之前原因：{3，3} 6这种case时候，若先insert map，依然会只有3一个k-v对。
		if vv, ok := m[strconv.Itoa(target-v)]; ok {
			kk, _ := strconv.Atoi(vv)
			if k != kk {
				return []int{k,  kk}
			}
		}

		m[strconv.Itoa(v)] = strconv.Itoa(k)
	}

	return []int{0, 0}
}