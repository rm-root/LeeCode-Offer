package main

//给你一根长度为 n 的绳子，请把绳子剪成整数长度的 m 段（m、n都是整数，n>1并且m>1），
//每段绳子的长度记为 k[0],k[1]...k[m-1] 。
//请问 k[0]*k[1]*...*k[m-1] 可能的最大乘积是多少？
//例如，当绳子的长度是8时，我们把它剪成长度分别为2、3、3的三段，
//此时得到的最大乘积是18。

//方法一
//func cuttingRope(n int) int {
//	// 11  3 3 3 2
//	// 10  3 3 3 1  => 3 3 2 2
//	// 9   3 3 3
//	// 8   3 3 2
//	// 7   3 3 1  => 3 2 2
//	// 6   3 3
//	// 5   3 2
//	// 4   3 1  => 2 2
//
//	// 特殊情况
//	// 3   2 1
//	// 2   1 1
//	// 1   1
//
//	if n <= 2 {
//		return 1
//	}
//	if n == 3 {
//		return 2
//	}
//	// 能分成几份3
//	parts := n / 3
//	another := n % 3
//	var result float64
//
//	switch another {
//	case 2:
//		result = math.Pow(3, float64(parts))
//		result *= 2
//	case 1:
//		result = math.Pow(3, float64(parts-1))
//		result *= 4
//	default:
//		result = math.Pow(3, float64(parts))
//	}
//	return int(result)
//
//}

//方法二
// 动态规划方法
func cuttingRope(n int) int {
	dp := make(map[int]int)
	dp[1] = 1
	for i := 2; i < n+1; i++ {
		// 切割点为j，j属于[1,i]，这里参照了他人的优化技巧
		for j := 1; j < (i/2 + 1); j++ {
			dp[i] = max(dp[i], max(j, dp[j])*max(i-j, dp[i-j]))
		}
	}
	return dp[n]
}

func max(x int, y int) int {
	if x > y {
		return x
	}
	return y
}

//func main() {
//	fmt.Print(cuttingRope(5))
//}
