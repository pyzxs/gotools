package ArrUtil

import "testing"

func TestCollapse(t *testing.T) {
	m := make([]map[string]interface{}, 0)
	m = append(m, []map[string]interface{}{
		{"id": "name", "age": 20},
		{"id": "laozhu", "age": 15},
	}...,
	)
	l := Collapse(m)
	t.Log(l)
}

func TestRandom(t *testing.T) {
	m := []interface{}{1, 2, "3", 4}
	l := Random(m)
	t.Log(l)
}
