package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcularImpuesto(t *testing.T) {
	//Given
	var sueldo1 float64 = 30000
	var sueldo2 float64 = 60000
	var sueldo3 float64 = 160000

	imp_esperado1 := float64(0)
	imp_esperado2 := sueldo2 * 0.17
	imp_esperado3 := sueldo3 * 0.27

	//Act
	resultado1 := CalcularImpuesto(sueldo1)
	resultado2 := CalcularImpuesto(sueldo2)
	resultado3 := CalcularImpuesto(sueldo3)

	assert.Equal(t, imp_esperado1, resultado1, "El impuesto menor a 50.000 no es correcto")
	assert.Equal(t, imp_esperado2, resultado2, "El impuesto mayor a 50.000 no es correcto")
	assert.Equal(t, imp_esperado3, resultado3, "El impuesto mayor a 150.000 no es correcto")
	//assert.NotEqual()
	//assert.Nil()

}
