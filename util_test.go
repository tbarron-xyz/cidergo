package cidergo

import (
	"fmt"
	"testing"
)

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

	b := true
	b2, err := itfToBool(b)
	if err != nil || b2 != b {
		t.Fail()
	}

	_, err = itfToBool(s)
	if err == nil {
		t.Fail()
	}

	arrayItf := []interface{}{}
	arrayItf = append(arrayItf, "abc")
	arrayString, err := itfToArrayString(arrayItf)
	if err != nil || arrayString[0] != "abc" {
		t.Fail()
	}

	arrayItf = append(arrayItf, 5)
	_, err = itfToArrayString(arrayItf)
	if err == nil {
		t.Fail()
	}

	notArrayItf := []int{}
	_, err = itfToArrayString(notArrayItf)
	if err == nil {
		t.Fail()
	}
}

func TestSprintfEscape(t *testing.T) {
	s := SprintfEscape("abc %v %v", `d"ef`, `99`)
	if s != `abc "d\"ef" "99"` {
		fmt.Println("s is:", s)
		t.Fail()
	}
}
