package twenty23

import (
	"io"
	"strings"
	"testing"
)

func TestDayOne(t *testing.T) {
	type args struct {
		reader io.Reader
		part2  bool
	}
	tests := []struct {
		name       string
		args       args
		wantOutput string
		wantErr    bool
	}{
		{
			name: "Part One",
			args: args{strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`), false},
			wantOutput: "142",
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotOutput, err := DayOne(tt.args.reader, tt.args.part2)
			if (err != nil) != tt.wantErr {
				t.Errorf("DayOne() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotOutput != tt.wantOutput {
				t.Errorf("DayOne() gotOutput = %v, want %v", gotOutput, tt.wantOutput)
			}
		})
	}
}
