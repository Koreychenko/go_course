package main

import "testing"

func TestRot13Transform(t *testing.T) {
	for input, want := range getTestCases() {
		t.Run(input, func(t *testing.T) {
			got, err := Rot13Transform(input)

			if err != nil {
				t.Errorf("en error occured %s", err)
			}

			assertCorrectMessage(t, got, want)
		})
	}
}

func getTestCases() map[string]string {
	return map[string]string{
		"Lbh penpxrq gur pbqr!": "You cracked the code!",
		"A":                     "N",
	}
}

func assertCorrectMessage(t *testing.T, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
