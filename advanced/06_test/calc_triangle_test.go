package main

import "testing"

func TestCalcTriangle(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1",args{a: 3,b: 4},3},
		{"2",args{a: 5,b: 12},13},
		{"3",args{a: 8,b: 15},17},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CalcTriangle(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("calcTriangle() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkCalcTriangle(b *testing.B)  {
	numA,numB,want := 3,4,5
	for i:=0;i<b.N;i++{
		if got := CalcTriangle(numA, numB); got != want {
			b.Errorf("calcTriangle() = %v, want %v", got,want)
		}
	}
}