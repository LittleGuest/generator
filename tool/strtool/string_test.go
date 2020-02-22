package strtool

import "testing"

func TestToCamelCase(t *testing.T) {
	t.Log(ToCamelCase("dd_dd", "_"))
}

func TestToPascal(t *testing.T) {
	t.Log(ToPascal("dd_dd", "_"))
}

func TestFirstLetterToUpper(t *testing.T) {
	t.Log(FirstLetterToUpper("gopher"))
	t.Log(FirstLetterToUpper("Gopher"))
}

func TestToString(t *testing.T) {
	t.Log(ToString(333))
	t.Log(ToString("333"))

	stru := struct {
		Id   string
		Name string
	}{
		Id:   "999",
		Name: "gopher",
	}
	t.Log(ToString(stru))
}

func TestIsBlank(t *testing.T) {
	t.Log(IsBlank("gopher"))
	t.Log(IsBlank(""))
	t.Log(IsBlank(" "))
}

func TestIsNotBlank(t *testing.T) {
	t.Log(IsNotBlank("gopher"))
	t.Log(IsNotBlank(""))
	t.Log(IsNotBlank(" "))
}
