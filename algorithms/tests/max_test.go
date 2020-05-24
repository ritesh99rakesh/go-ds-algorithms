package tests

import (
	"github.com/ritesh99rakesh/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestMaxFloat64s(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 int
	}{
		{"1", args{arr: []float64{1.0, 2, 3.4, 4.78, 5.7}}, 5.7, 4},
		{"2", args{arr: []float64{5.1, 4.5, 3.4, 2.2, 1.3}}, 5.1, 0},
		{"3", args{arr: []float64{1.4, 5.45, 4.6, 3.9, 2.0}}, 5.45, 1},
		{"4", args{arr: []float64{1.9, 2.3, 4.4, 4.4, 3.1}}, 4.4, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search.MaxFloat64s(tt.args.arr)
			if got != tt.want {
				t.Errorf("MaxFloat64s() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxFloat64s() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMaxInts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"1", args{arr: []int{1, 2, 3, 4, 5}}, 5, 4},
		{"2", args{arr: []int{5, 4, 3, 2, 1}}, 5, 0},
		{"3", args{arr: []int{1, 5, 4, 3, 2}}, 5, 1},
		{"4", args{arr: []int{1, 2, 4, 4, 3}}, 4, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search.MaxInts(tt.args.arr)
			if got != tt.want {
				t.Errorf("MaxInts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxInts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMaxStrings(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		{"1", args{arr: []string{"a", "b", "c", "d", "e"}}, "e", 4},
		{"2", args{arr: []string{"elephant", "dog", "cat", "ball", "apple"}}, "elephant", 0},
		{"3", args{arr: []string{"a", "e", "t", "b", "c"}}, "t", 2},
		{"4", args{arr: []string{"a", "b", "d", "d", "c"}}, "d", 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search.MaxStrings(tt.args.arr)
			if got != tt.want {
				t.Errorf("MaxStrings() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MaxStrings() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
