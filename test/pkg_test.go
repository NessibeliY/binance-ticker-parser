package test

import (
	"testing"

	"github.com/NessibeliY/binance-ticker-parser/pkg"
)

func TestDivideSlice(t *testing.T) {
	tests := []struct {
		name string
		arr  []string
		n    int
		want [][]string
	}{
		{
			name: "test1",
			arr:  []string{"apple", "banana", "pear"},
			n:    2,
			want: [][]string{
				{"apple", "pear"},
				{"banana"},
			},
		},
		{
			name: "test2",
			arr:  []string{"apple", "banana", "pear"},
			n:    3,
			want: [][]string{
				{"apple"},
				{"banana"},
				{"pear"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pkg.DivideSlice(tt.arr, tt.n)

			if !pkg.CompareSlicesOfSlices(got, tt.want) {
				t.Errorf("got %v;\nwant %v", got, tt.want)
			}
		})
	}
}

func TestCompareSlicesOfSlices(t *testing.T) {
	tests := []struct {
		name string
		arr1 [][]string
		arr2 [][]string
		want bool
	}{
		{
			name: "test1",
			arr1: [][]string{{"apple", "banana"}, {"pear"}},
			arr2: [][]string{{"apple", "banana"}, {"pear"}},
			want: true,
		},
		{
			name: "test2",
			arr1: [][]string{{"apple", "banana"}, {"pear"}},
			arr2: [][]string{{"apple", "banana"}},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pkg.CompareSlicesOfSlices(tt.arr1, tt.arr2)

			if got != tt.want {
				t.Errorf("got %t;\nwant %t", got, tt.want)
			}
		})
	}
}
