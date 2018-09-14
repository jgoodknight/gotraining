package strings

import (
	"testing"
)

type reverseTest struct {
	name          string
	test_input    string
	wanted_output string
}

func TestReverse(t *testing.T) {

	tests := []reverseTest{
		{"odd", "Jacob", "bocaJ"},
		{"one", "l", "l"},
		{"even", "looney", "yenool"},
		{"empty", "", ""},
		{"palindrome", "gohangasalamiimalasagnahog", "gohangasalamiimalasagnahog"},
		{"multibyte", "I ❤ NY", "YN ❤ I"},
	}
	for _, testStruct := range tests {

		test_functor := func(t *testing.T) {
			test_input := testStruct.test_input
			wanted_output := testStruct.wanted_output
			received_output := Reverse(test_input)

			if wanted_output != received_output {
				t.Errorf("Reverse of %q = %q, wanted %q", test_input, received_output, wanted_output)
			}
		}
		t.Run(testStruct.name, test_functor)
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse("abcdefghijklmnopqrstuvwxyz")
	}
}
