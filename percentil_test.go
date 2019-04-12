package percentil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPercentil_Prueba1(t *testing.T) {
	table := [][]float64{
		{0, 0},
		{1, 10},
		{2, 10},
		{3, 10},
		{4, 10},
		{5, 10},
		{6, 10},
		{7, 10},
		{8, 10},
		{9, 10},
		{10, 10},
	}

	p := Percentil{}
	p.SetTable(table)

	assert.Equal(t, float64(1), p.Calc(10).Result())
	assert.Equal(t, float64(2.5), p.Calc(25).Result())
	assert.Equal(t, float64(5), p.Calc(50).Result())
	assert.Equal(t, float64(7.5), p.Calc(75).Result())
	assert.Equal(t, float64(8.5), p.Calc(85).Result())
	assert.Equal(t, float64(9.9), p.Calc(99).Result())

}

func TestPercentil_Cocina(t *testing.T) {
	table := [][]float64{
		{152, 5},
		{160, 18},
		{168, 42},
		{176, 27},
		{184, 9},
	}
	p := Percentil{}
	p.SetTable(table)

	assert.Equal(t, float64(175.16), p.Calc(60).Result())
	assert.Equal(t, float64(177.99), p.Calc(71).Result())
	assert.Equal(t, float64(180.68), p.Calc(80).Result())
	assert.Equal(t, float64(191.1), p.Calc(99).Result())

}

func TestPercentil_Decimal1(t *testing.T) {
	table := [][]float64{
		{200, 85},
		{300, 90},
		{400, 120},
		{500, 70},
		{600, 62},
		{700, 36},
	}

	p := Percentil{}
	p.SetTable(table)
	assert.Equal(t, float64(334.17), p.Calc(25).Result())
	assert.Equal(t, float64(574.64), p.Calc(75).Result())
	assert.Equal(t, float64(787.14), p.Calc(99).Result())

}

func TestPercentil_Decimal(t *testing.T) {
	table := [][]float64{
		{0, 21},
		{1, 32},
		{2, 38},
		{3, 34},
		{4, 48},
		{5, 57},
		{6, 82},
		{7, 73},
		{8, 53},
		{9, 39},
		{10, 23},
	}

	p := Percentil{}
	p.SetTable(table)

	assert.Equal(t, float64(0.91), p.Calc(10).Result())
	assert.Equal(t, float64(3), p.Calc(25).Result())
	assert.Equal(t, float64(5.24), p.Calc(50).Result())
	assert.Equal(t, float64(6.86), p.Calc(75).Result())
	assert.Equal(t, float64(7.75), p.Calc(85).Result())
	assert.Equal(t, float64(9.78), p.Calc(99).Result())

}

func TestPercentil_SimpleColumna(t *testing.T) {
	table := []float64{
		1,
		8,
		9,
		9,
		10,
		10,
		11,
		12,
		13,
		14,
		15,
		16,
		17,
		18,
		20,
	}

	p := Percentil{}
	p.SetTable(table)

	assert.Equal(t, float64(10), p.Calc(40).Result())

	table = []float64{
		0,
		4,
		1,
		0,
		0,
		7,
		2,
		1,
		4,
		0,
		3,
		9,
		2,
		0,
		0,
		4,
		8,
		1,
		0,
		9,
		4,
	}
	p.SetTable(table)

	assert.Equal(t, float64(7.5), p.Calc(90).Result())

}

//0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 7, 8, 9, 9
