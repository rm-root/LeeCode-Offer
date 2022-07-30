package main

//编写一个函数，输入是一个无符号整数（以二进制串的形式），
//返回其二进制表达式中数字位数为 '1' 的个数（也被称为 汉明重量).）。
//官方方法：n&(n-1),这个运算结果恰为把n的二进制位中的最低位的1变为0，
//可以利用这个性质，不断进行n&(n-1)运算把n二进制中的1去掉，运算次数就等于n的二进制中1的个数。

func hammingWeight(num uint32) (ones int) {
	for ; num > 0; num &= num - 1 {
		ones++
	}
	return
}

//func hammingWeight(num uint32) (ones int) {
//	for i := 0; i < 32; i++ {
//		if 1<<i & num > 0 {
//			ones++
//		}
//	}
//	return
//}

//func main() {
//	hammingWeight(7)
//}
