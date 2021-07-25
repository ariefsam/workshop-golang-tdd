package main

import "testing"

func TestCreateTableName(t *testing.T) {
	err := CreateTableName()
	if err != nil {
		t.Error("Gagal buat tabel. " + err.Error())
	}
}
