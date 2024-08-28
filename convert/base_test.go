package convert

import (
	"reflect"
	"testing"
)

func TestStrToBoolean(t *testing.T) {
	s := StrToBoolean("false")
	t.Log(s, reflect.TypeOf(s).Kind(), reflect.ValueOf(s).Interface())
}

func TestForceConvert(t *testing.T) {
	u := []struct {
		ID   int
		Name string
	}{
		{ID: 1, Name: "gotools"},
		{ID: 2, Name: "study"},
	}

	p := []map[string]interface{}{}
	err := ForceCovert(u, &p)
	if err != nil {
		t.Error(err)
	}
	t.Log(p)
}

func TestForceConvert1(t *testing.T) {
	m := []map[string]interface{}{
		{"id": 10, "name": "gotools"},
		{"id": 20, "name": "test"},
	}

	type User struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	t.Log(m)
	u := []User{}
	err := ForceCovert(m, &u)
	if err != nil {
		t.Error(err)
	}
	t.Log(u)

	u1 := User{}
	err  = ForceCovert(m[1], &u1)
	if err != nil {
		t.Error(err)
	}
	t.Log(u1)
}
