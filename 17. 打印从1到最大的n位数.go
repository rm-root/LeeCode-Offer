package main

//func printNumbers(n int) []int {
//	maxNumber := int(math.Pow10(n))
//	ans := make([]int, maxNumber, maxNumber)
//	for i := 1; i < maxNumber; i++ {
//		ans[i] = i
//	}
//	return ans[1:]
//}

//考虑大数溢出的情况
//为了避免数字开头出现0，先把首位first固定，first取值范围为1~9
//用digit表示要生成的数字的位数，本题要从1位数一直生成到n位数，对每种数字的位数都生成一下首位，所以有个双重for循环
//生成首位之后进入递归生成剩下的digit - 1位数，从0~9中取值
//递归的中止条件为已经生成了digit位的数字，即index == digit，将此时的数[]byte类型的num转为int加到结果res中
//
//func printNumbers(n int) []int {
//	res := make([]int, int(math.Pow10(n)-1))
//	count := 0
//
//	//dfs函数表示从首位（index）开始生成到digit位数字
//	var dfs func(index int, digit int, num []byte)
//	dfs = func(index int, digit int, num []byte) {
//		if index == digit {
//			tmp, _ := strconv.Atoi(string(num))
//			res[count] = tmp
//			count++
//			return
//		}
//
//		//进入递归生成剩下的digit - 1位数
//		for i := '0'; i <= '9'; i++ {
//			num[index] = byte(i)
//			dfs(index+1, digit, num)
//		}
//	}
//
//	//从1开始生成n位数
//	for digit := 1; digit <= n; digit++ {
//		//生成首位
//		for first := '1'; first <= '9'; first++ {
//			num := make([]byte, digit)
//			num[0] = byte(first)
//			dfs(1, digit, num)
//		}
//	}
//	return res
//}

//func main() {
//	fmt.Print(printNumbers(2))
//}
