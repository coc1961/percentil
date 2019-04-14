package percentil

import (
	"errors"
	"math"
	"sort"
)

//New new
func New() *Percentil {
	return &Percentil{}
}

//Percentil Percentil
type Percentil struct {
	err         error
	result      float64
	data        [][]float64
	dataGrouped bool
}

//percentilTable percentilTable
type percentilTable struct {
	Value  float64
	Ni     float64
	NiAcum float64
}

//Error Error
func (p *Percentil) Error() error {
	return p.err
}

//Result Result
func (p *Percentil) Result() float64 {
	return p.result
}

//SetTable SetTable
func (p *Percentil) SetTable(data interface{}) *Percentil {
	p.result = -1
	p.data = nil
	p.dataGrouped = false
	if d, ok := data.([][]float64); ok {
		sort.Slice(d, func(i, j int) bool {
			return d[i][0] < d[j][0]
		})
		p.data = d
		p.dataGrouped = true
	} else if d, ok := data.([]float64); ok {
		p.data = make([][]float64, 0, len(d))
		for _, d1 := range d {
			p.data = append(p.data, []float64{d1, 0})
		}
		sort.Slice(p.data, func(i, j int) bool {
			return p.data[i][0] < p.data[j][0]
		})
		p.dataGrouped = false
	} else {
		p.err = errors.New("Invalid Data Type")
		p.result = -1
		p.data = nil
		return p
	}
	return p
}

//Calc Calc
func (p *Percentil) Calc(perc int) *Percentil {
	if p.data == nil {
		p.err = errors.New("Invalid Data Type")
		p.result = -1
		return p
	}
	if perc < 0 || perc > 99 {
		p.err = errors.New("Invalid Percentil")
		p.result = -1
		return p
	}
	p.result = -1
	p.err = nil

	if p.dataGrouped {
		return p.calcGrouped(perc)
	}
	return p.calcNoGrouped(perc)
}

func (p *Percentil) calcNoGrouped(perc int) *Percentil {
	data := p.data
	table := make([]percentilTable, len(data))
	acum := float64(0)
	for i := 0; i < len(data); i++ {
		acum += data[i][1]
		table[i] = percentilTable{
			Value:  data[i][0],
			Ni:     data[i][1],
			NiAcum: acum,
		}
	}
	N := float64(len(table))
	result := N * (float64(perc) / 100)
	if result == float64(int(result)) {
		p.result = table[int(result-1)].Value
		return p
	}
	value1 := table[int(result-1)].Value
	value2 := table[int(result)].Value
	decimal := (result - float64(int(result)))
	diff := (value2 - value1) * decimal
	p.result = math.Round((value1+diff)*100) / 100
	return p
}

func (p *Percentil) calcGrouped(perc int) *Percentil {
	table := make([]percentilTable, len(p.data))
	acum := float64(0)
	for i := 0; i < len(p.data); i++ {
		acum += p.data[i][1]
		table[i] = percentilTable{
			Value:  p.data[i][0],
			Ni:     p.data[i][1],
			NiAcum: acum,
		}
	}
	if len(table) == 0 {
		p.err = errors.New("Empty Data")
		p.result = -1
		return p
	}
	Intervalo := float64(0)         // Intervalo de incremento de la columna 0 de la tabla
	K := perc                       // Percentil
	N := table[len(table)-1].NiAcum // total de elementos
	Pos := ((float64(K) * float64(N)) / float64(100))
	ind := 0

	prevNum := float64(-1)
	for ind = 0; ind < len(table); ind++ {
		if prevNum == -1 {
			prevNum = table[ind].Value
		} else {
			Intervalo = table[ind].Value - prevNum
			break
		}
	}
	for ind = 0; ind < len(table); ind++ {
		if table[ind].NiAcum >= Pos {
			break
		}
	}
	ind1 := ind - 1
	if ind1 < 0 {
		ind1 = 0
	}
	if Pos == float64(int(Pos)) {
		ind1 = ind
	}
	LI := table[ind].Value
	Fi1 := table[ind1].NiAcum
	Fi := table[ind].Ni
	PercentilCalculado := LI + ((Pos-Fi1)/Fi)*Intervalo
	p.result = math.Round(PercentilCalculado*100) / 100
	return p
}
