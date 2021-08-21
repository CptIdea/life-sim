package befunge

import "testing"

func TestExecuteCode(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"zero",args{"@"},0},
		{"one",args{">1@"},1},
		{"sum",args{">11+@"},2},
		{"diff",args{">21-@"},1},
		{"multiple",args{">22*@"},4},
		{"division",args{">48/@"},2},
		{"infinity",args{"><"},0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExecuteCode(tt.args.code); got != tt.want {
				t.Errorf("ExecuteCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
