package search

import (
	"github.com/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestMaxFloat64s(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []float64{1.0, 2, 3.4, 4.78, 5.7}}, 4},
		{"2", args{arr: []float64{5.1, 4.5, 3.4, 2.2, 1.3}}, 0},
		{"3", args{arr: []float64{1.4, 5.45, 4.6, 3.9, 2.0}}, 1},
		{"4", args{arr: []float64{1.9, 2.3, 4.4, 4.4, 3.1}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.MaxFloat64s(tt.args.arr); got != tt.want {
				t.Errorf("MaxFloat64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxInts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []int{1, 2, 3, 4, 5}}, 4},
		{"2", args{arr: []int{5, 4, 3, 2, 1}}, 0},
		{"3", args{arr: []int{1, 5, 4, 3, 2}}, 1},
		{"4", args{arr: []int{1, 2, 4, 4, 3}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.MaxInts(tt.args.arr); got != tt.want {
				t.Errorf("MaxInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMaxStrings(t *testing.T) {
	type args struct {
		arr []string
		x   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []string{"a", "b", "c", "d", "e"}}, 4},
		{"2", args{arr: []string{"elephant", "dog", "cat", "ball", "apple"}}, 0},
		{"3", args{arr: []string{"a", "e", "t", "b", "c"}}, 2},
		{"4", args{arr: []string{"a", "b", "d", "d", "c"}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.MaxStrings(tt.args.arr); got != tt.want {
				t.Errorf("MaxStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
