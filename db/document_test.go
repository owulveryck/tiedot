package db_test

import (
	"github.com/owulveryck/tiedot/db"
	"testing"
)

type testStruct struct {
	ID      string
	Field1  int64
	Field3  float64
	private string
	Struct  subStruct
}

type subStruct struct {
	ID string
}

func TestUnmarshal(t *testing.T) {
	var test testStruct
	vals := map[string]interface{}{
		"ID":     "foo",
		"Field1": int64(1),
		"Field3": float64(2.0),
		"Struct": subStruct{"subid"},
	}
	err := db.Unmarshal(vals, &test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(test)
	vals = map[string]interface{}{
		"id":     "foo",
		"Field1": int64(1),
		"Field3": float64(2.0),
	}
	err = db.Unmarshal(vals, &test)
	if err != nil {
		t.Log(err)
	} else {
		t.Fail()
	}
}

func TestMarshal(t *testing.T) {
	test := testStruct{
		ID:     "foo",
		Field1: int64(1),
		Field3: float64(2.1),
		Struct: subStruct{"subid"},
	}
	t.Log(test)
	it := db.Marshal(test)
	t.Log(it)
	err := db.Unmarshal(it, &test)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(test)
}
