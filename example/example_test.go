package example

import (
	"testing"
)

func TestAdd(t *testing.T) {
	r := Add(2, 4)
	if r != 6 {
		t.Fatalf("Add(2, 4) error, expect:%d, actual:%d", 6, r)
	}
	t.Logf("test add succ")
}
func TestSub(t *testing.T) {
	r := Sub(3, 1)
	if r != 2 {
		t.Fatalf("Sub(3, 1) error, expect:%d, actual:%d", 2, r)
	}
	t.Logf("test Sub succ")
}
