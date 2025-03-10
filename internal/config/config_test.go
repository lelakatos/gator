package config

import "testing"

func TestRead(t *testing.T) {
	actual := Read()
	expected := Config{DbURL: "postgres://example"}

	if actual != expected {
		t.Errorf("Structs are not matching, expected=%+v, got=%+v", expected, actual)
	}
}
