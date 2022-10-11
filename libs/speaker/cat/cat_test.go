package cat

import "testing"

func TestSpeak(t *testing.T) {
	result := Speak("Hello World")
	expected :=
		" _._     _,-'\"\"`-._\n" +
			"(,-.`._,'(       |\\`-/|\n" +
			"    `-.-' \\ )-`( , o o)     ---->   Hello World\n" +
			"          `-    \\`_`\"'-"
	if result != expected {
		t.Errorf("Result: \n%s\n\n Expected: \n%s\n\n", result, expected)
	}
}
