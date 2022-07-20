package lotto_util

import "golang.org/x/exp/constraints"

func Contain[T constraints.Ordered](slices []T, target T) bool {
	for _, element := range slices {
		if element == target {
			return true
		}
	}
	return false
}