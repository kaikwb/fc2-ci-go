package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(15, 15)

	if total != 30 {
		t.Errorf("\r\nResultado da soma é inválido:\r\nResultado: %d\r\nEsperado: %d", total, 30)
	}
}
