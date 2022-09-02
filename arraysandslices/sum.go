package arraysandslices

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAllTails(slcOfSlices ...[]int) []int {
	var sums []int
	for _, nums := range slcOfSlices {
		if len(nums) > 0 {
			sums = append(sums, Sum(nums[1:]))
		} else {
			sums = append(sums, 0)
		}
	}
	return sums
}
