package compactnumber

import "testing"

func TestFormat(t *testing.T) {
	type args[T Formattable] struct {
		v T
	}
	type testCase[T Formattable] struct {
		name string
		args args[T]
		want string
	}
	tests := []testCase[uint64]{
		{name: "default", args: args[uint64]{v: 312}, want: "312"},
		{name: "thousands", args: args[uint64]{v: 1_000}, want: "1K"},
		{name: "thousands", args: args[uint64]{v: 2_499}, want: "2.5K"},
		{name: "thousands", args: args[uint64]{v: 3_500}, want: "3.5K"},
		{name: "thousands", args: args[uint64]{v: 4_600}, want: "4.6K"},
		{name: "thousands", args: args[uint64]{v: 4_999}, want: "5K"},
		{name: "thousands", args: args[uint64]{v: 5_100}, want: "5.1K"},
		{name: "millions", args: args[uint64]{v: 1_000_000}, want: "1M"},
		{name: "millions", args: args[uint64]{v: 2_100_000}, want: "2.1M"},
		{name: "millions", args: args[uint64]{v: 3_490_000}, want: "3.5M"},
		{name: "millions", args: args[uint64]{v: 4_500_000}, want: "4.5M"},
		{name: "millions", args: args[uint64]{v: 5_712_345}, want: "5.7M"},
		{name: "millions", args: args[uint64]{v: 6_999_999}, want: "7M"},
		{name: "billions", args: args[uint64]{v: 1_000_000_000}, want: "1B"},
		{name: "billions", args: args[uint64]{v: 1_990_000_000}, want: "2B"},
		{name: "billions", args: args[uint64]{v: 2_040_000_000}, want: "2B"},
		{name: "billions", args: args[uint64]{v: 3_490_000_000}, want: "3.5B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Format(tt.args.v); got != tt.want {
				t.Errorf("Format() = %v, want %v", got, tt.want)
			}
		})
	}
}
