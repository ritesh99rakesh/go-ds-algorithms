package tests

import (
	"github.com/ritesh99rakesh/go-ds-algorithms/algorithms/search"
	"testing"
)

func TestMinFloat64s(t *testing.T) {
	type args struct {
		arr []float64
	}
	tests := []struct {
		name  string
		args  args
		want  float64
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search.MinFloat64s(tt.args.arr)
			if got != tt.want {
				t.Errorf("MinFloat64s() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinFloat64s() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMinInts(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search.MinInts(tt.args.arr)
			if got != tt.want {
				t.Errorf("MinInts() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinInts() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestMinStrings(t *testing.T) {
	type args struct {
		arr []string
	}
	tests := []struct {
		name  string
		args  args
		want  string
		want1 int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := search.MinStrings(tt.args.arr)
			if got != tt.want {
				t.Errorf("MinStrings() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("MinStrings() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}