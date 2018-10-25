package main

import (
	"testing"
)

func Test_genRandom(t *testing.T) {
	tests := []struct {
		name    string
		args    uint
		wantOut uint
		wantErr bool
	}{
		{name: "teste uint", args: 68, wantOut: 32, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOut, err := genRandom(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("genRandom() error=%v, wantErr %v", err, tt.wantErr)
				return
			}

			if gotOut != tt.wantOut {
				t.Errorf("genRandom() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

func Test_genRandom2(t *testing.T) {
	type args struct {
		beg int
		end int
	}
	tests := []struct {
		name    string
		beg     int
		end     int
		wantOut int
		wantErr string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genRandom2(tt.beg, tt.end); got != tt.wantOut {
				t.Errorf("genRandom2() = %v, want %v", got, tt.wantErr)
			}
		})
	}
}

func Test_parseArg(t *testing.T) {
	type args struct {
		param string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseArg(tt.args.param); got != tt.want {
				t.Errorf("parseArg() = %v, want %v", got, tt.want)
			}
		})
	}
}
