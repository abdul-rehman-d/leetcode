package product_of_array_except_self

func ProductExceptSelf(nums []int) []int {
	leftProducts, rightProducts := make([]int, len(nums)), make([]int, len(nums))
	prefix, suffix := 1, 1
	for i, j := 0, len(nums)-1; i < len(nums); i, j = i+1, j-1 {
		leftProducts[i] = prefix
		prefix *= nums[i]
		rightProducts[j] = suffix
		suffix *= nums[j]
	}

	for i := range nums {
		nums[i] = leftProducts[i] * rightProducts[i]
	}

	return nums
}
