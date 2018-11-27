package slices

import "testing"

func TestGetWithout(t *testing.T) {
	t.Fatal("test not implemented")
}

func TestIndex(t *testing.T) {
	t.Fatal("test not implemented")
}

func TestAppendIfNotFound(t *testing.T) {
	t.Fatal("test not implemented")
}

func TestMergeStringSlices(t *testing.T) {
	var tests = []struct {
		slice1 []string
		slice2 []string
		skip   string
		want   []string
	}{
		{[]string{"s2", "s1"}, []string{}, "", []string{"s1", "s2"}},
		{[]string{"s2", "s1"}, []string{"s1"}, "", []string{"s1", "s2"}},
		{[]string{"s2", "s1"}, []string{"s1", "s2"}, "", []string{"s1", "s2"}},
		{[]string{"s2"}, []string{"s1", "s2"}, "", []string{"s1", "s2"}},
		{[]string{"s2"}, []string{"s1"}, "", []string{"s1", "s2"}},
		{[]string{"s2"}, []string{"s1"}, "s1", []string{"s2"}},
		{[]string{"s2"}, []string{"s1"}, "s3", []string{"s1", "s2"}},
		{[]string{}, []string{}, "", []string{}},
	}
	for _, test := range tests {
		if got := MergeStringSlices(test.slice1, test.slice2, test.skip); !isTheSame(got, test.want) {
			t.Errorf("mergeStringSlices(%+v, %+v, %s) = %v", test.slice1, test.slice2, test.skip, got)
		}
	}
}

func isTheSame(slice1, slice2 []string) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i, s := range slice1 {
		if s != slice2[i] {
			return false
		}
	}
	return true
}
