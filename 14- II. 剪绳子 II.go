package main

//const mod = 1000000007

func cuttingRope2(n int) int {

	//1 1
	//2 1
	//3 2
	//4 2 2
	//5 3 2
	//6 3 3
	//7 3 2 2
	//8 3 3 2

	if n <= 2 {
		return 1
	}

	if n == 3 {
		return 2
	}
	part := n / 3
	another := n % 3
	var result int
	switch another {
	case 2:
		result = Pow(3, part) * 2.0 % mod
	case 1:
		result = Pow(3, part-1) * 4.0 % mod
	default:
		result = Pow(3, part) % mod
	}
	return result
}

// 搞个快速幂
func Pow(x int, y int) int {
	var result = 1
	for y > 0 {
		if y%2 == 1 {
			result = result * x % mod
		}
		y = y / 2
		x = x * x % mod
	}
	return result
}

//
//func main() {
//	fmt.Print(cuttingRope2(999))
//}
