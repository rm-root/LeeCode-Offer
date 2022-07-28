package main

func movingCount(m int, n int, k int) int {
	vis := make([][]bool, m)
	for i := range vis {
		vis[i] = make([]bool, n)
	}
	return dfs(0, 0, m, n, k, vis)
}

func getSum(x int) int {
	sx := 0
	for x != 0 {
		sx += x % 10
		x /= 10
	}
	return sx
}

func dfs(i, j, m, n, k int, vis [][]bool) int {
	if i > m-1 || j > n-1 || getSum(i)+getSum(j) > k || vis[i][j] {
		return 0
	}
	vis[i][j] = true

	return dfs(i+1, j, m, n, k, vis) + dfs(i, j+1, m, n, k, vis) + 1
}
