package main

func TwoSum(nums []int, target int) []int {
	p := []int{0, 0}
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i := 0; i < len(nums); i++ {
		c := target - nums[i]
		if v, found := m[c]; found && i != v {
			p[0] = i
			p[1] = v
			break
		}
	}
	return p
}

func TwoSumOptimized(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		if v, found := m[target-v]; found && i != v {
			return []int{v, i}
		}
		m[v] = i
	}
	return []int{}
}

func TwoSumBruteForce(nums []int, target int) []int {
	p := []int{0, 0}
	for i, v1 := range nums {
		for j, v2 := range nums {
			if i != j {
				if v1+v2 == target {
					p[0] = j
					p[1] = i
				}
			}
		}
	}
	return p
}
