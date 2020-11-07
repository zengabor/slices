package slices

import "sort"

// WithoutThis returns a new slice without the specified item.
// If the specified item is not found, all items get included in the new slice.
func WithoutThis(slice []string, withoutThis string) []string {
	if len(slice) < 1 {
		return slice
	}
	cc := make([]string, 0, len(slice)-1)
	for _, s := range slice {
		if s != withoutThis {
			cc = append(cc, s)
		}
	}
	return cc
}

// Index returns the index of the specified item, or -1 if it's not found.
func Index(slice []string, needle string) (index int) {
	var s string
	for index, s = range slice {
		if s == needle {
			return index
		}
	}
	return -1
}

// Contains reports whether the specified item is found in the slice.
func Contains(slice []string, needle string) bool {
	return Index(slice, needle) != -1
}

// AppendIfNotFound returns a slice with the item appended if it's not yet in the slice.
func AppendIfNotFound(slice []string, s string) []string {
	if !Contains(slice, s) {
		return append(slice, s)
	}
	return slice
}

// Merge merges two slices into a single slice, skipping withoutThis.
// Empty values are also skipped.
func Merge(s1, s2 []string, withoutThis string) []string {
	unique := make(map[string]struct{})
	for _, v := range s1 {
		if len(v) == 0 {
			continue
		}
		if v == withoutThis {
			continue
		}
		unique[v] = struct{}{}
	}
	for _, v := range s2 {
		if len(v) == 0 {
			continue
		}
		if v == withoutThis {
			continue
		}
		unique[v] = struct{}{}
	}
	final := make([]string, len(unique))
	i := 0
	for k := range unique {
		final[i] = k
		i++
	}
	sort.Slice(final, func(one, two int) bool { return final[one] < final[two] })
	return final
}
