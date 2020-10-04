package slices

import "sort"

func GetWithout(slice []string, without string) []string {
	cc := make([]string, 0, len(slice)-1)
	for _, s := range slice {
		if s != without {
			cc = append(cc, s)
		}
	}
	return cc
}

func Index(slice []string, needle string) (index int) {
	var s string
	for index, s = range slice {
		if s == needle {
			return index
		}
	}
	return -1
}

func Contains(slice []string, needle string) bool {
	return Index(slice, needle) != -1
}

func AppendIfNotFound(slice []string, s string) (newSlice []string) {
	if Index(slice, s) != -1 {
		return append(slice, s)
	}
	return slice
}

func MergeStringSlices(s1, s2 []string, skip string) []string {
	unique := make(map[string]struct{})
	for _, v := range s1 {
		if len(v) < 1 {
			continue
		}
		if v == skip {
			continue
		}
		unique[v] = struct{}{}
	}
	for _, v := range s2 {
		if len(v) < 1 {
			continue
		}
		if v == skip {
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
