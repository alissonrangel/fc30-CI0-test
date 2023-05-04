package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(30, 30)

	if total != 900 {
		t.Errorf("Resultado da some é inválido: Resultado %d. Esperado: %d", total, 900)
	}
}
