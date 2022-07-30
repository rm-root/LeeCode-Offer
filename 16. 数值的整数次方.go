package main

//实现 pow(x, n) ，即计算 x 的 n 次幂函数（即，x^n）。
//不得使用库函数，同时不需要考虑大数问题。
func myPow(x float64, n int) float64 {
	if n < 0 {
		return 1 / quickMul(x, -n)
	}
	return quickMul(x, n)

}
func quickMul(x float64, n int) float64 {
	var res float64 = 1
	for n > 0 {
		if n%2 == 1 {
			res = res * x
		}
		n = n / 2
		x = x * x
	}

	return res
}

//func main() {
//	fmt.Print(myPow(2, -2))
//}
