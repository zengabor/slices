package slices

import (
	"fmt"
	"testing"
)

func TestWithoutThis(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	if fmt.Sprintf("%v", WithoutThis(s, "a")) != "[b c d e]" {
		t.Error(WithoutThis(s, "a"))
	}
	if fmt.Sprintf("%v", WithoutThis(s, "c")) != "[a b d e]" {
		t.Error(WithoutThis(s, "c"))
	}
	if fmt.Sprintf("%v", WithoutThis(s, "e")) != "[a b c d]" {
		t.Error(WithoutThis(s, "e"))
	}
	if fmt.Sprintf("%v", WithoutThis(s, "A")) != "[a b c d e]" {
		t.Error(WithoutThis(s, "A"))
	}
	if fmt.Sprintf("%v", WithoutThis(s, "x")) != "[a b c d e]" {
		t.Error(WithoutThis(s, "x"))
	}
	if fmt.Sprintf("%v", WithoutThis([]string{}, "x")) != "[]" {
		t.Error(WithoutThis([]string{}, "x"))
	}
}

func TestIndex(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	if Index(s, "a") != 0 {
		t.Error(Index(s, "a"))
	}
	if Index(s, "c") != 2 {
		t.Error(Index(s, "c"))
	}
	if Index(s, "e") != 4 {
		t.Error(Index(s, "e"))
	}
	if Index(s, "A") != -1 {
		t.Error(Index(s, "A"))
	}
	if Index(s, "x") != -1 {
		t.Error(Index(s, "x"))
	}
	if Index(s, "") != -1 {
		t.Error(Index(s, ""))
	}
}

func TestContains(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	if !Contains(s, "a") {
		t.Error(Contains(s, "a"))
	}
	if !Contains(s, "c") {
		t.Error(Contains(s, "c"))
	}
	if !Contains(s, "e") {
		t.Error(Contains(s, "e"))
	}
	if Contains(s, "A") {
		t.Error(Contains(s, "A"))
	}
	if Contains(s, "x") {
		t.Error(Contains(s, "x"))
	}
	if Contains(s, "") {
		t.Error(Contains(s, ""))
	}
}

func TestAppendIfNotFound(t *testing.T) {
	s := []string{"a", "b", "c", "d", "e"}
	if fmt.Sprintf("%v", AppendIfNotFound(s, "a")) != "[a b c d e]" {
		t.Error(AppendIfNotFound(s, "a"))
	}
	if fmt.Sprintf("%v", AppendIfNotFound(s, "e")) != "[a b c d e]" {
		t.Error(AppendIfNotFound(s, "e"))
	}
	if fmt.Sprintf("%v", AppendIfNotFound(s, "x")) != "[a b c d e x]" {
		t.Error(AppendIfNotFound(s, "x"))
	}
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
		{[]string{"s2", "", "s1"}, []string{"s1", "s2"}, "", []string{"s1", "s2"}},
		{[]string{"s2", "s1"}, []string{"s1", "s2", ""}, "", []string{"s1", "s2"}},
		{[]string{"s2"}, []string{"s1", "s2"}, "", []string{"s1", "s2"}},
		{[]string{"s2"}, []string{"s1"}, "", []string{"s1", "s2"}},
		{[]string{"s1"}, []string{"s2"}, "s1", []string{"s2"}},
		{[]string{"s2"}, []string{"s1"}, "s1", []string{"s2"}},
		{[]string{"s2"}, []string{"s1"}, "s3", []string{"s1", "s2"}},
		{[]string{}, []string{}, "", []string{}},
	}
	for _, test := range tests {
		if got := Merge(test.slice1, test.slice2, test.skip); !isTheSame(got, test.want) {
			t.Errorf("merge(%+v, %+v, %s) = %v", test.slice1, test.slice2, test.skip, got)
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
