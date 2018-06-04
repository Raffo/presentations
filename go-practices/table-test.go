func TestFlagParser(tt *testing.T) {
	for _, test := range []struct {
		name        string
		input       int
		expectedVal bool
	}{
		{"1 is not greater than 5", 1, false},
		{"3 is not greater than 5", 3, false},
		{"5 is not greater than 5", 5, false},
		{"7 is not greater than 5", 7, true},
	} {
		tt.Run(fmt.Sprintf("%v", test.name), func(t *testing.T) {
			tt.Log(test.name)
			out := isGreaterThanFive(test.input)
			if out != test.expectedVal {
				t.Errorf("Expected %b got %b with input %d", test.expectedVal, out, test.input)
			}
		})
	}
}
