package cidergo

import "testing"

func TestEscape(t *testing.T) {
	e1 := escape("Hello, World!")
	if e1 != `"Hello, World!"` {
		t.Fail()
	}
	e2 := escape(`"Hello," " World~`)
	if e2 != `"\"Hello,\" \" World~"` {
		t.Fail()
	}
}

func TestItfToXxx(t *testing.T) {
	s := "abc"
	s2, err := itfToString(s)
	if err != nil || s2 != s {
		t.Fail()
	}

	i := 3
	i2, err := itfToInt(i)
	if err != nil || i2 != i {
		t.Fail()
	}

	_, err = itfToInt(s)
	if err == nil {
		t.Fail()
	}

	_, err = itfToString(i)
	if err == nil {
		t.Fail()
	}
}
