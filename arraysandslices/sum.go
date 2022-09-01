package arraysandslices

func Sum(nums []int) int {
	var sum int
	for _, num := range nums {
		sum += num
	}
	return sum
}

func SumAll(slcOfSlices ...[]int) []int {
	var sums []int
	for _, nums := range slcOfSlices {
		sums = append(sums, Sum(nums))
	}
	return sums
}
