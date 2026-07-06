package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("length not equal")
	}
	ret := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			ret++
		}
	}
	return ret, nil
}
