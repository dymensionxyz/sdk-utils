package uslices_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/exp/maps"

	"github.com/dymensionxyz/sdk-utils/utils/uslices"
)

func TestToKeySet(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected map[int]struct{}
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			expected: map[int]struct{}{},
		},
		{
			name:     "Single element slice",
			input:    []int{1},
			expected: map[int]struct{}{1: {}},
		},
		{
			name:     "Multiple elements, no duplicates",
			input:    []int{1, 2, 3},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
		{
			name:     "Multiple elements with duplicates",
			input:    []int{1, 2, 2, 3, 1},
			expected: map[int]struct{}{1: {}, 2: {}, 3: {}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uslices.ToKeySet(tc.input)
			require.True(t, maps.Equal(tc.expected, result))
		})
	}
}

func TestMerge(t *testing.T) {
	testCases := []struct {
		name     string
		input    [][]int
		expected []int
	}{
		{
			name:     "No slices",
			input:    [][]int{},
			expected: []int{},
		},
		{
			name:     "Single empty slice",
			input:    [][]int{{}},
			expected: []int{},
		},
		{
			name:     "Single non-empty slice",
			input:    [][]int{{1, 2, 3}},
			expected: []int{1, 2, 3},
		},
		{
			name:     "Multiple non-empty slices",
			input:    [][]int{{1, 2}, {3, 4}, {5}},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Multiple slices, some empty",
			input:    [][]int{{1, 2}, {}, {3, 4}, {}, {5}},
			expected: []int{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uslices.Merge(tc.input...)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		mapFunc  func(int) int
		expected []int
	}{
		{
			name:     "Empty slice",
			input:    []int{},
			mapFunc:  func(x int) int { return x * 2 },
			expected: []int{},
		},
		{
			name:     "Simple mapping",
			input:    []int{1, 2, 3},
			mapFunc:  func(x int) int { return x * 2 },
			expected: []int{2, 4, 6},
		},
		{
			name:     "Identity mapping",
			input:    []int{1, 2, 3},
			mapFunc:  func(x int) int { return x },
			expected: []int{1, 2, 3},
		},
		{
			name:     "Mapping to constant",
			input:    []int{1, 2, 3},
			mapFunc:  func(int) int { return 42 },
			expected: []int{42, 42, 42},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := uslices.Map(tc.input, tc.mapFunc)
			require.Equal(t, tc.expected, result)
		})
	}
}
