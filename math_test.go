package main

import "testing"

func TestSoma(t *testing.T) {

	total := Soma(15, 15)
	if total != 30 {
		t.Errorf("O Resultado da soma é invalido: O resultado %d. Espearado: %d", total, 30)
	}
}
