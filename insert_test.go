package main

import "testing"

func TestInsert(t *testing.T) {
	err := InsertPerson("Arief", "Hidayatulloh")
	if err != nil {
		t.Error("Gagal insert person. " + err.Error())
	}
}
