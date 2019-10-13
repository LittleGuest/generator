package generator

import (
	"testing"
)

func TestCodeDB_List(t *testing.T) {
	list := CodeDB{}.List()
	for k, v := range list {
		t.Log(k, v)
	}
}
