package problems

//  -------------   submission start

func twoSum(nums []int, target int) []int {
	index := make(map[int]int, len(nums)) // Number[its position]
	for i, num := range nums {
		if ii, ok := index[target-num]; ok {
			return []int{ii, i}
		}
		index[num] = i

	}
	return nums
}

//  -------------   submission end
