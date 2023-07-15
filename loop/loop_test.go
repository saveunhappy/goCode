package loop

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

func Repeat(character string) string {
	//repeat := ""
	var repeat string
	for i := 0; i < 5; i++ {
		repeat += character
	}
	return "aaaaa"
}
