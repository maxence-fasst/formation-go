package format

import "testing"

func TestToBooleanString(t *testing.T) {
	type args struct {
		b bool
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should map true to oui",
			args: args{b: true},
			want: "oui",
		},
		{
			name: "should map false to non",
			args: args{b: false},
			want: "non",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToBooleanString(tt.args.b); got != tt.want {
				t.Errorf("ToBooleanString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToNumberString(t *testing.T) {
	type args struct {
		f int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should toString the number",
			args: args{f: 123},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToNumberString(tt.args.f); got != tt.want {
				t.Errorf("ToNumberString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestToFloatString(t *testing.T) {
	type args struct {
		f float64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "should work with no decimals",
			args: args{123},
			want: "123",
		},
		{
			name: "should work with few decimals",
			args: args{123.456},
			want: "123.456",
		},
		{
			name: "should work with decimals",
			args: args{123.456000789123},
			want: "123.456000789123",
		},
		{
			name: "should remove trailing zeros",
			args: args{123.456789000},
			want: "123.456789",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ToFloatString(tt.args.f); got != tt.want {
				t.Errorf("ToFloatString() = %v, want %v", got, tt.want)
			}
		})
	}
}
