package leetcode
/*
1 生成map 保存尚未找到答案时 数组元素值以及对应在数组中的下标
2 遍历数组
	2.1 当前遍历元素对应配对元素是否在map中 在就返回即可
    2.2 不在将当前元素和下标放入map
3 时间复杂度O(n)
*/
func twoSum(nums []int, target int) []int {
	// 申请足够的内存 避免重复申请内存 降低性能
	m := make(map[int]int,len(nums))
	// idx:数组下标		v:数组元素内容
	for idx, v := range nums {
		// idxOlder为target-v在map中的记录的原数组的下标位置
		if idxOlder, ok := m[target-v]; ok {
			return []int{idxOlder,idx}
		}
		m[v] = idx
	}
	return nil
}
