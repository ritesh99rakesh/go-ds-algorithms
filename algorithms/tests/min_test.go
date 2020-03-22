package search

import (
	"github.com/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestMinFloat64s(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []float64{1.0, 2, 3.4, 4.78, 5.7}}, 0},
		{"2", args{arr: []float64{5.1, 4.5, 3.4, 2.2, 1.3}}, 4},
		{"3", args{arr: []float64{1.4, 5.45, 4.6, 3.9, 2.0}}, 0},
		{"4", args{arr: []float64{1.9, 1.9, 4.4, 4.4, 3.1}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.MinFloat64s(tt.args.arr); got != tt.want {
				t.Errorf("MinFloat64s() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinInts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []int{1, 2, 3, 4, 5}}, 0},
		{"2", args{arr: []int{5, 4, 3, 2, 1}}, 4},
		{"3", args{arr: []int{5, 1, 4, 3, 2}}, 1},
		{"4", args{arr: []int{1, 2, 1, 4, 3}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.MinInts(tt.args.arr); got != tt.want {
				t.Errorf("MinInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMinStrings(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{arr: []string{"a", "b", "c", "d", "e"}}, 0},
		{"2", args{arr: []string{"elephant", "dog", "cat", "ball", "apple"}}, 4},
		{"3", args{arr: []string{"a", "e", "t", "b", "c"}}, 0},
		{"4", args{arr: []string{"a", "a", "d", "d", "c"}}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search.MinStrings(tt.args.arr); got != tt.want {
				t.Errorf("MinStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}