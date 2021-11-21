package main

// func TwoSum(nums []int, target int) []int {
//     p := []int{0,0}
//     m := make(map[int]int)
//     for i,v := range nums {
//         m[v] = i
//     }
//     for k,v := range m {
//         if target - v == m[k]
//             p[0] = k
//             p[1] = m[k]
//     }
//     return p
// }

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
