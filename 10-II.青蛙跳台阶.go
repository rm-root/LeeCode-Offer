package main

func numWays(n int) int {
	const mod = 1e9 + 7
	if n <= 1 {
		return 1
	}
	p, q, r := 1, 1, 1
	for i := 2; i <= n; i++ {
		p = q
		q = r
		r = (q + p) % mod
	}

	return r
}
