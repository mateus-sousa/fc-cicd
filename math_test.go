package main

import (
	"testing"
)

func TestSoma(t *testing.T) {
	total := Sum(15, 15)
	if total != 30 {
		t.Errorf("Valor recebido %d não bate com o valor esperado %d", total, 30)
	}
}
