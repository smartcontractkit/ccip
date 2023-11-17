package ccipcalc

// TODO: test
func FlattenUniqueSlice[T comparable](slices ...[]T) []T {
	seen := make(map[T]struct{})
	flattened := make([]T, 0)

	for _, sl := range slices {
		for _, el := range sl {
			if _, exists := seen[el]; exists {
				continue
			}
			flattened = append(flattened, el)
			seen[el] = struct{}{}
		}
	}

	return flattened
}
