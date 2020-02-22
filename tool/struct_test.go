package tool

import (
	"testing"
)

type Src struct {
	Id   int
	Name string
	Sex  string
}

type Dst struct {
	Id   int
	Name string
	data []interface{}
}

func TestCopyProperty(t *testing.T) {
	s1 := Src{
		Id:   1,
		Name: "测试测试",
		Sex:  "男",
	}
	d1 := Dst{}
	t.Log(CopyStructProperty(&s1, &d1))
	t.Log(s1, d1)

	s2 := Src{
		Id:   1,
		Name: "测试测试",
		Sex:  "男",
	}
	d2 := Dst{
		Id:   2,
		Name: "gopher",
		data: make([]interface{}, 0),
	}
	t.Log(CopyStructProperty(&d2, &s2))
	t.Log(s2, d2)
}
