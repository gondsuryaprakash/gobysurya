package leetcode

func ProductExceptSelf(nums []int) []int {

	product := []int{}

	prefixProdux := 1
	for i := 0; i < len(nums); i++ {
		product = append(product, prefixProdux)
		product[i] = prefixProdux
		prefixProdux *= nums[i]
	}

	suffixProduct := 1

	for i := len(nums) - 1; i >= 0; i-- {
		product[i] *= suffixProduct
		suffixProduct *= nums[i]
	}
	return product
}
