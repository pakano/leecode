package findmediansortedarrays

import (
	"fmt"
	"testing"
)

func TestFindMedianSortedArrays(t *testing.T) {
	{
		nums1 := []int{1, 3, 4, 7}
		nums2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		fmt.Println(findmediansortedarrays(nums1, nums2))
	}

	{
		nums1 := []int{1, 3}
		nums2 := []int{2}
		fmt.Println(findmediansortedarrays(nums1, nums2))
	}

	{
		nums1 := []int{1, 2}
		nums2 := []int{3, 4}
		fmt.Println(findmediansortedarrays(nums1, nums2))
	}
}
