package data

import "testing"

func TestLineToCSV(t *testing.T) {
	tests := []struct {
		in  string
		out []string
		err bool
	}{
		{
			in:  "jean,34,france",
			out: []string{"jean", "34", "france"},
			err: false,
		},
		{
			in:  "j,j,j,j",
			out: nil,
			err: true,
		},
		{
			in:  "j",
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := LineToCSV(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		for i, v := range test.out {
			if v != csv[i] {
				t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
			}
		}
	}
}

func TestReadFile(t *testing.T) {
	tests := []struct {
		in  string
		out []string
		err bool
	}{
		{
			in:  "test.csv",
			out: []string{"Damien,27,France", "Tom,5,USA", "QuentinV,8,Epicat"},
			err: false,
		},
		{
			in:  "j,j,j,j",
			out: nil,
			err: true,
		},
		{
			in:  "j",
			out: nil,
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := ReadFile(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		for i, v := range test.out {
			if v != csv[i] {
				t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
			}
		}
	}
}
