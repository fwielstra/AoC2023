package utils

// Contains is a generic function to determine if `e` is contained in slice `s`.
// https://stackoverflow.com/questions/10485743/contains-method-for-a-slice
// Note: beware of N x M comparisons for longer data sets.
func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}
