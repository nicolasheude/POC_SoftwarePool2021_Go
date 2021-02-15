package humanity

import "testing"

func TestNewHumanFromCSV(t *testing.T) {
	tests := []struct {
		in  []string
		out Human
		err bool
	}{
		{
			in:  []string{"jean", "34", "france"},
			out: Human{"jean", 34, "france"},
			err: false,
		},
		{
			in:  []string{"jean", "jean", "france"},
			out: Human{},
			err: true,
		},
		{
			in:  []string{"jean"},
			out: Human{},
			err: true,
		},
	}

	for _, test := range tests {
		csv, err := NewHumanFromCSV(test.in)
		if test.err && err != nil {
			continue
		}
		if err != nil {
			t.Errorf("Error when converting line [%s]: %v", test.in, err)
			continue
		}
		if test.out != *csv {
			t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
		}
	}
}

// func TestHumansFromCsvFile(t *testing.T) {
// 	tests := []struct {
// 		in  string
// 		out []Human
// 		err bool
// 	}{
// 		{
// 			in:  "test.csv",
// 			out: []Human{{"jean", 34, "france"}, {"nico", 18, "france"}},
// 			err: false,
// 		},
// 		{
// 			in:  "j, j",
// 			out: []Human{},
// 			err: true,
// 		},
// 		{
// 			in:  "",
// 			out: []Human{},
// 			err: true,
// 		},
// 	}

// 	for _, test := range tests {
// 		csv, err := NewHumansFromCsvFile(test.in)
// 		if test.err && err != nil {
// 			continue
// 		}
// 		if err != nil {
// 			t.Errorf("Error when converting line [%s]: %v", test.in, err)
// 			continue
// 		}
// 		if *csv != test.out {
// 			t.Errorf("Error when converting line [%s]:\n(exp)%v != (got)%v", test.in, test.out, csv)
// 		}
// 	}
// }
