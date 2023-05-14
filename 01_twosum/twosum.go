package twosum

func TwoSum(nums []int, target int) []int {
	var m = make(map[int]int)

	for idx, num := range nums {
		complement := target - num
		if idx0, ok := m[complement]; ok {
			return []int{idx0, idx}
		}
		m[num] = idx
	}
	return nil
}

/*
### 题目解析

使用查找表来解决该问题。

设置一个 map 容器 record 用来记录元素的值与索引，然后遍历数组 nums。

* 每次遍历时使用临时变量 complement 用来保存目标值与当前值的差值
* 在此次遍历中查找 record ，查看是否有与 complement 一致的值，如果查找成功则返回查找值的索引值与当前变量的值 i
* 如果未找到，则在 record 保存该元素与索引值 i
*/
