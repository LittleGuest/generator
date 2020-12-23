package tool_test

import (
	"testing"
)

func TestGetBirthByIdcard(t *testing.T) {
	t.Log(GetBirthByIdcard(""))
	t.Log(GetBirthByIdcard("11"))
	t.Log(GetBirthByIdcard("510823199609057250"))
}

func TestGetAgeByIdcard(t *testing.T) {
	t.Log(GetAgeByIdcard(""))
	t.Log(GetAgeByIdcard("11"))
	t.Log(GetAgeByIdcard("510823199609057250"))
}
