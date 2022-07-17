package main

func findRepeatNumber(nums []int) int {
	for i := 0; i < len(nums); i++ {
		for nums[i] != i { //注意这里一定不能将for改为if
			if nums[nums[i]] == nums[i] {
				return nums[i] //返回重复元素
			}
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i] //不等就交换
		}
	}
	return 1
}

//func FindRepeatNumber(nums []int) int {
//	var res int
//	numap := make(map[int]int)
//	for i := range nums {
//		if _, ok := numap[nums[i]]; ok { //固定语法，如果map里有以该元素为key的键值对，则ok为true，即找到重复元素，未找到ok为false，在map中加上该键值对
//			res = nums[i] //ok为true，找到重复元素，直接退出循环
//			break
//		}
//		numap[nums[i]]++
//	}
//	return res
//}
