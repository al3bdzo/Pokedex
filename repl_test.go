package main


import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input: "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: "something IS GOING WRONG",
			expected: []string{"something", "is", "going", "wrong"},
		}, 
		{
			input: "I need some   help, Please",
			expected: []string{"i", "need", "some", "help,", "please"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("something went wrong, actual: %v, expected: %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("something went wrong, actual Word: %v, expected Word: %v", word, expectedWord)
			}
		}
	}
}