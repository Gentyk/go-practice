package leetcode



func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	if n < 4 {
		return n - 1
	}

	array := make([]bool, n - 2)
	result := 0

	for i:=2; i<n; i++ {
		ind := i-2
		if array[ind] == false {
			result += 1
			for j:= 2; j < n/i; j++ {
				new_ind := j*i - 2
				if array[new_ind] == false {
					array[new_ind] = true
				}
			}
		}

	}
    return result
}