package main

func main() {

	n := int64(10_000)
	P := map[int64]bool{}
	for i := int64(1); i <= n; i++ {
		P[i*(3*i-1)/2] = true
	}

	for p1, _ := range P {
		for p2, _ := range P {
			if P[p1+p2] && P[p2-p1] {
				println(p2 - p1)
			}
		}
	}
}
