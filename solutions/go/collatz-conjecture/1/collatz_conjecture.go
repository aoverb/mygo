package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return -1, errors.New("n must be a positive number")
	}
	cnt := 0
	for n != 1 {
		cnt++
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
	}
	return cnt, nil
}
