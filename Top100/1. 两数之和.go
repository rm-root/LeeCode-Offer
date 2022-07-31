package main

//题目：给定一个整数数组 nums和一个整数目标值 target，请你在该数组中找出 和 为下目标值 target 的那两个整数，并返回它们的数组下标。
//你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
//你可以按任意顺序返回答案。
//2 <= nums.length <= 10^4
//-10^9 <= nums[i] <= 10^9
//-10^9 <= target <= 10^9
//只会存在一个有效答案

//方法一暴力枚举

//func twoSum(nums []int, target int) []int {
//
//	for i, x := range nums {
//		for j := i + 1; j < len(nums); j++ {
//			if x+nums[j] == target {
//				return []int{i, j}
//			}
//		}
//	}
//	return nil
//}
//时间复杂度分析时间复杂度：O(n^2) 其中 n 是数组中的元素数量。最坏情况下数组中任意两个数都要被匹配一次。
//空间复杂度：O(1)。
//

//方法二哈希表
func twoSum(nums []int, target int) []int {
	hashTable := map[int]int{}
	for i, x := range nums {
		//判断target-x是否再map里面，如果在v为map[target-x]的值,ok为true；否则v得到相应类型的零值；ok是false。
		if v, ok := hashTable[target-x]; ok {
			return []int{v, i}
		}
		hashTable[x] = i
	}
	return nil
}

//时间复杂度：O(N)
//空间复杂度：O(N)，其中N是数组中的元素数量。主要为哈希表的开销

//func main() {
//	fmt.Print(twoSum([]int{1, 4, 5}, 6))
//}
